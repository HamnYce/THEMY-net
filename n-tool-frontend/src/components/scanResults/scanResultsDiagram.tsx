"use client";

import React, { useEffect, useState } from 'react';
import DataDiagram from '@/components/ui/data-diagram'; // Ensure the path is correct
import { fetchDataFromJson } from '@/utils/resultsUtils';
import { formatCellValue } from '@/utils/resultsParser';

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

const ScanResultsDiagram: React.FC = () => {
  const [data, setData] = useState<Data[]>([]);
  const [error, setError] = useState<string | null>(null);
  const isTestMode = process.env.NEXT_PUBLIC_IS_TEST_MODE === 'true';
  const jsonFilePath = isTestMode 
    ? process.env.NEXT_PUBLIC_TEST_JSON_FILE_PATH
    : process.env.NEXT_PUBLIC_JSON_FILE_PATH;

  useEffect(() => {
    const getData = async () => {
      if (!jsonFilePath) {
        setError('Error fetching data: No JSON file path provided');
        return;
      }

      const { data, error } = await fetchDataFromJson(jsonFilePath);
      if (data) {
        const formattedData = data.map((item: any) => {
          const formattedItem: any = {};
          Object.keys(item).forEach((key) => {
            formattedItem[key] = formatCellValue(item[key]);
          });
          return formattedItem;
        });
        setData(formattedData);
      }
      setError(error);
    };
    getData();
  }, [jsonFilePath]);

  if (error) {
    return <div>{error}</div>;
  }

  return (
    <div>
      <h2>Diagram View</h2>
      {data.length > 0 ? <DataDiagram data={data} /> : <div>Loading...</div>}
    </div>
  );
};

export default ScanResultsDiagram;
