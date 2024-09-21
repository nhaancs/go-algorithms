tidy:
	go mod tidy && go mod vendor

test:
	go test -count 1 -race ./...