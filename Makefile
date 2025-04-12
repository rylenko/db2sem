.POSIX:

bin/:
	mkdir -p $@

bin/sqlc: | bin/
	GOBIN="$(realpath $(dir $@))" go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

.PHONY: sqlc
sqlc: | bin/sqlc
	./bin/sqlc generate -f internal/db/pg/sqlc.yaml
