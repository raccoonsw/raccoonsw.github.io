test:
	cd cmd/rest_api && export `cat .env` && go test

run:
	cd cmd/rest_api && export `cat .env` && go run main.go

run-grpc:
	cd cmd/grpc && export `cat .env` && go run main.go

docker-up:
	docker-compose -f docker-compose.ci.yml up -d

docker-down:
	docker-compose -f docker-compose.ci.yml down