//TODO: Replace hardcoded values with values fetched from actual data.

"use client";

import * as React from "react";
import { Card, CardHeader, CardFooter, CardContent, CardTitle } from "@/components/ui/card";

interface StatsProps {
  title: string;
  value: string;
  timestampLabel: string;
  timestamp: string;
}

const Stats: React.FC<StatsProps> = ({ title, value, timestampLabel, timestamp }) => {
  return (
    <Card className="bg-card text-card-foreground shadow-sm rounded-lg p-4">
      <CardHeader>
        <CardTitle className="text-xl font-semibold">{title}</CardTitle>
      </CardHeader>
      <CardContent>
        <p className="text-4xl font-bold">{value}</p>
      </CardContent>
      <CardFooter>
        <div className="flex flex-col">
          <span className="text-sm font-medium">{timestampLabel}</span>
          <p className="text-sm text-gray-500">{timestamp}</p>
        </div>
      </CardFooter>
    </Card>
  );
};

const OverViewOsStats: React.FC = () => {
  const stats = [
    { title: "Linux", value: "17", timestampLabel: "Last updated on:", timestamp: "2024-08-18 14:22:01" },
    { title: "Windows", value: "5", timestampLabel: "Last updated on:", timestamp: "2024-08-17 10:45:12" },
    { title: "MacOS", value: "2", timestampLabel: "Last updated on:", timestamp: "2024-08-19 08:30:50" }
  ];

  return (
    <div className="container mx-auto my-8">
      <h2 className="text-3xl font-bold mb-6">Server OS Count</h2>
      <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
        {stats.map((stat, index) => (
          <Stats
            key={index}
            title={stat.title}
            value={stat.value}
            timestampLabel={stat.timestampLabel}
            timestamp={stat.timestamp}
          />
        ))}
      </div>
    </div>
  );
};

export default OverViewOsStats;
