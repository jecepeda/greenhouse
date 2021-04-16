
sql-dev:
	pgcli -h localhost -p 5432 -U ghuser -w ghpassword -d greenhouse

start-db-test:
	docker run --name test-db --rm -e POSTGRES_USER=gtest -e POSTGRES_PASSWORD=gtest -e POSTGRES_DB=gtestdb -d -p 5431:5432 postgres:13.2

test:
	SQL_DB='host=localhost port=5431 user=gtest password=gtest dbname=gtestdb sslmode=disable' go test -cover ./...