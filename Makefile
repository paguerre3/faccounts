DCU=docker-compose up --build --abort-on-container-exit
docker-up:
	$(DCU)
test-build: docker-up