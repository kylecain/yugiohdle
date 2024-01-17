run:
	@templ generate
	@go run cmd/app/main.go

up:
	@go run cmd/up/main.go
