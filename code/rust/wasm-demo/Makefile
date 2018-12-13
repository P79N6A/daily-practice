SHELL := /bin/bash

all:
	rm -rf site/*.js site/*.wasm
	cargo build --target=wasm32-unknown-unknown --release
	mkdir -p site
	wasm-bindgen target/wasm32-unknown-unknown/release/wasm_demo.wasm --out-dir ./site
	\cp -rf ./index.html site/
	\cp -rf ./index.js site/
	\cp -rf ./package.json site/
	\cp -rf ./webpack.config.js site/
