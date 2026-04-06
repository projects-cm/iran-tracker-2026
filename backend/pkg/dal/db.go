package dal

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

// DB represents the database access layer
type DB struct {
	db *sql.DB
}

// NewDB opens a connection to Turso and initializes the schema
func NewDB(dbURL, authToken string) (*DB, error) {
	connStr := fmt.Sprintf("%s?authToken=%s", dbURL, authToken)
	db, err := sql.Open("libsql", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to turso: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping turso: %w", err)
	}

	d := &DB{db: db}
	if err := d.initSchema(); err != nil {
		return nil, err
	}
	return d, nil
}

func (d *DB) GetDB() *sql.DB {
	return d.db
}

func (d *DB) initSchema() error {
	schema := `
	CREATE TABLE IF NOT EXISTS figures (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		canonical_name TEXT UNIQUE NOT NULL,
		persian_name TEXT,
		role TEXT,
		tier INTEGER,
		current_status TEXT,
		parent_id INTEGER,
		last_update_id INTEGER,
		FOREIGN KEY(parent_id) REFERENCES figures(id)
	);

	CREATE TABLE IF NOT EXISTS reports (
		message_id INTEGER PRIMARY KEY,
		source TEXT NOT NULL,
		headline TEXT NOT NULL,
		raw_text TEXT,
		confidence_level INTEGER,
		status TEXT,
		previous_status TEXT,
		tier INTEGER,
		timestamp TEXT NOT NULL,
		entity_id INTEGER,
		FOREIGN KEY(entity_id) REFERENCES figures(id)
	);

	CREATE TABLE IF NOT EXISTS aliases (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		entity_id INTEGER,
		alias TEXT UNIQUE NOT NULL,
		FOREIGN KEY(entity_id) REFERENCES figures(id)
	);
	`
	_, err := d.db.Exec(schema)
	if err != nil {
		return fmt.Errorf("failed to initialize schema: %w", err)
	}
	return nil
}

// CreateFigure inserts a new figure and its aliases into the database
func (d *DB) CreateFigure(ctx context.Context, name, persian string, tier int, status string, role string, parentID *int, aliases []string) (int64, error) {
	res, err := d.db.ExecContext(ctx, "INSERT OR IGNORE INTO figures (canonical_name, persian_name, tier, current_status, role, parent_id) VALUES (?, ?, ?, ?, ?, ?)",
		name, persian, tier, status, role, parentID)
	if err != nil {
		return 0, fmt.Errorf("failed to insert figure %s: %w", name, err)
	}

	id, _ := res.LastInsertId()
	if id == 0 {
		err = d.db.QueryRowContext(ctx, "SELECT id FROM figures WHERE canonical_name = ?", name).Scan(&id)
		if err != nil {
			return 0, fmt.Errorf("failed to find figure id %s: %w", name, err)
		}
	}

	for _, alias := range aliases {
		_, err = d.db.ExecContext(ctx, "INSERT OR IGNORE INTO aliases (entity_id, alias) VALUES (?, ?)", id, alias)
		if err != nil {
			return 0, fmt.Errorf("failed to insert alias %s for %s: %w", alias, name, err)
		}
	}

	return id, nil
}

// Figure represents an Iranian leadership entity
type Figure struct {
	ID            int    `json:"id"`
	CanonicalName string `json:"canonical_name"`
	PersianName   string `json:"persian_name"`
	Role          string `json:"role"`
	Tier          int    `json:"tier"`
	CurrentStatus string `json:"current_status"`
	ParentID      *int   `json:"parent_id"`
	LastUpdateID  int    `json:"last_update_id"`
}

// Report represents a status update for a figure
type Report struct {
	MessageID       int    `json:"message_id"`
	Source          string `json:"source"`
	Headline        string `json:"headline"`
	RawText         string `json:"raw_text"`
	ConfidenceLevel int    `json:"confidence_level"`
	Status          string `json:"status"`
	PreviousStatus  string `json:"previous_status"`
	Tier            int    `json:"tier"`
	Timestamp       string `json:"timestamp"`
	EntityID        int    `json:"entity_id"`
}

func (d *DB) GetFigures(ctx context.Context) ([]Figure, error) {
	rows, err := d.db.QueryContext(ctx, "SELECT id, canonical_name, persian_name, role, tier, current_status, parent_id, last_update_id FROM figures")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var figures []Figure
	for rows.Next() {
		var f Figure
		var lastUpdateID sql.NullInt64
		if err := rows.Scan(&f.ID, &f.CanonicalName, &f.PersianName, &f.Role, &f.Tier, &f.CurrentStatus, &f.ParentID, &lastUpdateID); err != nil {
			return nil, err
		}
		if lastUpdateID.Valid {
			f.LastUpdateID = int(lastUpdateID.Int64)
		}
		figures = append(figures, f)
	}
	return figures, nil
}

func (d *DB) AddReport(ctx context.Context, r Report) error {
	tx, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, `
		INSERT INTO reports (message_id, source, headline, raw_text, confidence_level, status, previous_status, tier, timestamp, entity_id)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		r.MessageID, r.Source, r.Headline, r.RawText, r.ConfidenceLevel, r.Status, r.PreviousStatus, r.Tier, r.Timestamp, r.EntityID)
	if err != nil {
		return fmt.Errorf("failed to insert report: %w", err)
	}

	_, err = tx.ExecContext(ctx, "UPDATE figures SET current_status = ?, last_update_id = ? WHERE id = ?", r.Status, r.MessageID, r.EntityID)
	if err != nil {
		return fmt.Errorf("failed to update figure status: %w", err)
	}

	return tx.Commit()
}

// IsReportProcessed checks if a message ID has already been handled
func (d *DB) IsReportProcessed(ctx context.Context, msgID int) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM reports WHERE message_id = ?)"
	err := d.db.QueryRowContext(ctx, query, msgID).Scan(&exists)
	return exists, err
}
