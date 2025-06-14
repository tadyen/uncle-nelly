build: go-wasm-build bunbuild
	@echo "Built Go and Bun projects."
	if ! [ -d ./docs ]; then mkdir ./docs; fi
	cp -r ./web/dist/* -t ./docs/

bunbuild:
	@echo "Makefile: building bun project..."
	cd ./web && bun run build

go-wasm-build: copy-wasm-exec
	@echo "Makefile: building Go wasm project..."
	GOOS=js GOARCH=wasm go build -o main.wasm

copy-wasm-exec:
	cp $(shell go env GOROOT)/lib/wasm/wasm_exec.js ./

