#[allow(unused_imports)]
use serde::{Deserialize, Serialize};
use std::sync::mpsc;
use std::thread;
use threadpool::ThreadPool;

const START_PORT: u16 = 1;
const END_PORT: u16 = 60_000;

pub fn get_ports(network: String, mask: u8) -> ScanResult {
    let addr: Vec<String> = network.split('/').map(|s| s.to_owned()).collect();
    let (tx_tcp, rx_tcp) = mpsc::channel();
    let (tx_udp, rx_udp) = mpsc::channel();
    let num_threads = 1000;
    let pool = ThreadPool::new(num_threads);

    for port in START_PORT..=END_PORT {
        let tx_tcp = tx_tcp.clone();
        let tx_udp = tx_udp.clone();
        let addr2 = addr.clone();

        pool.execute(move || {
            let _ = tx_tcp.send(check_tcp_addr(&addr2[0], port));
            let _ = tx_udp.send(check_udp_addr(&addr2[0], port));
        });
    }

    drop(tx_tcp);
    drop(tx_udp);

    let mut open_tcp_ports: Vec<u16> = vec![];
    let mut closed_tcp_ports: Vec<u16> = vec![];
    let mut open_udp_ports: Vec<u16> = vec![];
    let mut closed_udp_ports: Vec<u16> = vec![];

    for (status, port) in rx_tcp.iter() {
        if status == "open" {
            open_tcp_ports.push(port);
        } else {
            closed_tcp_ports.push(port);
        }
    }

    for (status, port) in rx_udp.iter() {
        if status == "open" {
            open_udp_ports.push(port);
        } else {
            closed_udp_ports.push(port);
        }
    }

    ScanResult {
        open_tcp_ports,
        closed_tcp_ports,
        open_udp_ports,
        closed_udp_ports,
    }
}

fn check_tcp_addr(ip: &str, port: u16) -> (&'static str, u16) {
    let addr = format!("{}:{}", ip, port);
    match std::net::TcpStream::connect_timeout(
        &addr.parse().unwrap(),
        std::time::Duration::from_secs(1),
    ) {
        Ok(_) => ("open", port),
        Err(_) => ("closed", port),
    }
}

fn check_udp_addr(ip: &str, port: u16) -> (&'static str, u16) {
    let addr = format!("{}:{}", ip, port);
    let socket = std::net::UdpSocket::bind("0.0.0.0:0").expect("Failed to bind socket");
    socket
        .set_read_timeout(Some(std::time::Duration::from_secs(1)))
        .expect("Failed to set read timeout");

    match socket.send_to(&[0], &addr) {
        Ok(_) => match socket.recv_from(&mut [0; 1]) {
            Ok(_) => ("open", port),
            Err(_) => ("closed", port),
        },
        Err(_) => ("closed", port),
    }
}

#[derive(Debug, Serialize, Deserialize)]
pub struct Host {
    host: String,
    mac: String,
    vendor: String,
    hostname: String,
    open_ports: Vec<u16>,
    closed_ports: Vec<u16>,
}

#[derive(Serialize, Deserialize, Debug)]
pub struct ScanResult {
    open_tcp_ports: Vec<u16>,
    closed_tcp_ports: Vec<u16>,
    open_udp_ports: Vec<u16>,
    closed_udp_ports: Vec<u16>,
}
