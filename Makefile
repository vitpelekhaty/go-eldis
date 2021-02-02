.PHONY: test testAll
all: test

test: testMain testArchive testResponse

testMain:
	@echo "test go-eldis..."
	go test -v -timeout 30s

testArchive:
	@echo "test go-eldis/archive..."
	go test -v -timeout 30s github.com/vitpelekhaty/go-eldis/archive

testResponse:
	@echo "test go-eldis/response..."
	go test -v -timeout 30s github.com/vitpelekhaty/go-eldis/response

testAll: test
	@echo "run all tests..."
#	go test -v -timeout 30s . -tags=integration -args ${params}
