.PHONY: migrate-up
migrate-up:
	migrate -path ./migrations/user -database "postgres://postgres:postgres@localhost:5432/user?sslmode=disable" -verbose up
	migrate -path ./migrations/post -database "postgres://postgres:postgres@localhost:5433/post?sslmode=disable" -verbose up
#	migrate -path ./migrations/comment -database "postgres://postgres:postgres@localhost:5432/comments?sslmode=disable" -verbose up

.PHONY: migrate-down
migrate-down:
	migrate -path ./migrations/user -database "postgres://postgres:postgres@localhost:5432/user?sslmode=disable" -verbose down --all
	migrate -path ./migrations/post -database "postgres://postgres:postgres@localhost:5432/posts?sslmode=disable" -verbose down --all
#	migrate -path ./migrations/comment -database "postgres://postgres:postgres@localhost:5432/comments?sslmode=disable" -verbose down --all
