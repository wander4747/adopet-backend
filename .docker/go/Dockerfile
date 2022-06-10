FROM golang:1.18-alpine

ENV GO111MODULE=on

RUN apk update && apk add --no-cache



WORKDIR /app

COPY . .

RUN go install github.com/99designs/gqlgen.yaml

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/main ./cmd/graphql/main.go
CMD ["./bin/main"]
