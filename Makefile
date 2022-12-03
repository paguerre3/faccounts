DCU=docker-compose up --build --abort-on-container-exit
docker-up:
	$(DCU)
docker-push:
	docker-compose push
test-build: docker-up
release: docker-push