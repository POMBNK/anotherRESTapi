<h1 align="center">anotherRESTapi</h1>

<img src="https://img.shields.io/badge/made%20by-POMBNK-blue.svg"  alt="">


## Description:
This project is a test with techs like **GIN,REST, JWT, Postgres** with implementing uncle Bob's [clean architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) in **GO**

## TODO:
- [X] Endpoints
    - [X] SIGN-UP
    - [X] SIGN-IN
    - [X] CRUD endpoints
- [X] CRUD
    - [X] Create
    - [X] Read
    - [X] Update
    - [X] Delete
- [X] Authorization
    - [X] Registration
    - [X] Authentication
- [X] Database
    - [X] Schema
    - [X] Migration
- [X] Quick setup
    - [X] Docker image

## CRUD endpoints:

| HTTP Method | Endpoint          |
|-------------|-------------------|
| GET         | /lists            |
| GET         | /lists/{id}       |
| POST        | /lists            |
| PUT         | /lists/{id}       |
| DELETE      | /lists/{id}       |
| GET         | /lists/{id}/items |
| POST        | /lists/{id}/items |

## Schema:
![Schema](https://i.imgur.com/K0s9V7T.png "Schema")

## Project setup

### Create your own .env file in root with params:

```
DB_PASSWORD=<your db password, must compare with password in docker-compose.yml>
SALT=<your salt for password cash>
SIGN=<your sign for JWT token>
```
### Start command with bash to build and run app

```
docker-compose up --build
```
### Start command with bash to run app

```
docker-compose up -d
```