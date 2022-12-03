DCU=docker-compose up --build --abort-on-container-exit
docker-up:
	$(DCU)
dockerhub-login:
    docker login -u $(DOCKER_USERNAME) -p $(DOCKER_PASSWORD)
docker-push:
	docker-compose push
test-build: docker-up
release: dockerhub-login docker-push