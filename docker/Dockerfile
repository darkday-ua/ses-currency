FROM golang:1.20.4-alpine3.18 as builder

RUN apk add --update make \
    && rm -rf /tmp/* \
    && rm -rf /var/cache/apk/*

RUN mkdir /build


COPY . .
RUN cd ./currency_service  && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -o /build/currency_service

FROM alpine:3.9

COPY --from=builder /build/currency_service /app/currency_service
ENV APP_PORT 8080
EXPOSE ${APP_PORT}
RUN chmod +x /app/currency_service 
   

ENTRYPOINT ["/app/currency_service"]
