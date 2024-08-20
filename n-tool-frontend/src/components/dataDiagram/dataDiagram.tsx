"use client";

import React, { useState, useMemo } from 'react';
import ReactFlow, {
  ReactFlowProvider,
  Background,
  Controls,
  Node,
  Edge,
} from 'react-flow-renderer';
import { debugLog } from '@/utils/debugLogUtil';
import { NodeComponent, EdgeComponent, NodeDetailsDrawer, CustomTooltip } from './dataDiagramUi';


export interface Data {
  Apps: string | null;
  IP: string;
  Hostname: string;
  Status: string;
  Exposure: string;
  InternetAccess: string;
  OS: string;
  OSVersion: string;
  Usage: string;
  Location: string;
  Owners: string[];
  Dependencies: string[];
  CreatedAt: string;
  CreatedBy: string;
  RecordedAt: string;
  Access: { User: string; Role: string }[];
  ConnectsTo: string[];
  HostType: string;
  ExposedServices: string[];
  CPU: number;
  RAM: number;
  Storage: number;
  OpenPorts: string[];
  Range: string;
}

interface DataDiagramProps {
  data: Data[];
}

const nodeTypes = {
  custom: NodeComponent,
};

const edgeTypes = {
  custom: EdgeComponent,
};

const DataDiagram: React.FC<DataDiagramProps> = ({ data }) => {
  const [selectedNode, setSelectedNode] = useState<Data | null>(null);

  const nodes = useMemo(() => {
    const generatedNodes = data.map((node, index) => ({
      id: node.IP,
      data: {
        label: <CustomTooltip hostname={node.Hostname} ip={node.IP} range={node.Range} />,
        details: `${node.Usage} | ${node.Status}`,
      },
      position: { x: (index % 4) * 250, y: Math.floor(index / 4) * 200 }, // Positioned in a grid
      type: 'custom',
      range: node.Range, // Ensure range is included
    }));
    debugLog("Generated nodes: ", generatedNodes);
    return generatedNodes;
  }, [data]);

  const edges = useMemo(() => {
    const edgesList: Edge[] = [];
    nodes.forEach((node) => {
      nodes.forEach((targetNode) => {
        if (node.id !== targetNode.id && node.range === targetNode.range) {
          edgesList.push({
            id: `${node.id}-${targetNode.id}`,
            source: node.id,
            target: targetNode.id,
            type: 'smoothstep',
            animated: true,
            style: { stroke: 'var(--DD-edge-color)', strokeWidth: 2, fill: 'none' },
            markerEnd: 'url(#edge-circle)',
          });
        }
      });
    });
    debugLog("Generated edges: ", edgesList); // Debugging log to check edges
    return edgesList;
  }, [nodes]);

  debugLog("Nodes: ", nodes);
  debugLog("Edges: ", edges);

  const handleNodeClick = (_: React.MouseEvent, node: Node) => {
    const clickedNode = data.find((item) => item.IP === node.id);
    setSelectedNode(clickedNode || null);
  };

  return (
    <ReactFlowProvider>
      <div style={{ height: 600, position: 'relative' }}>
        <ReactFlow
          nodes={nodes}
          edges={edges}
          snapToGrid={true}
          snapGrid={[15, 15]}
          onNodeClick={handleNodeClick}
          nodeTypes={nodeTypes}
          edgeTypes={edgeTypes}
          fitView
        >
          <Controls showInteractive={false} />
          <Background color="var(--DD-bg-color)" gap={16} />
          <svg>
            <defs>
              <marker
                id="edge-circle"
                viewBox="-5 -5 10 10"
                refX="0"
                refY="0"
                markerUnits="strokeWidth"
                markerWidth="10"
                markerHeight="10"
                orient="auto"
              >
                <circle stroke="var(--DD-edge-marker-color)" strokeOpacity="0.75" r="2" cx="0" cy="0" />
              </marker>
            </defs>
          </svg>
        </ReactFlow>
        <NodeDetailsDrawer selectedNode={selectedNode} onClose={() => setSelectedNode(null)} />
      </div>
    </ReactFlowProvider>
  );
};

export default DataDiagram;
