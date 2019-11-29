# Udemy - Go 101 Curriculum

Estimated course time ~28 hours

## Section 1 - Introduction

- Course overview. About the author (Mac, Windows 10, Ubuntu 18, RaspberryPi)
- Resources. Code. Community (Q&A, Documentation, code sharing & others, Discord)
- Go in a nutshell. Go authors & philosophy
- Installation on Mac (Go, Git & Editor)
- Installation on Windows (Go, Git & Editor)
- Installation on Linux (Ubuntu & RaspberryPi) (Go, Git & Editor)
- Git basics & Github (clone udemy-go101, repo structure)
- `$GOPATH` & directory structure
- Hello World
- Basic Go commands (`run` `build` `get` `install` `test`, `fmt`)
- Binary size. Go vs C
- Section 1 recap
- [Quiz](https://github.com/steevehook/udemy-go101/blob/master/section_1-introduction/quiz)

## Section 2 - Language Basics (I)

- Section 2 overview
- Basic Go commands
- Packages (executable & non executable)
- Executable & non-executable packages
- Basic types
- **const** & **var** declarations
- `_` blank identifier
- Reserved keywords & operators
- Aggregate types overview
- Arrays
- Slices
- **len** function
- **cap** function
- **append** function
- **make** function
- Arrays/Slices addressing (range selections)
- Control structures overview (if, else, switch, for)
- if/else statement
- switch statement
- for statement
- Arrays/Slices looping
- Maps
- Maps - addressing & looping
- Maps - comma, ok syntax
- Nil map
- Maps - delete function
- Map implementation (Bucket example)
- Zero values
- Section 2 recap
- [Quiz](https://github.com/steevehook/udemy-go101/blob/master/section_2-language-basics-1/quiz)

## Section 3 - Language basics (II)

- Section 3 overview
- Strucs
- Structs - exported/unexported
- Structs - field promotion
- Empty struct
- Functions
- Function - multiple return (data, error)
- Closures
- Custom types
- Methods (func receivers)
- Pointers
- Pointer vs value (Everything is passed by value)
- Marshalling & Unmarshalling (Overview)
- Struct tags
- Structs - JSON
- Arrays/Slices JSON
- Maps JSON
- JSON Unmarshalling
- Section 3 recap
- [Quiz](https://github.com/steevehook/udemy-go101/blob/master/section_2-language-basics-2/quiz)
- [Assignment - Problem](https://github.com/steevehook/udemy-go101/blob/master/section_2-language-basics-2/assignment)
- [Assignment - Solution](https://github.com/steevehook/udemy-go101/blob/master/section_2-language-basics-2/assignment)

## Section 4 - Notes CLI App

- Introduction to Notes App (CLI app design & expectations)
- Intro to `os.Args`
- Read CLI args
- Create command switch
- Intro to `os` package (Read/Write/Delete file)
- Create note (write to file)
- Read note (read from file)
- Delete note (read from file & delete & write to file)
- Intro to `os.Signal`
- Graceful app shutdown
- Section 4 recap

## Section 5 - Interfaces

- Section 5 overview
- Interfaces overview
- STD lib interfaces (Error, Stringer, Handler)
- Reader
- Writer
- Handler
- interface{} type
- Type assertion
- Custom HTTP transport
- Section 5 recap
- [Quiz](https://github.com/steevehook/udemy-go101/blob/master/section_5-interfaces/quiz)
- [Assignment - Problem](https://github.com/steevehook/udemy-go101/blob/master/section_5-interfaces/assignment)
- [Assignment - Solution](https://github.com/steevehook/udemy-go101/blob/master/section_5-interfaces/assignment)

## Section 6 - Explore `net` & `net/http`

- Section 6 overview
- Explore `net` package
- Create a TCP server
- Create a UDP server
- Client/Server architecture
- net client & `net.Dial`
- Intro to HTTP protocol
- HTTP server with `net`
- Create HTTP router/mux
- Intro to query & path params 
- Intro to RegEx
- Add UUID request param
- Explore `net/http` package
- Create an HTTP client
- HTTP server with `net/http`
- Handler, HandlerFunc, Handle
- HTTP client/server example
- Section 6 recap
- [Quiz](https://github.com/steevehook/udemy-go101/blob/master/section_6-exploring-net-http/quiz)
- [Assignment - Problem](https://github.com/steevehook/udemy-go101/blob/master/section_6-exploring-net-http/assignment)
- [Assignment - Solution](https://github.com/steevehook/udemy-go101/blob/master/section_6-exploring-net-http/assignment)

## Section 7 - Notes CLI app improvements

- Section 7 overview
- Create a MemCache DB (`Get` & `Set` -> `map[uuid]Note`)
- Connect MemCache to Notes (`create` `read` `delete` without backup)
- Intro to mutexes (Overview & examples)
- Add mutex to Get and Set
- Background backup (once 10s)
- Improved backup - overview (add slice & uuid)
- Intro to UUID (`UUID.V4` `UUID.V5`)
- Intro to `satori/go.uuid`
- Add UUID to MemCache
- Improve backup with `[]UUID`
- Section 7 recap

## Section 8 - Errors

- Section 8 overview
- Apps & Errors design
- Errors in Go. (val, err) pattern
- Error interface & Custom errors
- **defer**
- **panic**
- **recover**
- Error types & Error values
- Panic recovery middleware
- Error checking & examples
- Type assertion recap
- Error handling good practices
- `Go 1.13` error features
- Section 8 recap
- [Quiz](https://github.com/steevehook/udemy-go101/blob/master/section_8-errors/quiz)
- [Assignment - Problem](https://github.com/steevehook/udemy-go101/blob/master/section_8-errors/assignment)
- [Assignment - Solution](https://github.com/steevehook/udemy-go101/blob/master/section_8-errors/assignment)

## Section 9 - Refactor Notes CLI app

- Section 9 overview
- Code smells & solutions
- Intro to SOLID
- Code split (files & packages)
- Avoid global vars
- Decoupling & func helpers
- From `File` to `ReadWriteCloser`
- Document exported symbols
- Section 9 recap

## Section 10 - Testing
- Section 10 overview
- Testing in Go (overview)
- Testing functions
- Testing types
- The power of interfaces (simple interfaces, simple implementations)
- Testify - test suites & mocks
- Mocks, stubs, fixtures
- Table driven testing
- Golden files
- Testing dependencies & Mocks
- Mock API external calls
- Mock database calls
- Custom build tags for testing
- Unit & integration testing
- Parallel testing
- Benchmark testing
- Test coverage script
- Section 10 recap
- [Quiz](https://github.com/steevehook/udemy-go101/blob/master/section_10-testing/quiz)
- [Assignment - Problem](https://github.com/steevehook/udemy-go101/blob/master/section_10-testing/assignment)
- [Assignment - Solution](https://github.com/steevehook/udemy-go101/blob/master/section_10-testing/assignment)

## Section 11 - Test Notes CLI app

- Section 11 overview
- Test **create**
- Test **read**
- Test **delete**
- Test **shutdown**
- Test **backup**
- Test MemCache
- Integration test **create**
- Integration test **read**
- Integration test **delete**
- Integration test **shutdown**
- Integration test **backup**
- Section 11 recap

## Section 12 - Expenses REST API

- Intro to Expenses REST API (Overview + list all dependencies)
- Install Expenses REST API deps
- Intro to **client**/**server** architecture
- Intro to REST (resources & verbs)
- Create a small REST API (with 2 endpoints)
- Intro to Databases. **SQL** vs **NoSQL**
- Intro to containerization
- Docker basic commands (`run` `ps` `pull` `exec` `logs` `rm` `build`)
- Intro to MariaDB & SQL
- Basic MariaDB CLI commands (CRUD)
- MariaDB Go example
- Intro to `ORM` (Object Relational Mapping)
- `ORM` vs `DAL` (Data Access Layer)
- Intro to `upper.io/db.v3`
- `MVC` overview (Model View Controller)
- `MCSR` overview (**Middleware**/**Controller**/**Service**/**Repo**)
- `CRUD` & API endpoints overview (**Create** **Read** **Update** **Delete**)
- Intro to `julienschmidt/httprouter`
- Create the app router (with endpoint placeholders)
- Create `GET /expenses`
- Intro to Postman. Test 1st endpoint (JSON Postman collection with all endpoints)
- Create `GET /expenses/{id}`
- Create `POST /expenses`
- Create `PATCH /expenses/{id}`
- Create `DELETE /expenses/{id}`
- Create **Expenses** service
- Create `notes` schema
- Intro to **Repository** pattern
- Create **Notes** repository
- Intro to `spf13/viper`
- Read from config file
- `Makefile` intro
- Add `build` `vendor` `clear` & `test` commands
- Section 12 recap

## Section 13 - Test Expenses REST API

- Section 13 overview
- Create service mocks
- Test `GET /expenses`
- Test `GET /expenses/{id}`
- Test `POST /expenses`
- Test `PATCH /expenses/{id}`
- Test `DELETE /expenses/{id}`
- Create repository mocks
- Test **Expenses** service
- Intro to `DATA-DOG/go-sqlmock`
- Test **Notes** repository
- Test `config` package
- Create app background startup
- Integration test `GET /expenses`
- Integration test `GET /expenses/{id}`
- Integration test `POST /expenses`
- Integration test `PATCH /expenses/{id}`
- Integration test `DELETE /expenses/{id}`
- Section 13 recap

## Section 14 - Expenses REST API Auth

- Section 14 overview
- Intro to middleware
- Create Auth middleware
- Intro to `justinas/alice`
- Create the middleware chain
- Connect **Auth** middleware
- Connect **Panic** middleware
- Intro to **JWT** & **OAuth2** (Overview)
- Intro to **Bearer** Token (HTTP headers)
- Use **JWT** in Auth middleware
- Create **JWT** test helper
- Test Auth middleware
- Adjust failing tests
- Create `POST /login`
- Test `POST /login`
- Create `GET /logout`
- Test `GET /logout`
- Create `GET /me`
- Test `GET /me`
- Integration test `POST /login`
- Integration test `GET /logout`
- Integration test `GET /me`
- Test all routes with Postman
- Future **OAuth2** & **JWT** considerations
- Section 14 recap

## Section 15 - Expenses REST API Session

- Section 15 overview
- Intro to **HTTP** session
- Intro to **Cookies**
- Create in-memory HTTP session
- Intro `text/template` & `html/template`
- Store HTTP session in cookie
- Persistent (stateless) HTTP session (Design & Overview)
- Intro to Redis (basic types & CLI commands - `GET`, `SET`, `HGET`, `HGETALL`, `HSET`, `HMSET`)
- Redis & Go example
- Intro to `docker-compose`
- Create **Session** repository
- Create **AddSession**
- Test **AddSession**
- Create **DeleteSession**
- Test **DeleteSession**
- **AddSession** on login
- **DeleteSession** on logout
- Adjust integration tests for `POST /login`
- Adjust integration tests for `GET /logout`
- Test auth routes with Postman
- Section 15 recap

## Section 16 - Package management

- Section 16 overview
- Explaining packages in depth
- Package management & go get
- Go mod (vgo)
- Intro to Dep
- Migration to Go mod
- Migrate Expenses REST API to Dep
- Migrate Expenses REST API to Go modules
- Section 16 recap
- [Quiz](https://github.com/steevehook/udemy-go101/blob/master/section_15-package-management/quiz)
- [Assignment - Problem](https://github.com/steevehook/udemy-go101/blob/master/section_15-package-management/assignment)
- [Assignment - Solution](https://github.com/steevehook/udemy-go101/blob/master/section_15-package-management/assignment)

## Section 17 - Concurrency

- Section 17 overview
- Concurrency vs Parallelism
- Channels
- Buffered channels
- Send/Receive only channels
- Concurrent func(s). The `go` keyword
- Concurrent func(s) facts
- for select
- Deadlock
- Race conditions & `-race` flag
- wait groups
- Mutexes
- Atomics
- Semaphore with channels
- Concurrency patterns
- Channels - FAN IN/OUT
- **Health Checker** example (HTTP round tripper)
- Intro to `gopkg.in/h2non/gock.v1`
- Test **Health Checker**
- Notes CLI app - Concurrency (Use multiple connections)
- Section 17 recap
- [Quiz](https://github.com/steevehook/udemy-go101/blob/master/section_16-concurrency/quiz)
- [Assignment - Problem](https://github.com/steevehook/udemy-go101/blob/master/section_16-concurrency/assignment)
- [Assignment - Solution](https://github.com/steevehook/udemy-go101/blob/master/section_16-concurrency/assignment)

## Section 18 - Real Time Chat

- Small overview of the chat app
- Introduction to web sockets
- Write the **Hub**
- Write the **Server**
- Write the **Client**
- Introduction to Go templates
- Glue the web client (copy paste the CSS & JS)
- Test the application & make sure it all works
- Introduction to MongoDB
- Basic MongoDB CLI commands (CRUD)
- MongoDB Go example
- Create **Messages** repository
- Connect the **Server** with **Messages** repository
- Improvements (Test the application)
- Section 18 recap

## Section 19 - Platform Specific Software

- Section 19 overview
- Intro to Go special files
- Intro to build tags
- `$GOOS` & `$GOARCH`
- Generate binaries checksum
- Test binaries (OSX, Linux, Windows, Raspberry Pi)
- Section 19 recap
- [Quiz](https://github.com/steevehook/udemy-go101/blob/master/section_18-platform-specific-software/quiz)
- [Assignment - Problem](https://github.com/steevehook/udemy-go101/blob/master/section_18-platform-specific-software/assignment)
- [Assignment - Solution](https://github.com/steevehook/udemy-go101/blob/master/section_18-platform-specific-software/assignment)

## Section 20 - Debugging

- Section 20 overview
- Intro to **Delve**
- Memory Footprint
- CPU usage
- Intro to `runtime/trace` & `x/net/trace`
- Intro to `pprof`
- Intro to `ldflags`
- Section 20 recap
- [Quiz](https://github.com/steevehook/udemy-go101/blob/master/section_19-debugging/quiz)

## Section 21 - Deploying Go apps

- Section 21 overview
- Introduction to SystemD
- Deploy to DigitalOcean
- Deploy to Amazon
- Deploy to Google Cloud
- Deploy to Heroku
- Go custom Docker image
- Add docker image to `docker-compose.yaml`
- Deploy **Real Time Chat** to Amazon
- Add Travis
- Into to Jenkins
- Section 21 recap
- [Quiz](https://github.com/steevehook/udemy-go101/blob/master/section_20-deploying-go-apps/quiz)
- [Assignment - Problem](https://github.com/steevehook/udemy-go101/blob/master/section_20-deploying-go-apps/assignment)
- [Assignment - Solution](https://github.com/steevehook/udemy-go101/blob/master/section_20-deploying-go-apps/assignment)

## Section 22 - Go best practices

- Section 22 overview
- CodeReviewComments
- Write docs
- Better copy paste than write wrong abstraction
- Avoid interface{} as much as you can
- Handle panics gracefully
- Naming packages
- Naming variables
- Always handle errors
- Use a logger
- Always check for race conditions
- Splitting packages/files
- Section 22 recap
- [Quiz](https://github.com/steevehook/udemy-go101/blob/master/section_21-go-best-practices/quiz)

## Farewell & Resources

- Resources, social media contacts (GopherTuts) & Good bye (fare well)

[Back](https://github.com/steevehook/udemy-go101)

---

Happy hacking gophers 🚀🚀🚀

<img src="https://github.com/steevehook/udemy-go101/raw/master/udemy-go101.svg?sanitize=true" width="150px"/>
