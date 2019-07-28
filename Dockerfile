FROM scratch
COPY go/bin/soma /soma
CMD ["/soma"]