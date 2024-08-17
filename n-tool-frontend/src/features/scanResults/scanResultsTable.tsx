"use client";

import React, { useEffect, useState } from 'react';
import DataTable from '@/components/ui/data-table'; 
import { fetchDataFromJson } from '@/utils/resultsUtils';
import { formatCellValue } from '@/utils/resultsParser';
import { ColumnDef } from '@tanstack/react-table';


interface Data {
  [key: string]: any; 
}

const ScanResultsTable: React.FC = () => {
  const [data, setData] = useState<Data[]>([]);
  const [error, setError] = useState<string | null>(null);
  const columns = process.env.NEXT_PUBLIC_JSON_COLUMNS ? process.env.NEXT_PUBLIC_JSON_COLUMNS.split(',') : [];
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
      setData(data);
      setError(error);
    };
    getData();
  }, [jsonFilePath]);

  if (error) {
    return <div>{error}</div>;
  }

  // Define the column definitions with proper types
  const columnDefs: ColumnDef<Data>[] = columns.map(col => ({
    accessorKey: col,
    header: col,
    cell: ({ row }) => formatCellValue(row.getValue(col)),
  }));

  return (
    <DataTable data={data} columns={columnDefs} />
  );
};

export default ScanResultsTable;
