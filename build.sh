gofmt -w .
golint .
go test . | grep -v ^ok | grep -v '^?'
