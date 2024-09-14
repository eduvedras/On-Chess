build:
	go build -o on-chess

run:
	./on-chess

brun:
	go build -o on-chess && ./on-chess

test:
	go test ./... -v
