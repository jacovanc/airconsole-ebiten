.PHONY: build server

build:
	cd game && env GOOS=js GOARCH=wasm go build -o ../web/build/game.wasm .

server:
	cd web && php -S 0.0.0.0:8000
