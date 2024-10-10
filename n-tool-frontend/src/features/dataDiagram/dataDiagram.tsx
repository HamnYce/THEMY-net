"use client";

import React, { useState, useEffect, useCallback } from 'react';
import ReactFlow, {
  ReactFlowProvider,
  useNodesState,
  useEdgesState,
  addEdge,
  Background,
  Controls,
  Connection,
  Edge,
  Node,
} from 'react-flow-renderer';
import { debugLog } from '@/utils/debugLogUtil';
import { NodeComponent, EdgeComponent, NodeDetailsDrawer, CustomTooltip } from './dataDiagramUi';

// Define data interface
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
  Range: string;  // The range property is defined here
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
  const [nodes, setNodes, onNodesChange] = useNodesState([]);
  const [edges, setEdges, onEdgesChange] = useEdgesState([]);
  const [selectedNode, setSelectedNode] = useState<Data | null>(null);

  useEffect(() => {
    // Explicitly define types for nodes and edges
    const generatedNodes = data.map((node, index) => ({
      id: node.IP,
      data: {
        label: <CustomTooltip hostname={node.Hostname} ip={node.IP} range={node.Range} />,
        details: `${node.Usage} | ${node.Status}`,
        range: node.Range,  // Include range here so it can be used for edge generation
      },
      position: { x: (index % 4) * 250, y: Math.floor(index / 4) * 200 }, 
      type: 'custom',
      draggable: true, // Ensure nodes are draggable
    }));

    const generatedEdges: Edge<any>[] = [];

    // Create edges only once per pair of nodes
    for (let i = 0; i < generatedNodes.length; i++) {
      for (let j = i + 1; j < generatedNodes.length; j++) {
        const node = generatedNodes[i];
        const targetNode = generatedNodes[j];

        // Connect nodes if their range matches
        if (node.data.range === targetNode.data.range) {
          generatedEdges.push({
            id: `${node.id}-${targetNode.id}`,
            source: node.id,
            target: targetNode.id,
            type: 'smoothstep',
            animated: true,
            style: { stroke: 'var(--DD-edge-color)', strokeWidth: 2, fill: 'none' },
            markerEnd: 'url(#edge-circle)',
          });
        }
      }
    }

    setNodes(generatedNodes);
    setEdges(generatedEdges);
  }, [data]);

  // Explicitly define 'params' type as Connection
  const onConnect = useCallback(
    (params: Connection) =>
      setEdges((eds) =>
        addEdge({ ...params, animated: true, style: { stroke: 'var(--DD-edge-color)' } }, eds)
      ),
    []
  );

  // Fix for handleNodeClick: Use correct Node type from 'react-flow-renderer'
  const handleNodeClick = (_: React.MouseEvent, node: Node<any>) => {
    const clickedNode = data.find((item) => item.IP === node.id);
    setSelectedNode(clickedNode || null);
  };

  return (
    <ReactFlowProvider>
      <div style={{ height: 600, position: 'relative' }}>
        <ReactFlow
          nodes={nodes}
          edges={edges}
          onNodesChange={onNodesChange}
          onEdgesChange={onEdgesChange}
          onConnect={onConnect}
          nodeTypes={nodeTypes}
          edgeTypes={edgeTypes}
          onNodeClick={handleNodeClick}
          snapToGrid={true}
          snapGrid={[15, 15]}
          fitView
        >
          <Controls showInteractive={false} />
          <Background color="var(--DD-bg-color)" gap={16} />
        </ReactFlow>
        <NodeDetailsDrawer selectedNode={selectedNode} onClose={() => setSelectedNode(null)} />
      </div>
    </ReactFlowProvider>
  );
};

export default DataDiagram;
