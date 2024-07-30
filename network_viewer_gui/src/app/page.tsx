"use client";

import { hostData as data1 } from "@/data/1";
import { hostData as data2 } from "@/data/2";
import { useState } from "react";
import { Host } from "@/types/host_type.ts";
import OSMatchesTable from "@/components/ui/os_matches_table";
import HostNamesTable from "@/components/ui/hostnames_table";
import PortsTable from "@/components/ui/ports_table";

export default function Home() {
  const [selectedHostIndex, setSelectedHostIndex] = useState(0);
  const [hostData, _] = useState([...data1, ...data2]);
  const [seeMoreDetails, setSeeMoreDetails] = useState(false);

  return (
    <main className="flex h-screen w-full columns-2 dark ">
      <div className="min-h-screen bg-background w-96 p-10 list-none overflow-auto">
        {hostData.map((host: Host, index: number) => {
          return (
            <div key={host.addresses[0].addr} className="flex flex-row">
              <h4
                onClick={() => {
                  setSelectedHostIndex(index);
                  setSeeMoreDetails(false);
                }}
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
              {selectedHostIndex === index && (
                <>
                  <div className="w-2"></div>
                  <p
                    className={
                      "text-xs flex flex-col justify-center cursor-pointer " +
                      (seeMoreDetails
                        ? "text-foreground"
                        : "text-muted-foreground")
                    }
                    onClick={() => setSeeMoreDetails(!seeMoreDetails)}
                  >
                    see {seeMoreDetails ? "less" : "more"}
                  </p>
                </>
              )}
            </div>
          );
        })}
      </div>
      <div className="h-screen w-screen bg-background text-foreground p-10 overflow-auto">
        {seeMoreDetails ? (
          <MoreDetailsPage host={hostData[selectedHostIndex]} />
        ) : (
          <OverView></OverView>
        )}
      </div>
    </main>
  );
}

function OverView() {
  return <h1>Hello world</h1>;
}

function MoreDetailsPage({ host }: { host: Host }) {
  return (
    <>
      <HostNamesTable hostnames={host.hostnames ?? []} />
      <PortsTable ports={host.ports} />
      {/* // TODO: convert osmatches table to tanstack table */}
      <OSMatchesTable host={host} />
      {/* // TODO: add addresses table */}
    </>
  );
}
