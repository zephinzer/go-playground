# Project Name
Short project description.

## Scope
- [x] Demonstrate serving of static files
- [ ] Demonstrate using a URL path parameter
- [ ] Demonstrate using a HTTP header
- [x] Demonstrate distributed sub-routers
- [x] Demonstrate directory/route parity
- [x] Demonstrate server middleware
- [x] Demonstrate accepting POST requests
- [x] Demonstrate accepting GET requests
- [x] Demonstrate accepting PUT requests
- [x] Demonstrate accepting DELETE requests
- [x] Demonstrate accepting OPTIONS requests
- [x] Demonstrate reading POST/PUT request body as string
- [x] Demonstrate reading POST/PUT request body as unstructured data

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

## Design Considerations
### Request Routing
| Router | URL | Stars | Forks | Issues | PRs | Contributors | License |
| --- | --- | --- | --- | --- | --- | --- | --- |
| github.com/julienschmidt/httprouter | https://github.com/julienschmidt/httprouter | 7705 | 767 | 45 | 13 | 33 | [Custom](https://github.com/julienschmidt/httprouter/blob/master/LICENSE) |
| github.com/gorilla/mux | https://github.com/gorilla/mux | 6984 | 834 | 17 | 5 | 67 | [BSD-3-Clause](https://github.com/gorilla/mux/blob/master/LICENSE) |
| github.com/go-chi/chi | https://github.com/go-chi/chi | 4084 | 273 | 17 | 7 | 48 | [MIT](https://github.com/go-chi/chi/blob/master/LICENSE) |

