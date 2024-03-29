FROM golang:1.21

WORKDIR /root

COPY go.mod .

RUN go mod download

COPY . .

RUN go build -o ./output .

EXPOSE 80

CMD ["./output"]
