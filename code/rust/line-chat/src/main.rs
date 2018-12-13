extern crate tokio;
#[macro_use]
extern crate futures;
extern crate bytes;

use tokio::io;
use tokio::net::{TcpListener, TcpStream};
use tokio::prelude::*;
use futures::sync::mpsc;
use futures::future::{self, Either};
use bytes::{BytesMut, Bytes, BufMut};

use std::collections::HashMap;
use std::net::SocketAddr;
use std::sync::{Arc, Mutex};

type Tx = mpsc::UnboundedSender<Bytes>;
type Rx = mpsc::UnboundedReceiver<Bytes>;

struct Shared {
    peers: HashMap<SocketAddr, Tx>,
}

fn main() {
    const ADDR_STR: &str = "127.0.0.1:12345";

    let addr = ADDR_STR.parse().unwrap();
    let listener = TcpListener::bind(&addr).unwrap();
    let state = Arc::new(Mutex::new(Shared::new()));

    let server = listener.incoming().for_each(move |socket| {
        process(socket, state.clone());
        Ok(())
    })
    .map_err(|err| {
        println!("accept running on {}", ADDR_STR);
    });

    tokio::run(server);
}

fn process(socket: TcpStream, state: Arc<Mutex<Shared>>) {
    let task = unimplemented!();

    tokio::spawn(task);
}