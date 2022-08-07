# hire.me
bemobi test

## Table of contents
* [General info](#general-info)
* [Technologies](#technologies)
* [Instalation](#instalation)
* [Usage](#usage)

## General info
Bemobi hire.me test.

## Technologies
Project is created with:
* Golang 1.18
* Docker
* Mysql 5.7

## Instalation
*You need to have Docker and docker-compose installed in your machine

run the following command:
```
$ docker-compose up -d --build
```

## Usage
Theres a Postman Collection for you to import and use to test the api [Postman Collection](Hire.me.postman_collection.json)

***Get Most Used***: request returns the top 10 most requested URL's

***Get By Alias***: request returns a previous saved URL, just replace the teste3 after retrive/ with a valid alias.

***Save With Custom Alias***: request to save a URL, the alias automatically created. Check the Query Params


***Save No Custom Alias***: request to save a URL with a predefined alias. Check the Query Params
