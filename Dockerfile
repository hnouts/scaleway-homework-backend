FROM golang:1.17

WORKDIR /go/src/github.com/hnouts/scaleway-homework-backend

COPY . .

RUN go mod download
RUN make clean
RUN make build

CMD ["./myapp"]
