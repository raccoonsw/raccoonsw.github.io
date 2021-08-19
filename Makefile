test:
	export `cat .env` && cd cmd/restapi && go test

run:
	export `cat .env` && go run cmd/restapi/main.go
