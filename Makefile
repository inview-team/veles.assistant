.PHONY: build run stop

build:
	docker-compose build

up:
	docker-compose up -d

down:
	docker-compose down

redis:
	docker run -d -p 6379:6379 --name my-redis redis:alpine

clean:
	docker-compose down
	docker system prune -af