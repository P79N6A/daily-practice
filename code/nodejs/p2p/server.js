const dgram = require('dgram')
const OS = require('os')
const println = console.log
const port = 3333

async function main() {
	const server = dgram.createSocket('udp4')

	server.on('message', (message, rinfo) => {
        const msg = message.toString('utf8').replace(OS.EOL, '')
		println(`received message: ${msg}, from:${rinfo.address}:${rinfo.port}`)
		server.send(message, 0, message.length, rinfo.port, rinfo.address)
	})

	server.on('listening', () => {
		const address = server.address()
		println(`listening on: ${address.address}:${address.port}`)
	})

	server.bind(port)
}

main()
