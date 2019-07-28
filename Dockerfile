FROM scratch
COPY go/bin/soma go/bin/
CMD ["/go/bin/soma"]