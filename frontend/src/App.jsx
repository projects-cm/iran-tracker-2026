import React, { useState, useCallback, useMemo } from 'react';
import { ReactFlow, Controls, Background, applyNodeChanges, applyEdgeChanges } from '@xyflow/react';
import '@xyflow/react/dist/style.css';
import FigureNode from './FigureNode';
import { Activity } from 'lucide-react';

// Dummy data embedded for local development (based on real Iranian leadership structure)
const DUMMY_FIGURES = [
  {
    id: "1", parentId: null, name: "Ali Khamenei",
    role: "Supreme Leader", status: "Dead", tier: 1,
  },
  {
    id: "2", parentId: "1", name: "Mojtaba Khamenei",
    role: "Son of Supreme Leader", status: "Missing", tier: 2,
  },
  {
    id: "3", parentId: "1", name: "Hossein Salami",
    role: "Commander of IRGC", status: "Alive", tier: 2,
  },
  {
    id: "4", parentId: "1", name: "Ahmad Vahidi",
    role: "Secretary of SNSC", status: "Critically Wounded", tier: 2,
  },
  {
    id: "5", parentId: "3", name: "Esmail Qaani",
    role: "Commander Quds Force", status: "Presumed Dead", tier: 3,
  },
  {
    id: "6", parentId: "3", name: "Amir Ali Hajizadeh",
    role: "Commander Aerospace Force", status: "Dead", tier: 3,
  },
  {
    id: "7", parentId: "4", name: "Masoud Pezeshkian",
    role: "President", status: "Missing", tier: 3,
  },
];

function buildGraph(data) {
  const nodes = [];
  const edges = [];

  if (!data || data.length === 0) return { nodes, edges };

  // Count items per tier for horizontal layout
  const tierCounts = {};
  const tierOffsets = {};
  data.forEach(fig => {
    tierCounts[fig.tier] = (tierCounts[fig.tier] || 0) + 1;
    tierOffsets[fig.tier] = 0;
  });

  data.forEach((fig) => {
    const count = tierCounts[fig.tier];
    const offset = tierOffsets[fig.tier];
    const spacing = 300;
    const x = (offset - (count - 1) / 2) * spacing;
    const y = (fig.tier - 1) * 240;

    tierOffsets[fig.tier]++;

    nodes.push({
      id: String(fig.id),
      type: 'figureNode',
      position: { x: 500 + x, y: y + 50 },
      data: {
        name: fig.canonical_name,
        role: fig.role,
        status: fig.current_status,
      },
    });

    if (fig.parent_id) {
      edges.push({
        id: `e${fig.parent_id}-${fig.id}`,
        source: String(fig.parent_id),
        target: String(fig.id),
        type: 'smoothstep',
        animated: fig.current_status !== 'Dead',
        style: { stroke: 'rgba(255, 255, 255, 0.15)', strokeWidth: 2 },
      });
    }
  });

  return { nodes, edges };
}

function App() {
  const [figures, setFigures] = useState([]);
  const [nodes, setNodes] = useState([]);
  const [edges, setEdges] = useState([]);
  const [loading, setLoading] = useState(true);

  React.useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch('/api/v1/figures');
        const data = await response.json();
        setFigures(data);
        const { nodes: newNodes, edges: newEdges } = buildGraph(data);
        setNodes(newNodes);
        setEdges(newEdges);
      } catch (error) {
        console.error("Failed to fetch intel:", error);
      } finally {
        setLoading(false);
      }
    };
    fetchData();
  }, []);

  const nodeTypes = useMemo(() => ({ figureNode: FigureNode }), []);

  const onNodesChange = useCallback((changes) => setNodes((nds) => applyNodeChanges(changes, nds)), []);
  const onEdgesChange = useCallback((changes) => setEdges((eds) => applyEdgeChanges(changes, eds)), []);

  const aliveCount = figures.filter(f => f.current_status === 'Alive').length;
  const deadCount = figures.filter(f => ['Dead', 'Presumed Dead'].includes(f.current_status)).length;

  if (loading) {
    return (
      <div className="h-screen w-full flex items-center justify-center bg-slate-950">
        <div className="flex flex-col items-center gap-4">
          <Activity className="w-12 h-12 text-green-500 animate-pulse" />
          <div className="text-slate-400 font-mono tracking-widest animate-pulse">CONNECTING TO INTEL NETWORK...</div>
        </div>
      </div>
    );
  }

  return (
    <div className="h-screen w-full flex flex-col bg-slate-950">
      {/* Header */}
      <header className="h-16 flex items-center justify-between px-6 border-b border-slate-800 glass-panel z-10 relative shrink-0">
        <div className="flex items-center gap-3">
          <Activity className="text-green-400 w-6 h-6" />
          <h1 className="text-xl font-bold tracking-widest text-slate-200">
            IRAN CASUALTY TRACKER
            <span className="text-xs text-red-500 font-mono ml-2 border border-red-500/50 px-2 py-0.5 rounded">CLASSIFIED</span>
          </h1>
        </div>
        <div className="flex gap-6 text-sm text-slate-400 font-mono">
          <span>{figures.length} Targets</span>
          <span className="text-green-400">{aliveCount} Alive</span>
          <span className="text-red-400">{deadCount} KIA</span>
          <span className="text-green-400">● Live Feed</span>
        </div>
      </header>

      {/* React Flow Canvas */}
      <main className="flex-1 w-full relative">
        <ReactFlow
          nodes={nodes}
          edges={edges}
          onNodesChange={onNodesChange}
          onEdgesChange={onEdgesChange}
          nodeTypes={nodeTypes}
          nodesDraggable={false}
          nodesConnectable={false}
          elementsSelectable={false}
          fitView
          fitViewOptions={{ padding: 0.3 }}
          className="bg-slate-950"
        >
          <Background color="#1e293b" gap={20} size={1} />
          <Controls className="fill-slate-400" />
        </ReactFlow>
      </main>
    </div>
  );
}

export default App;
