import React, { useState, useEffect, useCallback, useMemo } from 'react';
import { ReactFlow, Controls, Background, applyNodeChanges, applyEdgeChanges } from '@xyflow/react';
import '@xyflow/react/dist/style.css';
import FigureNode from './FigureNode';
import { Activity } from 'lucide-react';

function App() {
  const [nodes, setNodes] = useState([]);
  const [edges, setEdges] = useState([]);

  const nodeTypes = useMemo(() => ({ figureNode: FigureNode }), []);

  const onNodesChange = useCallback((changes) => setNodes((nds) => applyNodeChanges(changes, nds)), []);
  const onEdgesChange = useCallback((changes) => setEdges((eds) => applyEdgeChanges(changes, eds)), []);

  useEffect(() => {
    // Fetch dummy data from backend
    fetch('/api/v1/figures')
      .then(res => res.json())
      .then(data => {
        const initialNodes = [];
        const initialEdges = [];

        // Very basic layout generation based on tier
        const tierCounts = { 1: 0, 2: 0, 3: 0 };
        const tierOffsets = { 1: 0, 2: 0, 3: 0 };

        data.forEach(fig => tierCounts[fig.tier]++);

        data.forEach((fig) => {
          // Layout math
          const x = (tierOffsets[fig.tier] - (tierCounts[fig.tier]-1)/2) * 300;
          const y = (fig.tier - 1) * 200;
          
          tierOffsets[fig.tier]++;

          initialNodes.push({
            id: fig.id,
            type: 'figureNode',
            position: { x: window.innerWidth/2 + x - 100, y: y + 100 },
            data: { 
              name: fig.name, 
              role: fig.role, 
              status: fig.status 
            }
          });

          if (fig.parentId) {
            initialEdges.push({
              id: `e${fig.parentId}-${fig.id}`,
              source: fig.parentId,
              target: fig.id,
              type: 'smoothstep',
              animated: fig.status !== 'Dead',
              style: { stroke: 'rgba(255, 255, 255, 0.2)', strokeWidth: 2 }
            });
          }
        });

        setNodes(initialNodes);
        setEdges(initialEdges);
      });
  }, []);

  return (
    <div className="h-screen w-full flex flex-col bg-slate-950">
      <header className="h-16 flex items-center justify-between px-6 border-b border-slate-800 glass-panel z-10 relative">
        <div className="flex items-center gap-3">
          <Activity className="text-neon-green" />
          <h1 className="text-xl font-bold tracking-widest text-slate-200">IRAN CASUALTY TRACKER <span className="text-xs text-red-500 font-mono ml-2 border border-red-500/50 px-2 py-0.5 rounded">CLASSIFIED</span></h1>
        </div>
        <div className="flex gap-4 text-sm text-slate-400 font-mono">
          <span>{nodes.length} Targets Tracked</span>
          <span className="text-neon-green">System Online</span>
        </div>
      </header>
      
      <main className="flex-1 w-full h-full relative">
        <ReactFlow
          nodes={nodes}
          edges={edges}
          onNodesChange={onNodesChange}
          onEdgesChange={onEdgesChange}
          nodeTypes={nodeTypes}
          fitView
          className="bg-slate-950"
        >
          <Background color="#1e293b" gap={20} />
          <Controls className="fill-slate-400" />
        </ReactFlow>
      </main>
    </div>
  );
}

export default App;
