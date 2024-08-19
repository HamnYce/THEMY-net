"use client";

import * as React from "react";
import {
  Card,
  CardHeader,
  CardFooter,
  CardContent,
  CardTitle,
} from "@/components/ui/card";

interface StatsProps {
  title: string;
  value: string;
  timestampLabel: string;
  timestamp: string;
}

const Stats: React.FC<StatsProps> = ({
  title,
  value,
  timestampLabel,
  timestamp,
}) => {
  return (
    <Card className=" shadow-sm rounded-lg">
      <CardHeader>
        <CardTitle>{title}</CardTitle>
      </CardHeader>
      <CardContent>
        <p className="text-2xl font-bold">{value}</p>
      </CardContent>
      <CardFooter>
        <div className="flex flex-col">
          <span className="text-sm font-medium">{timestampLabel}</span>
          <p className="text-sm">{timestamp}</p>
        </div>
      </CardFooter>
    </Card>
  );
};

export default Stats;
