FROM        golang:alpine
RUN         mkdir -p /app
WORKDIR     /app
COPY        . .
RUN         go mod download
RUN         go build -o app
EXPOSE 5050
ENTRYPOINT  ["./app"]

