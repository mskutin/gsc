FROM golang:1.14.2-alpine as BUILDER
WORKDIR /go/src/app
COPY . .
RUN cd /go/src/app/ && go build -o gsc

FROM alpine:3.12
COPY --from=BUILDER /go/src/app/ /app/
ENTRYPOINT ["/app/gsc"]