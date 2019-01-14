FROM golang:latest as builder
WORKDIR /go/src/github.com/jicg/App
COPY . .
RUN go get github.com/tools/godep
#RUN go get
RUN godep go build -a -ldflags="-w -s"

FROM alpine as final
MAINTAINER jicg <284077319@qq.com>
RUN apk --update upgrade \
&& apk add sqlite
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
RUN rm -rf /var/cache/apk/*
WORKDIR /app
COPY --from=builder /go/src/github.com/jicg/App/App /app/App
COPY --from=builder /go/src/github.com/jicg/App/views /app/views
COPY --from=builder /go/src/github.com/jicg/App/static /app/static
COPY --from=builder /go/src/github.com/jicg/App/conf /app/conf
EXPOSE 8080
ENTRYPOINT ["/app/App"]