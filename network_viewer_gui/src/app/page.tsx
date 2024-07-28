"use client";

import { hostData } from "@/data/1";
import { useState } from "react";
import { Host } from "@/types/host_type.ts";
import OSMatchesTable from "@/components/host_tables/os_matches_table";
import HostNamesTable from "@/components/host_tables/hostnames_table";
import PortsTable from "@/components/host_tables/ports_table";

export default function Home() {
  const [selectedHostIndex, setSelectedHostIndex] = useState(0);

  return (
    <main className="flex h-screen w-full columns-2 dark ">
      <div className="min-h-screen bg-background w-96 p-10 list-none overflow-auto">
        {hostData.map((host: Host, index: number) => {
          return (
            <h4
              key={host.addresses[0].addr}
              onClick={() => setSelectedHostIndex(index)}
              className={
                "cursor-pointer rounded" +
                " " +
                (selectedHostIndex === index
                  ? "text-foreground"
                  : "text-muted-foreground")
              }
            >
              {host.addresses[0].addr}
            </h4>
          );
        })}
      </div>
      <div className="h-screen w-screen bg-background text-foreground p-10 overflow-auto">
        <HostNamesTable hostnames={hostData[selectedHostIndex].hostnames} />
        <PortsTable host={hostData[selectedHostIndex]} />
        <OSMatchesTable host={hostData[selectedHostIndex]} />
      </div>
    </main>
  );
}
