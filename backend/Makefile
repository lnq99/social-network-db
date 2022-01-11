live:
	@export PATH=$(PATH):$(shell go env GOPATH)/bin ; cd cmd ; gin -p 8080 run cmd/main.go

run:
	# go build ./...
	go run cmd/main.go

ui:
	cd ../client && npm run build

swag:
	swag init -g cmd/main.go

setup:
	go get github.com/onsi/ginkgo/v2/ginkgo

test: unit_all integration e2e e2e_bdd

cover:
	go test -v -tags="unit integration e2e" -coverprofile unit_cover.out ./...
	go tool cover -html=unit_cover.out

unit_all:
	go test -v -cover -tags=unit ./...
	@echo -e '\n'

unit:
	-go test -tags=unit -v -cover ./pkg/auth
	@echo
	-go test -tags=unit -v -cover ./internal/repository
	@echo
	-go test -tags=unit -v -cover ./internal/service
	@echo -e '\n'
	

integration:
	go test -tags=integration -v ./test/integration
	@echo -e '\n'

e2e:
	go test -tags=e2e -v ./test/e2e
	@echo -e '\n'

e2e_100:
	go test -tags=e2e -count 100 ./test/e2e

e2e_bdd:
	go test -v -tags=e2e_bdd ./test/e2e
	@# ginkgo -tags=e2e_bdd ./test/e2e
	@echo -e '\n'


prof:
	go test -count 100 -tags="e2e" ./test/e2e -cpuprofile prof.cpu -memprofile prof.mem


postgres:
	docker start postgres14 ||\
	docker run --name postgres14 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=pw -d postgres:14-alpine
	@echo

createdb:
	docker exec -it postgres14 createdb --username=root --owner=root socialnetwork_test

dropdb:
	docker exec -it postgres14 dropdb socialnetwork_test

migrateup:
	migrate -path schema -database "postgresql://root:pw@localhost:5433/socialnetwork_test?sslmode=disable" -verbose up

migratedown:
	migrate -path schema -database "postgresql://root:pw@localhost:5433/socialnetwork_test?sslmode=disable" -verbose down

# migrate create -ext sql -dir ./schema -seq init

psql:
	docker exec -it postgres14 psql -d socialnetwork_test

fix_dirty:
	docker exec -it postgres14 psql -d socialnetwork_test -c 'update schema_migrations set dirty=false;'

show:
	docker exec -it postgres14 psql -d socialnetwork_test -c 'select * from Profile'
	docker exec -it postgres14 psql -d socialnetwork_test -c 'select * from Post'
