FROM golang:1.19.3-alpine AS builder

WORKDIR /usr/src/app
COPY . .
WORKDIR /usr/src/app/cmd
RUN go build -o ../build/out/squadhealth


FROM alpine:3.17
WORKDIR /usr/bin
COPY --from=builder /usr/src/app/build/out/squadhealth .
ENTRYPOINT [ "./squadhealth" ]
