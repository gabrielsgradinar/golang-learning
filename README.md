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


# Go Routines

- `go` + nome_da_funcão
- rodar simultâneamente diversas tarefas
- se o programa acaba, as threads que estão executando são finalizadas
- `wait groups`
    - adicionar quantidade de tarefas / operações
    - informar que você terminou uma operação
    - esperar até que as operações sejam finalizadas