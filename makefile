up:
	docker-compose up --build -d

down:
	docker-compose down

backend run:
	cd ./backend && go run main.go

backend build:
	cd ./backend && go build

backend build && run:
	cls && cd ./backend && go build && go-wallet-api.exe

tidy:
	cd ./backend && go mod tidy