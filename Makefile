# Load enviroment variables
include ./config/.env

# Export enviroment variables to commands
export

# Variables
go_cover_file=coverage.out

help:: ## Show this help
	@ fgrep -h "##" $(MAKEFILE_LIST) | sort | fgrep -v fgrep | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

run:: ## Run go application
	@ go run . -port=$(HTTP_PORT)

test:: ## Do the tests in go
	@ go test -race -coverprofile $(go_cover_file) ./...

coverage:: test ## See of the coverage of the code on your default navigator, see more in https://go.dev/blog/cover
	@ go tool cover -html=$(go_cover_file)

open-front:: ## Open in default browser de the front
	open http://localhost:9000
