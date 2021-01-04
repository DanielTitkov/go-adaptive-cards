.PHONY: test
test: 
	go test -coverprofile cover.out

.PHONY: run
run:
	go run examples/main.go