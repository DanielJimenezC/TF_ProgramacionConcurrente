# TF_ProgramacionConcurrente

## Api

#### PreRequisites

Install Go Package pq.

    go get github.com/lib/pq

Install Go Package mux.

    go get github.com/gorilla/mux

Install Go Package yaml

    got get gopkg.in/yaml.v3

Install Go Package bcrypt

    got get golang.org/x/crypto/bcrypt

#### Build & Run Api

Go to CovidAppApi package then build it.

    go build

Run CovidAppApi.

    ./CovidAppApi

If everything is alright, the message "Server is up and running..." will be displayed at localhost/5000.

Api is running on localhost:5000

### API Routes

#### Users
GET (Get all users)

    /api/users

    response:
    [
        {
            id: "id",
            username: "username"
        }
    ]

GET (Get user by Id)

    /api/users/{id}

    response:
    {
        id: "id",
        username: "username"
    }

PUT (Update user)

    /api/users/{id}

    request body:
    {
        username: "username",
        password: "password"
    }   

POST (Signup - Create user)

    /api/users/signup

    request body:
    {
        username: "username",
        password: "password"
    }    

POST (User login)

    /api/users/login

    request body:
    {
        username: "username",
        password: "password"
    }  

DELETE (Delete user by Id)

    /api/users/{id}
