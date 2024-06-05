run: 
	go run cmd/*.go

instal_sqlc :
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

install_goose :
	go install github.com/pressly/goose/v3/cmd/goose@latest


goose_up:
	cd sql/schema && goose postgres postgres://postgres:root@localhost:5432/attempt2 up

goose_down:
	cd sql/schema && goose postgres postgres://postgres:root@localhost:5432/attempt2 down

sqlc:
	sqlc generate