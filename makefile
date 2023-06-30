.PHONY: build run clean all

build:
	go build -o myapp

run:
	./myapp

clean:
	rm -f myapp

all: clean build run