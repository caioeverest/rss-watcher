FROM golang:1.10-alpine

RUN  mkdir -p /go/src \
  && mkdir -p /go/bin \
  && mkdir -p /go/pkg
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH

# now copy your app to the proper build path
RUN mkdir -p $GOPATH/src/github.com/caioever/rss-watcher/
ADD . $GOPATH/src/github.com/caioever/rss-watcher/

WORKDIR $GOPATH/src/github.com/caioever/rss-watcher/
RUN go build -o main .
CMD ["/go/src/github.com/caioever/rss-watcher/main"]
