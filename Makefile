test:
	export `cat .env` && go test

run:
	export `cat .env` && go run main.go
