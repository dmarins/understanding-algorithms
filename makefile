.PHONY: tests test-binary-search

tests: 
	go clean -testcache
	go test -cover -p=1 ./...

test-binary-search:
	@echo "▶️  Run tests for binary_search.go..."
	go test -cover -p=1 ./binary-search/