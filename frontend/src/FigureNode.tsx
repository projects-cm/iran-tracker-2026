import { Handle, Position } from '@xyflow/react';
import type { NodeProps, Node } from '@xyflow/react';
import type { FigureNodeData, FigureStatus } from './types';

const STATUS_COLOR_MAP: Record<FigureStatus, string> = {
  Alive: 'text-green-400 bg-green-400/10',
  Missing: 'text-gray-400 bg-gray-400/10',
  'Critically Wounded': 'text-orange-400 bg-orange-400/10',
  Dead: 'text-red-500 bg-red-500/10',
  'Presumed Dead': 'text-red-400 bg-red-400/10',
};

const STATUS_CLASS_MAP: Record<FigureStatus, string> = {
  Alive: 'status-Alive',
  Missing: 'status-Missing',
  'Critically Wounded': 'status-CriticallyWounded',
  Dead: 'status-Dead',
  'Presumed Dead': 'status-Dead',
};

// Fix the types in FigureNode.tsx
export default function FigureNode({ data }: NodeProps<Node<FigureNodeData>>) {
  const status = data.status || 'Missing';
  const statusColor = STATUS_COLOR_MAP[status as FigureStatus] ?? 'text-gray-400';
  const glowClass = STATUS_CLASS_MAP[status as FigureStatus] ?? 'status-Missing';

  return (
    <div
      className={`glass-panel p-4 min-w-[200px] flex flex-col items-center justify-center border-l-4 ${glowClass}`}
    >
      <Handle type="target" position={Position.Top} className="opacity-0" />

      <div className="w-16 h-16 rounded-full bg-slate-800 mb-2 border-2 border-slate-700 flex items-center justify-center overflow-hidden">
        <span className="text-xl font-bold text-slate-400">
          {data.name.charAt(0)}
        </span>
      </div>

      <h3 className="text-lg font-bold text-white text-center">{data.name}</h3>
      <p className="text-xs text-slate-300 mb-2 text-center">{data.role}</p>

      <div className={`px-3 py-1 rounded-full text-xs font-semibold ${statusColor}`}>
        {data.status.toUpperCase()}
      </div>

      <Handle type="source" position={Position.Bottom} className="opacity-0" />
    </div>
  );
}
