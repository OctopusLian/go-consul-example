run:
	@echo "=============starting server============="
	go run cmd/main.go

consul-up:
	@echo "=============running a temporary consul============="
	docker run --rm -d --name=c1 -p 8500:8500 consul agent -dev -client=0.0.0.0 -bind=0.0.0.0

consul-down:
	@echo "=============stopping the temporary consul============="
	docker stop c1

redis-up:
	@echo "=============running a temporary redis============="
	docker run --rm --name redis-docker -d -p 6379:6379 redis

redis-down:
	@echo "=============stopping the temporary redis============="
	docker stop redis-docker
