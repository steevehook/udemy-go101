# Udemy - Go 101 Curriculum

Estimated course time ~28 hours

## Section 1 - Introduction

- Course Introduction & About the author
- Course rules & useful resources (Q&A, Documentation, code sharing & others, Discord)
- Go in a nutshell. Go authors & philosophy
- Installation on Mac (Go & Git)
- Installation on Windows (Go & Git)
- Installation on Linux (Go & Git)
- Git basics & Github (clone udemy-go101) 
- $GOPATH & directory structure
- Hello World
- Section 1 recap
- Section 1 quiz

## Section 2 - Language Basics (I)

- Section 2 overview
- Basic Go commands
- Packages (executable & non executable)
- Executable & non-executable packages
- Basic Types
- **const** & **var** declarations
- Reserved keywords & operators
- Aggregate types overview
- Arrays
- Slices
- **len** function
- **cap** function
- **append** function
- **make** function
- Arrays/Slices addressing
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
- Map implementation (Bucket)
- Zero values
- Section 2 recap
- Section 2 quiz

## Section 3 - Language basics (II)

- Section 3 overview
- Strucs
- Structs - exported / unexported
- Structs - field promotion
- Functions
- Function - multiple return (data, error)
- Closures
- Custom types
- Methods (function receivers)
- Pointers
- Pointer vs value (Everything is passed by value)
- Marshalling & Unmarshalling (Overview)
- Structs JSON
- Arrays/Slices JSON
- Maps JSON
- JSON Unmarshalling
- Section 3 recap
- Section 3 quiz

## Section 4 - Notes CLI App

- Introduction to Notes App (CLI app design & expectations)
- Intro to `os.Args`
- Read CLI args
- Create command switch
- Intro to `os` package (Read/Write/Delete file)
- Create note
- Read note
- Delete note
- Intro to `archive/zip` package (Zip/unzip files)
- Backup notes
- Ingest notes
- Intro to `os.Signal`
- Graceful app shutdown
- Section 4 recap

## Section 5 - Interfaces

- Section 5 overview
- Interfaces - general overview
- Interfaces in Std Lib
- Interface{} type
- Implement some STD lib interfaces & use them
- Interfaces - Type Assertion

## Section 6 - Exploring `net` & `net/http`

- Section 6 overview
- Exploring the `net` package

## Section 7 - Notes CLI app improvements

- Section 7 overview
- Create a memory cache database
- Connect the memory cache
- Introduction to Mutexes (Overview & examples)
- Fine tune the app - (add mutexes)
- Background backup
- Backup locking

## Section 8 - Errors

- Section 8 overview
- Errors in Go. (val, err) pattern
- Error interface & Custom errors
- What is defer? Example
- What is panic? Example
- Panic & recover example
- Panic recovery middleware
- Type assertion recap
- Error checking & examples
- Build a custom HTTP transport
- Error handling good practices
- Assignment (Problem)
- Assignment (Solution)
- Section 5 recap

## Section 9 - Refactor Notes CLI app

- Section 9 overview

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
- Assignment (Test a piece of code with dependencies)
- Assignment (Solution)
- Section 6 recap

## Section 11 - Testing Notes CLI app

- Section 11 overview
- Small overview on Section 11
- Code smells & bad architecture design & solutions
- SOLID - explanation
- Explain how the app refactoring
- Replace implementations with abstract interfaces
- Extract big functionalities in function helpers
- Create separate types & unit test them individually
- Unit test each component
- Write docs for exported symbols
- Integration test the whole CLI app
- Section 7 recap

## Section 12 - Expenses REST API

- Introduction to Expenses REST API (Overview)
- Introduction to client - server architecture
- Introduction to HTTP protocol
- Introduction to REST (resources & verbs example)
- Introduction to net/http package
- Create a small HTTP server with 2 REST endpoints
- Introduction to Databases. SQL vs NoSQL
- Introduction to containerization
- Introduction to Docker - Basic commands
- Introduction to MariaDB & SQL
- Basic MariaDB CLI commands (CRUD)
- MariaDB Go example
- Docker MariaDB tinkering
- Configure the app structure & dependencies
- MVC overview
- Middleware/Controller/Service/DB Layer overview
- CRUD explanation & API endpoints overview
- Create GET /expenses
- Test GET /expenses
- Introduction to Postman. Testing the first endpoint
- Create GET /expenses/{id}
- Test GET /expenses/{id}
- Create POST /expenses
- Test POST /expenses
- Create PATCH /expenses/{id}
- Test PATCH /expenses/{id}
- Create DELETE /expenses/{id}
- Test DELETE /expenses/{id}
- Database layer. Repository pattern
- Create Notes repository
- Test Notes repository
- Read from config file
- Testing the config file
- Add integration tests for Notes REST API
- Assignment (Finish couple of endpoints & test them)
- Section 7 recap

## Section 13 - Expenses REST API Auth

- Section 13 overview
- Introduction to router middleware
- Create Auth middleware
- Guard private routes with Auth middleware
- Introduction to JWT (Overview)
- Adjust Auth middleware to use JWT
- Adjust all failing tests, after implementing Auth middleware
- Testing all routes using Postman and Auth bearer token
- Assignment (add auth middleware & tests for some private routes)
- Assignment solution
- Section 9 recap

## Section 14 - Package management

- Section 14 overview
- Explaining packages in depth
- Package management & go get
- Go mod (vgo)
- Migration to Go mod
- Migrate Expenses REST API to Go modules
- Section 10 recap

## Section 15 - Concurrency

- Section 15 overview
- Concurrency is not Parallelism. 1 CPU concurrency vs multi CPU concurrency
- Channels
- Buffered channels
- Send only / receive only channels
- Concurrent functions. The `go` keyword
- Concurrent functions facts
- for select
- wait groups
- Mutexes
- Atomics
- Buffered vs unbuffered channels
- Concurrency patterns
- Channels - FAN IN / FAN OUT
- Common concurrency issues & race conditions
- -race flag
- Health check example (HTTP roundtripper)
- Enhance Notes App to use multiple connections
- Assignment
- Section 8 recap

## Section 16 - Real Time Chat

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

## Section 17 - Deploying Go apps

- Section 17 overview
- Introduction to SystemD
- DigitalOcean
- Amazon
- Google Cloud
- Heroku
- Docker
- Deploy "Real Time Chat" app
- Jenkins
- Assignment (Deploy Expenses REST API)
- Assignment (Solution)
- Section 13 recap

## Section 18 - Go best practices

- Section 18 overview
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
- Section 14 recap

## Farewell & Resources

- Resources, social media contacts (GopherTuts) & Good bye (fare well)

[Back](https://github.com/steevehook/udemy-go101)

---

Happy hacking gophers 🚀🚀🚀

<img src="https://github.com/steevehook/udemy-go101/raw/master/udemy-go101.svg?sanitize=true" width="150px"/>