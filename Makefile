FILE_COVER ?= cover.out

run/api:
	go run ./cmd/api/main.go


test/unit:
	go test ./... -p 2 -coverprofile=./${FILE_COVER}

cover/read: 
	go tool cover -func=./${FILE_COVER}

cover/read/html: 
	go tool cover -html=./${FILE_COVER}
