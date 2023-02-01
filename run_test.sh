go test -coverprofile=coverage.out main_test.go main.go && go tool cover -html=coverage.out
