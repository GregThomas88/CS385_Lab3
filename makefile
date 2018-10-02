all: go

go:
	docker build -t go .
	docker run -d --name goapp go

