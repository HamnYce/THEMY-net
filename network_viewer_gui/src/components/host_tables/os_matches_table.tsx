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

export default function OSMatchesTable({ host }: Props) {
  return (
    <Table>
      <TableCaption>Os Matches</TableCaption>
      <TableHeader>
        <TableRow>
          <TableHead>Name</TableHead>
          <TableHead>Accuracy (%)</TableHead>
          <TableHead>Vendor</TableHead>
          <TableHead>Generation</TableHead>
          <TableHead>Type</TableHead>
          <TableHead>Accuracy</TableHead>
          <TableHead>Os Family</TableHead>
          <TableHead>CPEs</TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        {host.os.os_matches?.map((osMatch) => {
          return (
            <TableRow key={osMatch.name}>
              <TableCell>{osMatch.name}</TableCell>
              <TableCell>{osMatch.accuracy}</TableCell>
              <TableCell>{osMatch.os_classes[0].vendor}</TableCell>
              <TableCell>{osMatch.os_classes[0].os_generation}</TableCell>
              <TableCell>{osMatch.os_classes[0].type}</TableCell>
              <TableCell>{osMatch.os_classes[0].accuracy}</TableCell>
              <TableCell>{osMatch.os_classes[0].os_family}</TableCell>
              <TableCell>{osMatch.os_classes[0].cpes.join(", ")}</TableCell>
            </TableRow>
          );
        })}
      </TableBody>
    </Table>
  );
}
