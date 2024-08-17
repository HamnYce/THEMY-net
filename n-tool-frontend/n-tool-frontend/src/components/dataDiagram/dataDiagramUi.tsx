import React from 'react';
import {
  Handle,
  Position,
  NodeProps,
  EdgeProps,
  getBezierPath,
} from 'react-flow-renderer';
import { Tooltip, TooltipTrigger, TooltipContent, TooltipProvider } from '@/components/ui/tooltip';
import { Card, CardHeader, CardContent } from '@/components/ui/card';
import { Drawer, DrawerTrigger, DrawerContent } from '@/components/ui/drawer';

// Import Data interface from dataDiagram.tsx
import { Data } from './dataDiagram';

// Node Component
export const NodeComponent: React.FC<NodeProps> = ({ data }) => (
  <Card style={{ backgroundColor: 'var(--DD-card-bg-color)' }}>
    <CardHeader style={{ color: 'var(--DD-text-color)' }}>
      <Handle type="target" position={Position.Left} />
      {data.label}
      <Handle type="source" position={Position.Right} />
    </CardHeader>
    <CardContent style={{ color: 'var(--DD-text-color)', opacity: 0.8 }}>
      <div>
        {data.details.split(' | ')[0]} |{' '}
        <span style={getStatusStyle(data.details.split(' | ')[1])}>
          {data.details.split(' | ')[1]}
        </span>
      </div>
    </CardContent>
  </Card>
);

const getStatusStyle = (status: string) => {
  const trimmedStatus = status.trim().toLowerCase();
  if (trimmedStatus === 'online') {
    return { color: 'green', opacity: 1 };
  } else if (trimmedStatus === 'offline') {
    return { color: 'red', opacity: 0.4 };
  }
  return { color: 'var(--DD-text-color)', opacity: 0.8 };
};

// Edge Component
export const EdgeComponent: React.FC<EdgeProps> = ({
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

// Drawer Component for Selected Node Details
export const NodeDetailsDrawer: React.FC<{ selectedNode: Data | null, onClose: () => void }> = ({ selectedNode, onClose }) => (
  <Drawer open={!!selectedNode} onClose={onClose}>
    <DrawerTrigger>
      <button className="hidden"></button>
    </DrawerTrigger>
    <DrawerContent className="bg-[color:var(--DD-bg-color)] text-[color:var(--DD-text-color)] p-4">
      {selectedNode && (
        <>
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
            className="mt-4 bg-primary px-4 py-2 rounded"
            onClick={onClose}
            style={{ backgroundColor: 'var(--primary)', color: 'var(--primary-foreground)' }}
          >
            Close
          </button>
        </>
      )}
    </DrawerContent>
  </Drawer>
);

// Tooltip Component
export const CustomTooltip: React.FC<{ hostname: string, ip: string, range: string }> = ({ hostname, ip, range }) => (
  <TooltipProvider>
    <Tooltip>
      <TooltipTrigger asChild>
        <span>{`${hostname} (${ip})`}</span>
      </TooltipTrigger>
      <TooltipContent>
        <span>Subnet: {range}</span>
      </TooltipContent>
    </Tooltip>
  </TooltipProvider>
);
