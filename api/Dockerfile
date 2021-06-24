FROM golang:1.16-alpine AS build

COPY . /app

WORKDIR /app

# disable dynamic linking, build statically
ENV CGO_ENABLED=0
RUN go build -o wordcounter .

# we have a static binary, don't need the whole alpine image anymore
FROM scratch

# needed for https certificates, just take from alpine image for now
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /app/wordcounter /app/wordcounter

EXPOSE 8080
ENTRYPOINT ["/app/wordcounter"]
