
run: bin/marvel
	@echo ---- Running marvel ----
	./bin/marvel

clean:
	@echo ---- Clean up ----
	rm ./bin -fvr

build: bin/marvel
bin/marvel:
	@echo ---- Building marvel binary ----
	go build -o bin/marvel ./cmd/marvel

test:
	@echo ---- Running test for gateway ----
	go test ./internal/gateway

	@echo ---- Running test for marvel ----
	go test ./internal/marvel

	@echo ---- Running test for memorydb ----
	go test ./internal/memorydb

	@echo ---- Running test for scraper ----
	go test ./internal/scraper

dock-build:
	@echo ---- Building docker image ----
	docker-compose build

dock-up:
	@echo ---- Running docker service ----
	docker-compose up -d

dock-ps:
	@echo ---- Process list ----
	docker-compose ps

dock-logs:
	@echo ---- Docker logs ----
	docker-compose logs -f

dock-down:
	@echo ---- Stop docker service ----
	docker-compose down
