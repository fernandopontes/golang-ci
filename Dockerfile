FROM scratch

ADD ./src /go/src/soma
WORKDIR /go/src/soma

COPY go/bin/soma /soma
CMD ["/soma"]