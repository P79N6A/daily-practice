const wasm = import("./wasm_demo")

export function duang(name) {
  alert(name)
}

function main(wasm) {
  wasm.greet("Jack")
}

// passed in Safari
wasm.then(main)