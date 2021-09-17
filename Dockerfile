FROM golang:latest

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV mongoURI mongodb://mongo:27017
ENV alo testggs

RUN go build

CMD ["./go-mongo"]