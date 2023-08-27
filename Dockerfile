FROM golang:latest
WORKDIR /app
COPY . ./
RUN GOOS=linux go build -o /sandwichyum-docker
EXPOSE 3200
ENTRYPOINT ["/sandwichyum-docker"]