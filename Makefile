test:
	cd cmd/rest_api && export `cat .env` && go test

run:
	cd cmd/rest_api && export `cat .env` && go run main.go

run_grpc:
	cd cmd/grpc && export `cat .env` && go run main.go
