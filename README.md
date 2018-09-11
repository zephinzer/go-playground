# Experiments in Go
A repository with code for myself to learn and practice Golang using practical applications.

## Running Applications
All projects use `Makefile`s to execute.

In any directory, run `make init` to initialise the project.

To run the project in development, use `make dev`.

To run the project in production, use `make prod`.

## Creating A New Project
To create a new project from the template, use `make create PROJECT=${PROJECT_NAME}` where `${PROJECT_NAME}` is the name of your new project in `kebab-case`.

## Design and Technical Decisions

### Logger

| Name | URL | Stars | Forks | Issues | PRs | Contributors | License |
| --- | --- | --- | --- | --- | --- | --- |
| Zap | https://github.com/uber-go/zap | 5079 | 364 | 25 | 6 | 48 | [MIT](https://github.com/uber-go/zap/blob/master/LICENSE.txt) |

### Router

| Name | URL | Stars | Forks | Issues | PRs | Contributors | License |
| --- | --- | --- | --- | --- | --- | --- | --- |
| github.com/julienschmidt/httprouter | https://github.com/julienschmidt/httprouter | 7705 | 767 | 45 | 13 | 33 | [Custom](https://github.com/julienschmidt/httprouter/blob/master/LICENSE) |
| github.com/gorilla/mux | https://github.com/gorilla/mux | 6984 | 834 | 17 | 5 | 67 | [BSD-3-Clause](https://github.com/gorilla/mux/blob/master/LICENSE) |
| github.com/go-chi/chi | https://github.com/go-chi/chi | 4084 | 273 | 17 | 7 | 48 | [MIT](https://github.com/go-chi/chi/blob/master/LICENSE) |