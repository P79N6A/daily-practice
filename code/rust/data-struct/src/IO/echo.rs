use std::net::TcpListener;
use std::io;
use std::thread::spawn;

pub fn serve(addr: &str) -> io::Result<()> {
    let listener = TcpListener::bind(addr)?;
    println!("listening on {}", addr);

    loop {
        let (mut stream, addr) = listener.accept()?;
        println!("connection received from {}", addr);

        let mut write_stream = stream.try_clone()?;
        spawn(move || {
            io::copy(&mut stream, &mut write_stream)
                .expect("error in client thread: ");
            println!("connection closed");
        });
    }
}