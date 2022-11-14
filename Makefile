install:
	@go mod tidy
	@go get github.com/kyoh86/richgo

test:
	@richgo test ./datastruct --cover -v