FROM golang:1.17
ENV GO111MODULE=on

WORKDIR /go/src/bello-auth
COPY . .

RUN go build -o dist/bello-app cmd/bello-app-auth-server/main.go

CMD ["dist/bello-app"]
