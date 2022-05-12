# builder image
FROM golang:1.18.2 as builder

WORKDIR /build
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -o app .

# generate clean, final image for end users
FROM alpine:3.14

WORKDIR /app

COPY --from=builder /build/app .

# executable
ENTRYPOINT [ "./app" ]

# arguments that can be overridden
CMD [ "serve" ]
