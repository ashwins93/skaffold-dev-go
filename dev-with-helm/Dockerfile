FROM golang:1.19
WORKDIR /code
COPY main.go .
COPY go.mod .
COPY reflex.conf .
RUN go install github.com/cespare/reflex@latest

CMD ["reflex", "-c", "/code/reflex.conf"]
