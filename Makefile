FILE_COVER ?= cover.out

run/api:
	go run ./cmd/api/main.go


test/unit:
	go test ./... -p 2 -coverprofile=./${FILE_COVER}

cover/read: 
	go tool cover -func=./${FILE_COVER}

cover/read/html: 
	go tool cover -html=./${FILE_COVER}

# TODO: make this generic to only need to provide interface location
# mocks:
# 	mockgen -source=./src/domain/usecases/find_customer_by_tax_id.go -destination=./tests/mocks/mocks.go -package=mocks