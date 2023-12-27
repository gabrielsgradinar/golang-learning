# golang-learning

# modules
 - go mod init "name"
 - go mod tidy
 - go work init ./repative_paths
 - go mod tidy -e

# testes

- go test
- go test -coverprofile=coverage.out
- go tool cover -html=coverage.out
- go test -bench=.
- go test -bench=. -count=5 -benchtime=3s

- fuzzing - testes de mutação
- go test -fuzz=. -run=^#
- go test -fuzz=. -fuzztime=10s -run=^#