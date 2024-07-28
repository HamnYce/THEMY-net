import {
  Table,
  TableHeader,
  TableBody,
  TableCell,
  TableHead,
  TableRow,
} from "@/components/ui/table";
import { Host, Port } from "@/types/host_type";
import {
  createColumnHelper,
  flexRender,
  getCoreRowModel,
  useReactTable,
} from "@tanstack/react-table";

type Props = {
  readonly ports: Port[];
};

export default function PortsTable({ ports }: Props) {
  const columnHelper = createColumnHelper<Port>();
  const defaultColumns = [
    columnHelper.accessor("id", {
      header: "ID",
      cell: (info) => info.getValue(),
    }),
    columnHelper.accessor("protocol", {
      header: "Protocol",
      cell: (info) => info.getValue(),
    }),
    columnHelper.accessor("service.name", {
      header: "Service Name",
      cell: (info) => info.getValue(),
    }),
    columnHelper.accessor("state.state", {
      header: "Port State",
      cell: (info) => info.getValue(),
    }),
  ];

  const table = useReactTable({
    data: ports,
    columns: defaultColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <Table>
      <TableHeader>
        <TableRow>
          {table.getFlatHeaders().map((header) => (
            <TableHead key={header.id}>
              {flexRender(header.column.columnDef.header, header.getContext())}
            </TableHead>
          ))}
        </TableRow>
      </TableHeader>
      <TableBody>
        {table.getRowModel().rows.map((row) => (
          <TableRow key={row.id}>
            {row.getVisibleCells().map((cell) => (
              <TableCell key={cell.id}>
                {flexRender(cell.column.columnDef.cell, cell.getContext())}
              </TableCell>
            ))}
          </TableRow>
        ))}
      </TableBody>
    </Table>
  );
}
