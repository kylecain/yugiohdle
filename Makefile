run:
	@npx npx tailwindcss -i ./dist/main.css -o ./dist/tailwind.css
	@templ generate
	@go run cmd/app/main.go

up:
	@go run cmd/up/main.go
