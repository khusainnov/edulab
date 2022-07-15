m-up:
	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5434/postgres?sslmode=disable' up
m-down:
	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5434/postgres?sslmode=disable' down

dpsql_up:
	docker run --name=edulab -e POSTGRES_PASSWORD='qwerty' -p 5434:5432 -d --rm postgres

d-exec:
	docker exec -it 082cee8c6fce /bin/bash