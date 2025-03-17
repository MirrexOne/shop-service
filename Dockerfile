FROM golang:1.24

WORKDIR ${GOPATH}/avito-shop/
COPY . ${GOPATH}/avito-shop/

RUN go build -o /build ./internal/app \
    && go clean -cache -modcache

EXPOSE 8081

CMD ["/build"]