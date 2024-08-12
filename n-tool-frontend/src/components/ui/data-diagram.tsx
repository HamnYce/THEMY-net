"use client";
//TODO: Clean up this messy code and fix the logic
import React, { useState, useMemo } from 'react';
import ReactFlow, {
  ReactFlowProvider,
  Background,
  Controls,
  Handle,
  Position,
  getBezierPath,
  NodeProps,
  EdgeProps,
  Node,
  Edge,
} from 'react-flow-renderer';
import {
  Drawer,
  DrawerTrigger,
  DrawerContent,
} from '@/components/ui/drawer';
import { Tooltip, TooltipTrigger, TooltipContent, TooltipProvider } from '@/components/ui/tooltip';
import { Card, CardHeader, CardContent } from '@/components/ui/card';
import { debugLog } from '@/utils/debugLogUtil';

// Define the type for your data
interface Data {
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

const NodeComponent: React.FC<NodeProps> = ({ data }) => (
  <Card style={{ backgroundColor: 'var(--DD-card-bg-color)' }}>
    <CardHeader>
      <Handle type="target" position={Position.Left} />
      {data.label}
      <Handle type="source" position={Position.Right} />
    </CardHeader>
    <CardContent>
      <div>{data.details}</div>
    </CardContent>
  </Card>
);

const EdgeComponent: React.FC<EdgeProps> = ({
  id,
  sourceX,
  sourceY,
  targetX,
  targetY,
  sourcePosition,
  targetPosition,
  markerEnd,
}) => {
  const edgePath = getBezierPath({
    sourceX,
    sourceY,
    sourcePosition,
    targetX,
    targetY,
    targetPosition,
  });

  return (
    <path
      id={id}
      className="react-flow__edge-path"
      d={edgePath}
      markerEnd={markerEnd}
      style={{ stroke: 'var(--DD-edge-color)', strokeWidth: 2, fill: 'none' }}
    />
  );
};

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
        label: (
          <TooltipProvider>
            <Tooltip>
              <TooltipTrigger asChild>
                <span>{`${node.Hostname} (${node.IP})`}</span>
              </TooltipTrigger>
              <TooltipContent>
                <span>Subnet: {node.Range}</span>
              </TooltipContent>
            </Tooltip>
          </TooltipProvider>
        ),
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
        {selectedNode && (
          <Drawer open={true} onClose={() => setSelectedNode(null)}>
            <DrawerTrigger>
              <button className="hidden"></button>
            </DrawerTrigger>
            <DrawerContent className="bg-[color:var(--DD-bg-color)] text-[color:var(--DD-text-color)] p-4">
              <h2 className="text-xl font-bold mb-4">
                {selectedNode.Hostname} ({selectedNode.IP})
              </h2>
              <ul className="list-disc pl-5 space-y-1">
                <li>Status: {selectedNode.Status}</li>
                <li>Exposure: {selectedNode.Exposure}</li>
                <li>Internet Access: {selectedNode.InternetAccess}</li>
                <li>OS: {selectedNode.OS} {selectedNode.OSVersion}</li>
                <li>Usage: {selectedNode.Usage}</li>
                <li>Location: {selectedNode.Location}</li>
                <li>Owners: {Array.isArray(selectedNode.Owners) ? selectedNode.Owners.join(', ') : selectedNode.Owners}</li>
                <li>Dependencies: {Array.isArray(selectedNode.Dependencies) ? selectedNode.Dependencies.join(', ') : selectedNode.Dependencies}</li>
                <li>Created At: {selectedNode.CreatedAt}</li>
                <li>Created By: {selectedNode.CreatedBy}</li>
                <li>Recorded At: {selectedNode.RecordedAt}</li>
                <li>Host Type: {selectedNode.HostType}</li>
                <li>Exposed Services: {Array.isArray(selectedNode.ExposedServices) ? selectedNode.ExposedServices.join(', ') : selectedNode.ExposedServices}</li>
                <li>CPU: {selectedNode.CPU}</li>
                <li>RAM: {selectedNode.RAM} GB</li>
                <li>Storage: {selectedNode.Storage} GB</li>
                <li>Open Ports: {Array.isArray(selectedNode.OpenPorts) ? selectedNode.OpenPorts.join(', ') : selectedNode.OpenPorts}</li>
                <li>Range: {selectedNode.Range}</li>
              </ul>
              <button
                className="mt-4 bg-primary text-primary-foreground px-4 py-2 rounded"
                onClick={() => setSelectedNode(null)}
                style={{ backgroundColor: 'var(--primary)', color: 'var(--primary-foreground)' }}
              >
                Close
              </button>
            </DrawerContent>
          </Drawer>
        )}
      </div>
    </ReactFlowProvider>
  );
};

export default DataDiagram;
