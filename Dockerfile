FROM golang:1.23

WORKDIR /vesperis_proxy

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -o /vesperis_proxy_app

EXPOSE 25565

CMD ["/vesperis_proxy_app"]