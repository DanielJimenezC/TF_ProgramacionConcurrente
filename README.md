# TF_ProgramacionConcurrente

## Api

#### PreRequisites

Install Go Package pq.

    go get github.com/lib/pq

Install Go Package mux.

    go get github.com/gorilla/mux

#### Build & Run Api

Go to CovidAppApi package then build it.

    go build

Run CovidAppApi.

    ./CovidAppApi

If everything is alright, the message "Server is up and running..." will be displayed.

Api is running on localhost:5000

### API Routes

#### Users
GET

    /api/users

    response:
    {
        id: "id",
        username: "username",
        password: "password"
    }

POST

    /api/users

    request body:
    {
        username: "username",
        password: "password"
    }    