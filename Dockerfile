FROM golang

RUN mkdir /api

COPY . /api

WORKDIR  /api

RUN go env -w GO111MODULE=on

RUN go mod init Project

RUN go get github.com/spf13/viper
RUN go get github.com/gorilla/mux

RUN go env -w GO111MODULE=auto

RUN go build -o app .

CMD ["/api/app"]
