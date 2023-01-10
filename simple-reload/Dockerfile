FROM golang:1.19 as builder
WORKDIR /code
COPY main.go .
COPY go.mod .
COPY .air.toml .
# install air
RUN go install github.com/cosmtrek/air@latest

CMD ["air", "-c", ".air.toml"]
