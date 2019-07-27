run:
	go run cmd/bello-app-auth-server/main.go

generate-swagger:
	rm -rf apimodel restapi/operations && swagger generate server -m apimodel