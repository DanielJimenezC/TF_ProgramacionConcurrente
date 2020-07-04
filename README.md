# CovidApp
Developers:
- [Daniel Jimenez](https://github.com/DanielJimenezC)
- [Luis Kcomt](https://github.com/kcomt)


# Index
1. [Project Description](#Project%20Description)
2. [PreRequisites](#PreRequisites)
3. [BackEnd](#Backend)
    1. [PreRequisites](#API%20PreRequisites)
    2. [Build & Run Api](#Build%20&%20Run%20Api)
    3. [API Routes](#API%20Routes)
4. [FrontEnd](#FrontEnd)
    1. [PreRequisites](#FrontEnd%20PreRequisites)
    2. [Run Web Application](#Run%20Web%20Application)
5. [Bibliography](#Bibliography)

&nbsp;

# Project Description
<p style='text-align: justify;'>The project is a web application that helps predict if a person is infected with Covid-19, taking into account their location and identifying risk groups. The project has a frontend developed with React and a Backend developed with Go.</p> 

&nbsp;

# PreRequisites
<p style='text-align: justify;'>To run CovidApp (Backend & Frontend) you must have Go and Node.js installed.
If you don't have them, you can download and install from the following links.</p>
Go

    https://golang.org

Node.Js

    https://nodejs.org/es/

&nbsp;

# Backend
<p style='text-align: justify;'>The backend was developed using Go. We use Clean Architecture to divide the application into layers in order to decouple the code and make the maintenance or change of any library or connection easy.</p> 

&nbsp;

## API PreRequisites

Install Go Package pq.

    go get github.com/lib/pq

Install Go Package mux.

    go get github.com/gorilla/mux

Install Go Package yaml.

    got get gopkg.in/yaml.v3

Install Go Package bcrypt.

    got get golang.org/x/crypto/bcrypt

&nbsp;

## Build & Run Api

Go to CovidAppApi package, then build it.

    go build

Run CovidAppApi.

    ./CovidAppApi

If everything is alright, the message "Server is up and running..." will be displayed at localhost/5000.

*Api is running on localhost:5000*

&nbsp;

## API Routes

### USERS

**GET** *(Get all users)* <a name="GET"></a>

    /api/users

    response:
    [
        {
            id: "id",
            username: "username"
        }
    ]

&nbsp;

**GET** *(Get user by Id)*

    /api/users/{id}

    response:
    {
        id: "id",
        username: "username"
    }

&nbsp;

**PUT** *(Update user)*

    /api/users/{id}

    request body:
    {
        username: "username",
        password: "password"
    }   

&nbsp;

**POST** *(Signup - Create user)*

    /api/users/signup

    request body:
    {
        username: "username",
        password: "password"
    }    

&nbsp;

**POST** *(User login)*

    /api/users/login

    request body:
    {
        username: "username",
        password: "password"
    }  

&nbsp;

**POST** *(Prediction of Covid in users)*

    /api/users/prediction

    request body:
    {
        "edad": "14",
        "peso": "66.39",
        "distrito": "Ventanilla",
        "tos": "false",
        "fiebre": "true",
        "dificultadRespirar": "false",
        "perdidaOlfato": "true",
        "enfermo": "true"
    }

    reponse body:
    {
        "enfermo": "true",
        "chance": 56
    }

&nbsp;

**POST** *(addBlock)*

    /api/users/blockchain

    request body:
    {
        "name": "name",
        "birthday": "01/01/2001",
        "dni": "07745011",
        "telefono": "90149558"
    }

&nbsp;

**DELETE** *(Delete user by Id)*

    /api/users/{id}

&nbsp;

# FrontEnd
<p style='text-align: justify;'>The web application was developed using React, a Javascript library. This application receives, and sends, information in Json format to the REST service "CovidApi". The web application has a login page, a module to know the risk groups, a module to perform a virtual diagnosis and a section for frequently asked questions.</p> 

&nbsp;

## FrontEnd PreRequisites

Install node modules using

    npm install

&nbsp;

## Start Web Application

After started CovidAppApi, start web application using

    npm start

&nbsp;

# Bibliography
- Go ( https://golang.org/doc/ )
- React ( https://es.reactjs.org/docs/getting-started.html )
- Package Pq ( https://godoc.org/github.com/lib/pq )
- Package Mux ( https://godoc.org/github.com/gorilla/mux )
- Package Yaml.v3 ( https://godoc.org/gopkg.in/yaml.v3 )
- Package Bcrypt ( https://godoc.org/golang.org/x/crypto/bcrypt )

&nbsp;

<p style='text-align: center;'>July - 2020</p>