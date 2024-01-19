help:
	@echo "You can perform the following:"
	@echo ""
	@echo "  check         Format, vet, and unit test Go code"
	@echo "  cover         Run & show test coverage in html"
	@echo "  int           Run integration tests"
	@echo "  int-cover     Run & show integration test coverage in html"
	@echo "  lint          Lint Go code using staticcheck"

check:
	@echo 'Formatting, vetting, and testing Go code'
	go fmt ./...
	go vet ./...
	go test ./... -cover

lint:
	@echo 'Linting code using staticcheck'
	staticcheck -f stylish ./...

cover:
	@echo 'Unit test coverage in html'
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

int:
	@echo 'Run integration tests'
	go test ./... -tags=integration

int-cover:
	@echo 'Integration test coverage in html'
	go test -coverprofile=coverage.out -tags=integration
	go tool cover -html=coverage.out
