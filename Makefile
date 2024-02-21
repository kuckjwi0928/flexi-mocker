.PHONY: clean
clean:
	@echo "Cleaning up..."
	@rm -f flexi-mocker

.PHONY: build
build: clean
	@echo "Building application..."
	@go build -o flexi-mocker ./cmd/main.go

.PHONY: run
run:
	cp config.toml.template config.toml
	@go run ./cmd/main.go
