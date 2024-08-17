use core::time::Duration;
use std::net::{IpAddr, Ipv4Addr, SocketAddr, TcpStream, UdpSocket};

pub fn check_tcp_addr<'a>(ip: &str, port: u16) -> (&'a str, u16) {
    let ip = parse_ip(ip);
    let addr = SocketAddr::new(IpAddr::V4(Ipv4Addr::new(ip[0], ip[1], ip[2], ip[3])), port);
    match TcpStream::connect_timeout(&addr, Duration::from_secs(5)) {
        Ok(_) => ("open", port),
        Err(_) => ("closed", port),
    }
}

pub fn check_udp_addr<'a>(ip: &str, port: u16) -> (&'a str, u16) {
    let ip = parse_ip(ip);
    let addr = SocketAddr::new(IpAddr::V4(Ipv4Addr::new(ip[0], ip[1], ip[2], ip[3])), port);
    let socket = UdpSocket::bind(addr).unwrap();
    match socket.send_to(&[0], addr) {
        Ok(_) => ("open", port),
        Err(_) => ("closed", port),
    }
}

fn parse_ip(ip: &str) -> Vec<u8> {
    let ip = ip
        .split('.')
        .map(|s| s.parse::<u8>().unwrap())
        .collect::<Vec<_>>();
    ip
}
