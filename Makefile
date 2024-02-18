run:
	cd src && air -c .air.toml

update:
	cd src && go mod tidy

unit-test:
	cd src && go test -v -cover ./...