run:
	cd src && air -c .air.toml

update:
	cd src && go mod tidy

test:
	cd src/tests/integration && go test -v