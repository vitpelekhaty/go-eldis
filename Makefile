.PHONY: test testAll
all: test

test: testMain testArchive testDate testResponse

testMain:
	@echo "test go-eldis..."
	go test -v -timeout 30s

testArchive:
	@echo "test go-eldis/archive..."
	go test -v -timeout 30s github.com/vitpelekhaty/go-eldis/archive

testDate:
	@echo "test go-eldis/date..."
	go test -v -timeout 30s github.com/vitpelekhaty/go-eldis/date

testResponse:
	@echo "test go-eldis/response..."
	go test -v -timeout 30s github.com/vitpelekhaty/go-eldis/response

testAll: test
	@echo "run integration tests..."
	go test -v -timeout 30s . -tags=integration -args ${params}
