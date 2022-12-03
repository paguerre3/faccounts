IMAGE=paguerre/faccounts
DCR=docker-compose up --build --abort-on-container-exit
docker-compose-up:
	$(DCR) -t $(IMAGE)
docker-push:
	docker push $(IMAGE)
test-build: docker-compose-up
release: docker-compose-up docker-push