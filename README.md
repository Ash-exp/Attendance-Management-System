# Attendance Management System

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

## **Table of Contents**

- [About the project](#about-the-project)
  - [API docs](#api-docs)
- [Getting started](#getting-started)
  - [Layout](#layout)
- [Run the project](#run-the-project)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## About the project

_A School Attendance Management System where there are three entities Principal, Teacher and Students where all of them will have different functionalities._

### API docs

All the API endpoints can be found in the [Postman Import Link](https://api.postman.com/collections/13277644-0383300d-4299-486a-bd08-e85538c6e5f4?access_key=PMAT-01GPRTR76H3X2FYD1W9JV4HT8P)

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
└── .env
└── example.env
├── .gitignore
├── main.go
├── go.sum
├── go.mod
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
- `main.go` project entry-point.
- `README.md` is a detailed description of the project.

## Run the project

```bash
go build main.go
go run main.go
```
