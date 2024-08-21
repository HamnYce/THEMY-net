//This is the home page which will display detailed "stats" for all things scanned. Data can be updated here by calling a "fetch" from the DB
"use client";
import React, { useState } from "react";
import {
  ResizableHandle,
  ResizablePanel,
  ResizablePanelGroup,
} from "@/components/ui/resizable";
import OverViewOsStats from "@/features/overviewDataDisplay/overviewOsStats";

export default function Home() {
  return (
    <div>
      <h1>Network Dashboard</h1>
      <OverViewOsStats />

      <ResizablePanelGroup direction="horizontal">
        <ResizablePanel>
          {" "}
          <div className="bg-blue-500 text-white p-4">
            Tailwind CSS is working? Color will be different if it is.
          </div>
        </ResizablePanel>
        <ResizableHandle withHandle />
        <ResizablePanel>Two</ResizablePanel>
      </ResizablePanelGroup>
    </div>
  );
}
