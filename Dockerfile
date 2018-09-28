FROM golang:1.9

COPY application/main.go application/

WORKDIR application/

CMD ["go", "run", "main.go"]

