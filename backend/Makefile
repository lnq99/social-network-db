live:
	@export PATH=$(PATH):$(shell go env GOPATH)/bin ; cd cmd ; gin -p 8080 run cmd/main.go

run:
	go build ./...
	go run cmd/main.go

ui:
	cd ../client && npm run build
