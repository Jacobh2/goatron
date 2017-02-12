FROM golang:1.7.5

WORKDIR /go/src/app

COPY . .

# Use port 8080
ENV PORT 8080

# Set the GOBIN env to be able to install
ENV GOBIN /go/bin

# Envs for MySQL
ENV MYSQL_USER todo
ENV MYSQL_PASSWORD 1234
ENV MYSQL_DB_NAME todo
ENV MYSQL_ADDRESS "mysql:3306"

RUN go get -d

RUN go install ./main/main.go

CMD ["main"]

