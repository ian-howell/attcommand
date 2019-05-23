.PHONY: build
build:
	go build -o airshipctl

.PHONY: clean
clean:
	rm -fr airshipctl
