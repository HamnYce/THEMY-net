export interface Host {
  distance: Distance;
  end_time: number;
  ip_id_sequence: IpIdSequence;
  os: Os;
  start_time: number;
  status: Status;
  tcp_sequence: TcpSequence;
  tcp_ts_sequence: TcpTsSequence;
  times: Times;
  trace: Trace;
  uptime: Uptime;
  comment: string;
  addresses: Address[];
  extra_ports: ExtraPort[];
  hostnames: Hostname[];
  host_scripts: any;
  ports: Port[];
  smurfs: any;
}

export interface Distance {
  value: number;
}

export interface IpIdSequence {
  class: string;
  values: string;
}

export interface Os {
  ports_used: PortsUsed[];
  os_matches?: OsMatch[];
  os_fingerprints?: OsFingerprint[];
}

export interface PortsUsed {
  state: string;
  proto: string;
  port_id: number;
}

export interface OsMatch {
  name: string;
  accuracy: number;
  line: number;
  os_classes: OsClass[];
}

export interface OsClass {
  vendor: string;
  os_generation: string;
  type: string;
  accuracy: number;
  os_family: string;
  cpes: string[];
}

export interface OsFingerprint {
  fingerprint: string;
}

export interface Status {
  state: string;
  reason: string;
  reason_ttl: number;
}

export interface TcpSequence {
  index: number;
  difficulty: string;
  values: string;
}

export interface TcpTsSequence {
  class: string;
  values: string;
}

export interface Times {
  srtt: string;
  rttv: string;
  to: string;
}

export interface Trace {
  proto: string;
  port: number;
  hops: any;
}

export interface Uptime {
  seconds: number;
  last_boot: string;
}

export interface Address {
  addr: string;
  addr_type: string;
  vendor: string;
}

export interface ExtraPort {
  state: string;
  count: number;
  reasons: Reason[];
}

export interface Reason {
  reason: string;
  count: number;
}

export interface Hostname {
  name: string;
  type: string;
}

export interface Port {
  id: number;
  protocol: string;
  owner: Owner;
  service: Service;
  state: State;
  scripts: any;
}

export interface Owner {
  name: string;
}

export interface Service {
  device_type: string;
  extra_info: string;
  high_version: string;
  hostname: string;
  low_version: string;
  method: string;
  name: string;
  os_type: string;
  product: string;
  proto: string;
  rpc_num: string;
  service_fp: string;
  tunnel: string;
  version: string;
  configuration: number;
  cpes: any;
}

export interface State {
  state: string;
  reason: string;
  reason_ip: string;
  reason_ttl: number;
}
