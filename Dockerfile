FROM golang:1.19

WORKDIR /go-app

COPY go.mod . go.sum ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-go-app

EXPOSE 8080

CMD [ "/docker-go-app" ]