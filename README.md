# ToDo List Application
The ToDo List application allows users to manage their tasks efficiently.

# DB Design
We will run postgres db in a docker container:
1. docker pull postgres
2. docker run --name local-postgres -e POSTGRES_USER=ankit -e POSTGRES_PASSWORD=app123 -e POSTGRES_DB=todo_app_db -p 5432:5432 -d postgres

# Follow steps for database:
CREATE DATABASE todo_app_db;
GRANT ALL PRIVILEGES ON DATABASE todo_app_db TO ankit;
psql -d todo_app_db -U ankit

# Swagger documentation
1. swag init
2. run project: go run main.go
3. go to user http://localhost:<PORT>/swagger/index.html

# Useful resources for building and testing this app
1. Application requirement pdf
2. Postman collection
refer from path: https://github.com/AnkitDhawale/TodoListApp/main/resources