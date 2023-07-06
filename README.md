# scaleway-homework-backend

A backend service that interacts with a database to manager servers. The service provides an API that allows you to create and retrieve servers.

### Prerequisites

Before you start using the API, make sure you have the .env file in the root directory.
> not a great practice but i figured you would like it better to just have a one line setup

>    To be safe : 

>   DB_HOST=db
    DB_USER=myuser
    DB_NAME=mydatabase
    DB_PASSWORD=example

### How to Run

just enter ```docker-compose up```

This will start the API server on port 8000 and the database on port 5432.

**Wait for app service to receive a successful healthcheck**

![Alt text](/misc/image.png)

## API Endpoints

Below are the available API endpoints.

#### POST /server
This endpoint is used to create a new server in the database.

**Request Body:**
```
{
    "name": "test",
    "type": "small",
    "status": "starting"
}
```
**Example Request:**
```
curl --location 'localhost:8000/server' \
--header 'Content-Type: application/json' \
--data '{
    "name": "testCurl",
    "type": "small",
    "status": "running"
}'
```

**Example Response:**
```
{"ID":7,"CreatedAt":"2023-07-06T12:21:39.950736381Z","UpdatedAt":"2023-07-06T12:21:39.950736381Z","DeletedAt":null,"Name":"testCurl","Type":"small","Status":"running"}
```

#### GET /servers
This endpoint is used to retrieve all servers in the database.

**Example Request:**
```
curl --location 'localhost:8000/servers'
```

**Example Response:**
```
[
    {"ID":1,"CreatedAt":"2023-07-06T12:21:39.950736381Z","UpdatedAt":"2023-07-06T12:21:39.950736381Z","DeletedAt":null,"Name":"test","Type":"small","Status":"starting"},
    {"ID":2,"CreatedAt":"2023-07-06T12:21:39.950736381Z","UpdatedAt":"2023-07-06T12:21:39.950736381Z","DeletedAt":null,"Name":"test","Type":"small","Status":"starting"},
...]
```

#### GET /servers/{id}
This endpoint is used to retrieve a server by id.

**Example Request:**
```
curl --location 'localhost:8000/servers/1'
```

**Example Response:**
```
{"ID":1,"CreatedAt":"2023-07-06T12:21:39.950736381Z","UpdatedAt":"2023-07-06T12:21:39.950736381Z","DeletedAt":null,"Name":"test","Type":"small","Status":"starting"}
```