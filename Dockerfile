FROM golang:1.20-alpine


ENV PROJECT_DIR=/app \
    GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR /app
RUN mkdir "/build"
COPY ./api . 
RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon
RUN go mod tidy
RUN go mod download
ENTRYPOINT CompileDaemon -build="go build -o /build/app" -command="/build/app"
EXPOSE 3000
# RUN go build -o api ./main.go
# CMD ["./api"]