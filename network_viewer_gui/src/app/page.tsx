"use client";

import { hostData } from "@/data/1";
import { useState } from "react";
import { Host } from "@/types/host_type.ts";
import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import OSMatchesTable from "@/components/host_tables/os_matches_table";

export default function Home() {
  const [selectedHostIndex, setSelectedHostIndex] = useState(0);

  // TODO: add buttons on the left panel to select between the different IPs
  return (
    <main className="flex h-screen w-full columns-2 bg-green-500">
      <div className="flex-col min-h-screen bg-red-400 w-96 p-10 list-none">
        {hostData.map((host: Host, index: number) => {
          return (
            <h4
              key={host.addresses[0].addr}
              onClick={() => setSelectedHostIndex(index)}
              className={
                "cursor-pointer" +
                (selectedHostIndex === index ? " bg-blue-500" : "")
              }
            >
              {host.addresses[0].addr}
            </h4>
          );
        })}
      </div>
      <div className="flex-col h-screen w-screen bg-blue-400 p-10">
        <OSMatchesTable host={hostData[selectedHostIndex]} />
      </div>
    </main>
  );
}
