.PHONY: build
build:
	go build -o airshipctl

.PHONY: update
update:
	go get github.com/ian-howell/airshipctl@master

.PHONY: clean
clean:
	rm -fr airshipctl
