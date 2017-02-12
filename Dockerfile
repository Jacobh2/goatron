FROM golang:1.7.5

WORKDIR /go/src/app

COPY . .

ENV PORT 8080
ENV GOBIN /go/bin

RUN go get -d

RUN go install ./main/main.go

CMD ["main"]

