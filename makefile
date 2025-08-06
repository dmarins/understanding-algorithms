.PHONY: tests test-binary-search test-selection-sort test-quick-sort

tests: 
	go clean -testcache
	go test -cover -p=1 ./...

test-linear-search:
	@echo "▶️  Run tests for linear_search.go..."
	go test -cover -p=1 ./search/linear-search/

test-binary-search:
	@echo "▶️  Run tests for binary_search.go..."
	go test -cover -p=1 ./search/binary-search/

test-selection-sort:
	@echo "▶️  Run tests for selection_sort.go..."
	go test -cover -p=1 ./sort/selection-sort/

test-quick-sort:
	@echo "▶️  Run tests for quick_sort.go..."
	go test -cover -p=1 ./sort/quick-sort/