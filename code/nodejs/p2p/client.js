const dgram = require('dgram')
const OS = require('os')
const println = console.log

const host = process.argv[2] || '127.0.0.1'
const port = parseInt(process.argv[3], 10) || 3333

async function main() {
	const client = dgram.createSocket('udp4')

	process.stdin.resume()
	process.stdin.on('data', (data) => {
		client.send(data, 0, data.length, port, host)
	})

	client.on('message', (message) => {
		const msg = message.toString('utf8').replace(OS.EOL, '')
		println('received message from UDP server:', msg)
	})

	const message = new Buffer('hello UDP, from holy crap.')
	await client.send(message, 0, message.length, port, host)

}

main()
