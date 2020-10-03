default:
	@echo "=============building Local API============="
	docker-compose up
	docker exec -it graphsproject_redis_1 redis-cli FLUSHALL

up: default
	@echo "=============starting api locally============="
	docker-compose up -d

logs:
	docker-compose logs -f

build:
	docker-compose build

down:
	docker-compose down

test:
	go test -v -cover ./...

clean: down
	@echo "=============cleaning up============="
	rm -f api
	docker system prune -f
	docker volume prune -f