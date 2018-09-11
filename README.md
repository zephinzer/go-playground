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

| Name | URL | Stars | Forks | Issues | PRs | Contributors |
| --- | --- | --- | --- | --- | --- | --- |
| Zap | https://github.com/uber-go/zap | 5079 | 364 | 25 | 6 | 48 |
