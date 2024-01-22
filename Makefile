FILE_COVER ?= cover.out

run/api:
	go run ./cmd/api/main.go


test/unit:
	go test ./... -p 2 -coverprofile=./${FILE_COVER}

cover/read: 
	go tool cover -func=./${FILE_COVER}

cover/read/html: 
	go tool cover -html=./${FILE_COVER}


# USAGE EXAMPLE: make INTERFACE_PATH=src/domain/usecases/factory_customer.go FILE_NAME=factory_customer.go mocks
mocks:
	mockgen -source=./${INTERFACE_PATH} -destination=./tests/mocks/mocks_${FILE_NAME} -package=mocks
