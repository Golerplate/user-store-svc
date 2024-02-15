run:
	air -c .air.toml

test:
	cd src/tests/integration && go test -v