.PHONY: tests test-binary-search test-selection-sort

tests: 
	go clean -testcache
	go test -cover -p=1 ./...

test-binary-search:
	@echo "▶️  Run tests for binary_search.go..."
	go test -cover -p=1 ./binary-search/

test-selection-sort:
	@echo "▶️  Run tests for selection_sort.go..."
	go test -cover -p=1 ./selection-sort/