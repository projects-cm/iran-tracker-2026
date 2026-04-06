import React from 'react';
import { Handle, Position } from '@xyflow/react';

const statusColorMap = {
  "Alive": "text-green-400 bg-green-400/10",
  "Missing": "text-gray-400 bg-gray-400/10",
  "Critically Wounded": "text-orange-400 bg-orange-400/10",
  "Dead": "text-red-500 bg-red-500/10",
  "Presumed Dead": "text-red-400 bg-red-400/10"
};

const statusClassMap = {
  "Alive": "status-Alive",
  "Missing": "status-Missing",
  "Critically Wounded": "status-CriticallyWounded",
  "Dead": "status-Dead",
  "Presumed Dead": "status-Dead"
};

function FigureNode({ data }) {
  const statusColor = statusColorMap[data.status] || "text-gray-400";
  const glowClass = statusClassMap[data.status] || "status-Missing";

  return (
    <div className={`glass-panel p-4 min-w-[200px] flex flex-col items-center justify-center border-l-4 ${glowClass}`}>
      <Handle type="target" position={Position.Top} className="opacity-0" />
      
      <div className="w-16 h-16 rounded-full bg-slate-800 mb-2 border-2 border-slate-700 flex items-center justify-center overflow-hidden">
         {/* Placeholder for actual image */}
         <span className="text-xl font-bold text-slate-400">{data.name.charAt(0)}</span>
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

export default FigureNode;
