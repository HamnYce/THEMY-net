import {
  Table,
  TableHeader,
  TableBody,
  TableCaption,
  TableCell,
  TableHead,
  TableRow,
} from "@/components/ui/table";
import { Host } from "@/types/host_type";

type Props = {
  readonly host: Host;
};
export default function PortsTable({ host }: Props) {
  return (
    <Table>
      <TableCaption>Ports Information</TableCaption>
      <TableHeader>
        <TableRow>
          <TableHead>ID</TableHead>
          <TableHead>Protocol</TableHead>
          <TableHead>Service Name</TableHead>
          <TableHead>Port State</TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        {host.ports.map((port) => {
          return (
            <TableRow key={"" + host.addresses[0].addr + port.id}>
              <TableCell>{port.id}</TableCell>
              <TableCell>{port.protocol}</TableCell>
              <TableCell>{port.service.name}</TableCell>
              <TableCell>{port.state.state}</TableCell>
            </TableRow>
          );
        })}
      </TableBody>
    </Table>
  );
}
