FROM golang:latest AS build

WORKDIR /app
COPY go.mod .
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ipcheck .

FROM alpine:latest

WORKDIR /app
COPY --from=build /app/ipcheck .

CMD ["./ipcheck"]
