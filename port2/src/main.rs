mod ip;
mod scan;

use crate::scan::get_ports;
use actix_web::{post, App, HttpResponse, HttpServer, Responder};
use serde::{Deserialize, Serialize};

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| App::new().service(scanner))
        .bind(("100.112.234.55", 8080))?
        .run()
        .await
}

#[post("/")]
async fn scanner(req_body: String) -> impl Responder {
    let ip: Ip = serde_json::from_str(&req_body).unwrap();
    let addr = ip.addr;
    let mask = ip.mask.trim().parse::<u8>().unwrap();
    let scan_result = get_ports(addr, mask);
    let json = serde_json::to_string(&scan_result).unwrap();
    HttpResponse::Ok().body(json)
}

#[derive(Serialize, Deserialize, Debug)]
struct Ip {
    addr: String,
    mask: String,
}
