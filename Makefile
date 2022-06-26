run:
	docker-compose --env-file .env up -d


doc-gen:
	swag init -g cmd/s3corp-golang-fresher/main.go -o api
