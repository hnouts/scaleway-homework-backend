# build stage
FROM golang:1.20-bullseye AS builder

WORKDIR /go/src/github.com/hnouts/scaleway-homework-backend

COPY . .

RUN go mod download
RUN make clean
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp .

# final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /go/src/github.com/hnouts/scaleway-homework-backend/myapp .

CMD ["./myapp"]
