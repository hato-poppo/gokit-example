FROM golang:1.16.2

WORKDIR /go/src/hot_reload_docker
COPY . .
ENV GO111MODULE=on

RUN apt-get update
RUN apt-get install -y \
  git \
  vim

RUN go mod init
RUN go get github.com/pilu/fresh
RUN go get github.com/go-kit/kit/transport/http

# CMD ["fresh"]

EXPOSE 8080
CMD ["go", "run", "main.go"]