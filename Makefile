build:
	@go build -o bin/pet-reminder

run: build
	@./bin/pet-reminder

test: 
	@go test -v ./...c
	