.PHONY: compose-up
compose-up:
	docker compose up -d

.PHONY: compose-down
compose-down:
	docker compose down -d

.PHONY: migrate-up
migrate-up:
	migrate -path ./migrations/user -database "postgres://postgres:postgres@localhost:5432/user?sslmode=disable" -verbose up
	migrate -path ./migrations/post -database "postgres://postgres:postgres@localhost:5433/post?sslmode=disable" -verbose up

.PHONY: migrate-down
migrate-down:
	migrate -path ./migrations/user -database "postgres://postgres:postgres@localhost:5432/user?sslmode=disable" -verbose down --all
	migrate -path ./migrations/post -database "postgres://postgres:postgres@localhost:5432/post?sslmode=disable" -verbose down --all

.PHONY: build-auth
build-auth:
	docker build -t mikehsu0618/reddit-auth -f ./deployments/docker/auth/dockerfile .
	docker push mikehsu0618/reddit-auth

.PHONY: build-user
build-user:
	docker build -t mikehsu0618/reddit-user -f ./deployments/docker/user/dockerfile .
	docker push mikehsu0618/reddit-user

.PHONY: build-post
build-post:
	docker build -t mikehsu0618/reddit-post -f ./deployments/docker/post/dockerfile .
	docker push mikehsu0618/reddit-post

.PHONY: build-reddit
build-reddit:
	docker build -t mikehsu0618/reddit -f ./deployments/docker/reddit/dockerfile .
	docker push mikehsu0618/reddit

.PHONY: kube-apply
kube-apply:
	kubectl apply -f ./deployments/kube \
	-f ./deployments/kube/auth \
	-f ./deployments/kube/post \
	-f ./deployments/kube/reddit \
	-f ./deployments/kube/user

.PHONY: kube-delete
kube-delete:
	kubectl delete -f ./deployments/kube \
	-f ./deployments/kube/auth \
	-f ./deployments/kube/post \
	-f ./deployments/kube/reddit \
	-f ./deployments/kube/user