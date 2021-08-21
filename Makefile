help:
	@echo "You can perform the following:"
	@echo ""
	@echo "  check         Format, lint, vet, and test Go code"
	@echo "  cover         Show test coverage in html"

check:
	@echo 'Formatting, linting, vetting, and testing Go code'
	go fmt ./...
	golint ./...
	go vet ./...
	go test ./... -cover

cover:
	@echo 'Test coverage in html'
	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out
