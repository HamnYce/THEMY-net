import {
  Table,
  TableHeader,
  TableBody,
  TableCaption,
  TableCell,
  TableHead,
  TableRow,
} from "@/components/ui/table";
import { Host, Hostname } from "@/types/host_type";
import {
  createColumnHelper,
  flexRender,
  getCoreRowModel,
  useReactTable,
} from "@tanstack/react-table";

type Props = {
  readonly hostnames: Hostname[];
};

export default function HostNamesTable({ hostnames }: Props) {
  const columnHelper = createColumnHelper<Hostname>();

  const defaultColumns = [
    columnHelper.accessor("name", {
      header: "Hostname",
      cell: (info) => info.getValue(),
    }),
    columnHelper.accessor("type", {
      header: "Type",
      cell: (info) => info.getValue(),
    }),
  ];

  const table = useReactTable({
    columns: defaultColumns,
    data: hostnames,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <Table>
      <TableCaption>Host Names</TableCaption>
      <TableHeader>
        <TableRow>
          {table.getFlatHeaders().map((header) => (
            <TableHead key={header.id}>{header.id}</TableHead>
          ))}
        </TableRow>
      </TableHeader>
      <TableBody>
        {table.getRowModel().rows.map((row) => {
          return (
            <TableRow key={row.id}>
              {row.getVisibleCells().map((cell) => (
                <TableCell key={cell.id}>
                  {flexRender(cell.column.columnDef.cell, cell.getContext())}
                </TableCell>
              ))}
            </TableRow>
          );
        })}
      </TableBody>
    </Table>
  );
}
