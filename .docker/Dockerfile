FROM golang:1.16-alpine

LABEL maintainer="Alejandro Gonzalez R <alejandrogonzalr@gmail.com>"

RUN apk update && apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server .

EXPOSE 8090

CMD ./server
