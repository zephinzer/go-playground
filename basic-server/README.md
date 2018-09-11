# Basic Server
A basic HTTP server with production ready configuration settings.

## Scope
- [x] Listen on a port specified through a configuration file
- [x] Listen on a port specified through the environment
- [x] A default port is available if no configuration is specified
- [x] Respond with `"hello world"` on the apex endpoint

## Running
Prepare the project by running:

```bash
make init;
```

Run the project in development by running:

```bash
make dev;
```

Run the project in production by running:

```bash
make prod;
```

To build just the executable:

```bash
make compile;
```
