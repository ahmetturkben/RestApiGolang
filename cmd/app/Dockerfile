FROM golang:1.16.5

ARG BRANCH=main

RUN git clone –single-branch –branch ${BRANCH} https://github.com/ahmetturkben/RestApiGolang.git

WORKDIR /restapi/golang/

RUN go build -o /cmd/app main.go

EXPOSE 8080

ENTRYPOINT /server