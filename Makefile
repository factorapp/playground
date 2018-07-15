.PHONY: wasm
wasm:
	GOARCH=wasm GOOS=js go build -o app/example.wasm ./client

.PHONY: server-app
server-app:
	go build -o server-app ./server

.PHONY: make-parse
make-parse:
	cd cmd/parse && go install && cd ../../

.PHONY: parse
parse:
	parse

.PHONY: run
run: make-parse parse wasm server-app
	./server-app
