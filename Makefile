mock-dependencies:
	mockery --keeptree --output mocks/service --dir service --all
	mockery --keeptree --output mocks/repository --dir repository --all
