"use client";

import React, { useState } from 'react';
import ScanResultsTable from './scanResultsTable';
import ScanResultsDiagram from './scanResultsDiagram';
import {
  NavigationMenu,
  NavigationMenuList,
  NavigationMenuItem,
  NavigationMenuTrigger,
} from '@/components/ui/navigation-menu';

/**
 * ScanResults Component
 * Includes a navigation menu to switch between table and diagram views.
 */
const ScanResults: React.FC = () => {
  const [view, setView] = useState('table'); // State to manage current view

  return (
    <div>
      <NavigationMenu>
        <NavigationMenuList>
          <NavigationMenuItem>
            <NavigationMenuTrigger onClick={() => setView('table')}>Table</NavigationMenuTrigger>
          </NavigationMenuItem>
          <NavigationMenuItem>
            <NavigationMenuTrigger onClick={() => setView('diagram')}>Diagram</NavigationMenuTrigger>
          </NavigationMenuItem>
        </NavigationMenuList>
      </NavigationMenu>
      <div className="mt-4">
        {view === 'table' ? <ScanResultsTable /> : <ScanResultsDiagram />}
      </div>
    </div>
  );
};

export default ScanResults;
