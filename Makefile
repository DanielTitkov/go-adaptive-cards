.PHONY: test
test: 
	go test -coverprofile cover.out -v

.PHONY: run
run:
	go run examples/main.go