FROM golang:1.7.5

WORKDIR /go/src/app

COPY . .

# Envs for the app
ENV PORT 8080
ENV BASE_URL "/api/v1/todos"
ENV ADDRESS "0.0.0.0"

# Set the GOBIN env to be able to install
ENV GOBIN /go/bin

# Envs for MySQL
ENV MYSQL_USER todo
ENV MYSQL_PASSWORD 1234
ENV MYSQL_DB_NAME todo
# Dont use with compose
#ENV MYSQL_ADDRESS "mysql:3306"

RUN go get -d

RUN go install ./main/main.go

CMD ["main"]

