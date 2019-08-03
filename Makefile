run:
	go run cmd/bello-app-auth-server/main.go

test-env:
	docker-compose -f docker-compose.yml up -d

generate-swagger:
	rm -rf apimodel restapi/operations && swagger generate server -m apimodel