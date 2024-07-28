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

export default function HostNameTable({ host }: Props) {
  return (
    <Table>
      <TableCaption>Host Names</TableCaption>
      <TableHeader>
        <TableRow>
          <TableHead>Hostname</TableHead>
          <TableHead>Type</TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        {host.hostnames.map((hostname) => {
          return (
            <TableRow key={hostname.name}>
              <TableCell>{hostname.name}</TableCell>
              <TableCell>{hostname.type}</TableCell>
            </TableRow>
          );
        })}
      </TableBody>
    </Table>
  );
}
