.PHONY: build server

build:
	cd game && env GOOS=js GOARCH=wasm go build -o ../web/build/game.wasm github.com/jacovanc/airconsole-ebiten/game

server:
	cd web && php -S 0.0.0.0:8000
