# Attendance Management System

## **Table of Contents**

- [About the project](#about-the-project)
  - [API docs](#api-docs)
- [Prerequisites](#prerequisites)
  - [Linux](#linux)
  - [macOS](#macos)
  - [Windows](#windows)
- [Getting started](#getting-started)
  - [Layout](#layout)
- [Run the project](#run-the-project)
- [Access the project logs](#access-the-project-logs)
- [Configure Database through PgAdmin](#configure-database-through-pgadmin)
- [Entry Point](#entry-point)
- [Possible Issues](#possible-issues)
- [Shutting down the application](#shutting-down-the-application)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## About the project

_A School Attendance Management System where there are three entities Principal, Teacher and Students where all of them will have different functionalities._

### API docs

All the API endpoints can be found in the [Postman Import Link](https://api.postman.com/collections/13277644-0383300d-4299-486a-bd08-e85538c6e5f4?access_key=PMAT-01GPRTR76H3X2FYD1W9JV4HT8P)

## Prerequisites

This project uses `Docker` for containerization and requires Docker to be installed in your system. Refer the following if you don't have Docker yet.

### Linux

Run this quick and easy install script provided by Docker:

```sh
curl -sSL https://get.docker.com/ | sh
```

If you're not willing to run a random shell script, please see the [installation](https://docs.docker.com/engine/installation/linux/) instructions for your distribution.

### macOS

Download and install [Docker Community Edition](https://www.docker.com/community-edition). if you have Homebrew-Cask, just type `brew install --cask docker`. Or Download and install [Docker Toolbox](https://docs.docker.com/toolbox/overview/). [Docker For Mac](https://docs.docker.com/docker-for-mac/) is nice, but it's not quite as finished as the VirtualBox install. [See the comparison](https://docs.docker.com/docker-for-mac/docker-toolbox/).

> **NOTE** Docker Toolbox is legacy. You should to use Docker Community Edition, See [Docker Toolbox](https://docs.docker.com/toolbox/overview/).

### Windows

Instructions to install Docker Desktop for Windows can be found [here](https://docs.docker.com/desktop/windows/install/)

Once installed, open powershell as administrator and run:

```powershell
# Display the version of docker installed:
docker version
```

## Getting started

Below we describe the conventions or tools specific to golang project.

### Layout

```tree
├── .github
├── api
│   ├── controllers
│   │   └── attendance.controller.go
│   │   └── attendance.teacher.controller.go
│   │   └── db.go
│   │   └── home.controller.go
│   │   └── routes.go
│   │   └── teacher.controller.go
│   │   └── user.controller.go
│   ├── middlewares
│   │   └── middlewares.go
│   ├── models
│   │   └── attendance.go
│   │   └── attendance.teacher.go
│   │   └── teacher.go
│   │   └── user.go
│   ├── responses
│   │   └── json.go
│   ├── seeds
│   │   └── seeder.go
│   └── utils
│       └── formaterror
│           └── formaterror.go
├── .env
├── .gitignore
├── docker-compose.yml
├── Dockerfile
├── example.env
├── go.sum
├── go.mod
├── main.go
└-─ README.md
```

A brief description of the layout:

- `api` places most of project business logic and locate
  - `controllers` package - route handlers and controllers
  - `middlewares` package - middleware configurations
  - `models` package - db models and methods
  - `responses` package - Json handlers
  - `seed` package - db migrations and model association
  - `utils` package - error formater
- `.gitignore` varies per project, but all projects need to ignore `.env` file.
- `example.env` dummy structure of env variables
- `docker-compose.yml` docker services and their configurations
- `Dockerfile` project image instructions
- `main.go` project entry-point.
- `README.md` is a detailed description of the project.

## Run the project

```bash
docker-compose up -d
```

## Access the project logs

```bash
docker logs full_app
```

## Configure Database through PgAdmin

Go to your browser and open:

```link
http://localhost:5050
```

- Login by the credentials used in the environment veriables at `.env` file.
- Choose and create a server with `HOST`, `PORT`, `USERNAME` and `PASSWORD` **as per the environment values only**.

## Entry Point

Now the application has successfully started. All the project API end-points are running and open at:

```link
http://localhost:8080
```

## Possible Issues

1. Issue with environmental variable: Maybe you used invalid details in the .env file, this might break your build. Carefully read the logs when you run docker-compose up. If you identify any and fix, terminate the process using `Ctrl+C`.

   - Stop running process using :

     ```bash
     docker-compose down --remove-orphans --volumes
     ```

   - Then to build again, use :

     ```bash
     docker-compose up --build
     Observe the --build.
     ```

2. Any other issue: If you have any other issue, try and read the logs and fix them. then follow the exact same process to build again.

## Shutting down the application

After successful testing of the application, you can shut it down using:

```bash
docker-compose down
```

Or to also remove the volumes too:

```bash
docker-compose down --remove-orphans --volumes
```

If you want to remove all dangling images, run:

```bash
docker system prune
```

To remove all unused images, not just dangling ones, run:

```bash
docker system prune -a
```

To remove all unused images not just dangling ones and volumes, run:

```bash
docker system prune -a --volumes
```
