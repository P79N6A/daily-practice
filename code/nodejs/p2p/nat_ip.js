const os = require('os')
const println = console.log

function main() {
	const ifaces = os.networkInterfaces()
  Object.keys(ifaces).forEach(name => {
    const nc = ifaces[name][0]

    if (name === 'en1') {
      println(name, nc)
    }
  })
}

main()
