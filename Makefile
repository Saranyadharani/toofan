.PHONY: build install clean

build:
	go build -o toofan .

install:
	go install .

clean:
	rm -f toofan

run:
	go run .
