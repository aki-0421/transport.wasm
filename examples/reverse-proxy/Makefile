.PHONY: docker
docker:
	docker compose build && docker compose up || true

.PHONY: dev
dev:
	docker compose exec app bash -c 'wrangler dev --port 8976'

.PHONY: build
build:
	rm -rf build
	mkdir build
	cp js/* build
	tinygo build -o ./build/app.wasm -target=wasm -opt=z ./main.go

.PHONY: publish
publish:
	docker compose exec app bash -c 'wrangler publish'