benchmark:
	go test -short -bench BenchmarkSolvers -json -run NONE ./... | go run cmd/report.go > public/benchmark.json

lint:
	golangci-lint run ./...
test:
	go test ./...
