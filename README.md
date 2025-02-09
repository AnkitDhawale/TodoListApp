# ToDo List Application

# DB Design
We will run postgres db in a docker container:
1. docker pull postgres
2. docker run --name local-postgres -e POSTGRES_USER=ankit -e POSTGRES_PASSWORD=app123 -e POSTGRES_DB=todo_app_db -p 5432:5432 -d postgres

# Create a database
CREATE DATABASE todo_app_db;

# Create a new user and grant privileges
GRANT ALL PRIVILEGES ON DATABASE todo_app_db TO ankit;

# Connect to the database
psql -d todo_app_db -U ankit

# Swagger documentation
1. swag init
2. run project: go run main.go
3. go to user http://localhost:<PORT>/swagger/index.html

# Following resources are kept in path github.com/AnkitDhawale/TodoListApp/resources
1. Application requirement pdf
2. Postman collection 

# Using docker image from dockerhub public repo:
1. Build & push image: \
docker build -t your-dockerhub-username/todo-app:latest . \
docker login: Enter username & password \
docker push your-dockerhub-username/todo-app:latest

2. Pull & run it:
docker pull ankit8796/todo-app:latest
docker run -p 8080:8080 ankit8796/todo-app:latest
