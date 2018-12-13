use tokio;
use tokio::io;
use tokio::net::TcpListener;
use tokio::prelude::*;

pub fn serve(addr: &str) {
    let server_addr = addr.parse().unwrap();
    let tcp = TcpListener::bind(&server_addr).unwrap();

    let server = tcp.incoming()
        .map_err(|e| eprintln!("accept failed = {:?}", e))
        .for_each(|sock| {
            let (reader, writer) = sock.split();

            let bytes_copied = io::copy(reader, writer);
            
            let handle_conn = bytes_copied
                .then(|res| {
                    println!("wrote message; success={:?}", res.is_ok());
                    Ok(())
                });
            
            tokio::spawn(handle_conn)
        });

    tokio::run(server);
}