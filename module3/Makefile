build:
	docker build --tag bjy168/httpserver -f ./Dockerfile ../module2

push:
	docker push bjy168/httpserver

run:
	docker run -p 8080:8080 --name httpserver -d bjy168/httpserver httpserver
