/**
 * Shared domain types for the Iran Casualty Tracker frontend.
 * These mirror the JSON shapes returned by the Go backend.
 */

export type FigureStatus =
  | 'Alive'
  | 'Dead'
  | 'Presumed Dead'
  | 'Critically Wounded'
  | 'Missing';

export interface Figure {
  id: number;
  canonical_name: string;
  persian_name: string;
  role: string;
  tier: number;
  current_status: FigureStatus;
  parent_id: number | null;
  last_update_id: number;
}

export interface Report {
  message_id: number;
  source: string;
  headline: string;
  raw_text: string;
  confidence_level: number;
  status: FigureStatus;
  previous_status: FigureStatus;
  tier: number;
  timestamp: string;
  entity_id: number;
}

/** Data embedded in each ReactFlow node */
export interface FigureNodeData {
  name: string;
  role: string;
  status: FigureStatus;
  [key: string]: unknown;
}

/** Config shape for a single ad slot */
export interface AdSlotConfig {
  enabled: boolean;
  label: string;
  html: string | null;
  fallbackWidth: number;
  fallbackHeight: number;
}

/** Full ad configuration */
export interface AdConfig {
  enabled: boolean;
  slots: Record<string, AdSlotConfig>;
}
