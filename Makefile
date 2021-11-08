.PHONY: test testAll
all: test

test: testArchive testDate testResponse

testArchive:
	@echo "test go-eldis/archive..."
	go test -v -timeout 30s github.com/vitpelekhaty/go-eldis/archive

testDate:
	@echo "test go-eldis/date..."
	go test -v -timeout 30s github.com/vitpelekhaty/go-eldis/date

testResponse:
	@echo "test go-eldis/response..."
	go test -v -timeout 30s github.com/vitpelekhaty/go-eldis/response

iTestUOMList:
	@echo "run connection.UOMList() integration test..."
	go test -v . -tags=integration -run TestConnection_UOMList -args ${params}

iTestListForDevelopment:
	@echo "run connection.ListForDevelopment() integration test..."
	go test -v . -tags=integration -run TestConnection_ListForDevelopment -args ${params}
