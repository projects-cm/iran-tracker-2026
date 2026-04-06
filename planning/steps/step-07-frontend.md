# Step 07: React Dashboard (Frontend) ✅

## Goal:
Build the Vite/React application with a hierarchical node-based UI for tracking.

## Actions:
1.  **Vite App Creation**:
    -   Initialized via `yarn create vite frontend --template react`.
    -   Installed: `@xyflow/react`, `lucide-react`, `tailwindcss`, `@tailwindcss/vite`.
2.  **UI Implementation** (`src/App.jsx`, `src/FigureNode.jsx`):
    -   **React Flow** hierarchical node network reflecting military/political chain of command.
    -   Custom `FigureNode` component with glassmorphism styling.
    -   Status-based glow effects: Green (Alive), Orange (Wounded), Red (Dead), Gray (Missing).
    -   Nodes locked in position (no dragging) — layout is predetermined by tier.
    -   Animated edges between parent-child nodes (disabled for Dead figures).
3.  **Header Bar**:
    -   "IRAN CASUALTY TRACKER" title with red "CLASSIFIED" badge.
    -   Live stats: target count, alive count, KIA count.
    -   "System Online" status indicator.
4.  **Embedded Dummy Data**:
    -   7 figures with realistic statuses for standalone UI development.
    -   No backend dependency required for frontend iteration.

## Running Locally:
```bash
.\run_frontend.bat
# Opens at http://localhost:5173/
```

## Pending:
-   Connect to live backend API (replace embedded data with fetch).
-   Add intel feed sidebar with report history.
-   Add click-to-expand detail view for each figure.

## Status: COMPLETE — UI functional with dummy data
