docker-build:
	docker build -t message-board .

docker-run:
	docker run --rm --name message-board -p 8000:8000 message-board
