# Project_Restaurant_Angular_GO

## Introduction

A simple app where you can order food. The application has a login system and is mostly based on cookies. Users authenticate to the app with an email address and password. Each user has own role. Admin role allows to add, delete, modifidied products. 

## Install developer environment

### Backend

First download all packages

`go get .`

Ensure that there is MySQL Database on your system. Migrate schema using

`go run . --migrations=up`

to clear all schema use command

`go run . --migrations=down`

You can only run a few migrations with this command

`go run . --migrations=up X`

where `X` is a positive integer number indicating how many migrations will be up.

Roll back schema by step can be done using this command

`go run . --migrations=down X`

where `X` is a negative integer number indicating how many migrations will be rolled back.

After setup database and schema you can run application typing `go run ./...`

### Frontend

Install all packages. Make sure Angular is properly installed

`npm install`

Run application

`npm run start`

## Technologies

- GO 1.20
- Angular 15
- MariaDB
- Go Gin
- Golang-migrate

## Screens

![](/images/screen1.png)
![](/images/screen2.png)
![](/images/screen3.png)
![](/images/screen4.png)
![](/images/screen5.png)
![](/images/screen6.png)
![](/images/screen7.png)
![](/images/screen8.png)
![](/images/screen9.png)
![](/images/screen10.png)
![](/images/screen11.png)
![](/images/screen12.png)