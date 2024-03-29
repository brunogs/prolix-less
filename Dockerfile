FROM golang:1.22.1-alpine3.19 AS build
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o app .
FROM alpine:3.19.1
WORKDIR /root/
COPY --from=build /app/app .
COPY web/ web/
EXPOSE 3000

CMD ["./app"]
