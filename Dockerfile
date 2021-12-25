FROM golang:1.17 AS builder

WORKDIR /opt

COPY . .

RUN go build

FROM gcr.io/distroless/cc-debian11 AS runtime

COPY --from=builder /opt/special-tribble /usr/local/bin/special-tribble

ENTRYPOINT ["/usr/local/bin/special-tribble"]
