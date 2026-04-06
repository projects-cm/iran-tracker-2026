# Makefile for Iranian Leadership Casualty Tracker

.PHONY: run-backend run-frontend seed auth test install build lint test-all

# --- Development ---

run-backend:
	cd backend && go run ./cmd/server

run-frontend:
	cd frontend && yarn dev

seed:
	cd backend && go run ./cmd/admin/seed

auth:
	cd backend && go run ./cmd/admin/auth

# --- Testing ---

test:
	cd backend && go test ./...

test-v:
	cd backend && go test -v ./...

# --- Setup ---

install:
	cd backend && go mod download
	cd frontend && yarn install

# --- Production ---

build:
	cd backend && go build -o ../bin/server ./cmd/server
	cd frontend && yarn build

lint:
	cd backend && golangci-lint run
