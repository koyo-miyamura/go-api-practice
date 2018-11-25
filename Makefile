test:
	go test $(shell go list ./...)

test-nocach:
	go test $(shell go list ./...) -count=1

setup:
	dep ensure
	go run lib/migration/main.go
	go run lib/seed/main.go
