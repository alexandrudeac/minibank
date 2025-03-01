BUILD_DIR=./build
COVERAGE_DIR=${BUILD_DIR}/coverage
COVERAGE_OUT_PATH=${COVERAGE_DIR}/coverage.out
COVERAGE_HTML_PATH=${COVERAGE_DIR}/coverage.html
DB_URL=postgresql://root:secret@localhost:5432/minibank?sslmode=disable

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)
sqlc:
	sqlc generate
test:
	mkdir -p ${COVERAGE_DIR} && go test -coverprofile=${COVERAGE_OUT_PATH}  ./... && go tool cover -html=${COVERAGE_OUT_PATH} -o ${COVERAGE_HTML_PATH}