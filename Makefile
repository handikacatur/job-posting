migrate-compile:
	@echo ">> Building Binart migration..."
	go build -o build/migrate ./cmd/migrate/

migrate-up:
	$(MAKE) migrate-compile
	@echo ">> Starting migrate..."
	@./build/migrate -cmd=up & echo $$! > $@;

migrate-down:
	$(MAKE) migrate-compile
	@echo ">> Starting to down migration..."
	@./build/migrate -cmd=down & echo $$! > $@;

run:
	go run cmd/http/main.go