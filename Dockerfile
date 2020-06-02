FROM golang:1.13
LABEL maintainer="indecember"

WORKDIR /go/src/groupie-tracker
COPY . .

EXPOSE 8181

RUN go build -o main .

CMD ["./main"]