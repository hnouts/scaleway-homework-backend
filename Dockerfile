# build stage
FROM golang:1.20-bullseye AS builder

WORKDIR /go/src/github.com/hnouts/scaleway-homework-backend

COPY . .

RUN go mod download
RUN make clean
RUN make build

# final image
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /go/src/github.com/hnouts/scaleway-homework-backend/myapp .

CMD ["./myapp"]
