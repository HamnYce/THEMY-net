"use client"
import React, { useState } from 'react';
import {
  ResizableHandle,
  ResizablePanel,
  ResizablePanelGroup,
} from "@/components/ui/resizable"
 
export default function Home() {


  return (

    <div>
    <div>Home</div>
   
  
    <ResizablePanelGroup direction="horizontal">
       
    <ResizablePanel> <div className="bg-blue-500 text-white p-4">
      Tailwind CSS is working? Color will be different if it is.
    </div></ResizablePanel>
    <ResizableHandle withHandle />
    <ResizablePanel>Two</ResizablePanel>
  </ResizablePanelGroup>
  </div>
  );
}
