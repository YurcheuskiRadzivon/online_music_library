FROM golang:1.22.9-alpine as builder
WORKDIR /online_library
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o migrate cmd/migrate/main.go
RUN go build -o app cmd/app/main.go

FROM alpine:latest
WORKDIR /online_library
COPY --from=builder /online_library/migrate ./migrate
COPY --from=builder /online_library/app ./app
COPY --from=builder /online_library/.env ./
CMD ["./migrate"]
