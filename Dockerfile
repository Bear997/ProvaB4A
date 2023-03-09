FROM golang:1.20-alpine

WORKDIR /app

COPY ./api . 
RUN go mod tidy
RUN go mod download
EXPOSE 3000
RUN go build -o api ./main.go
CMD ["./api"]