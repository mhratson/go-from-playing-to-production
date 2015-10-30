# Makefile
test:
	go test

re-test:
	reflex -r '\.go' $(MAKE) test

dev-deps:
	go get github.com/cespare/reflex
