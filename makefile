.PHONY: build run clean all

build:
	CGO_ENABLED=0 GOOS=linux go build -a -o myapp .

run:
	./myapp

clean:
	rm -f myapp

all: clean build run