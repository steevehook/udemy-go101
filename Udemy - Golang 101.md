# Udemy - Go 101 Plan

Estimated course time >= 20 hours

## Chapter 1 - Introduction

- Course Introduction & About the author
- Course overview. Chapter 1 overview
- Course rules & useful resources (Q&A, Documentation, code sharing & others, Discord)
- Go in a nutshell. Go authors & philosophy
- Go installation on Mac
- Go installation on Linux
- Go installation on Windows
- Git installation on Mac/Linux
- Git installation on Windows
- Git basic commands & Github
- Cloning the course from GitHub & small directories overview
- $GOPATH & directory structure (project layout)
- Hello World
- Chapter 1 recap

## Chapter 2 - Language Basics (I)

- Small overview on Chapter 2
- Explaining the basic & important Go commands
- Packages
- Package types (executable & non executable)
- Basic Types in Go
- **const** & **var** declarations
- Reserved keywords & operators
- Small overview on aggregate types
- Arrays
- Slices
- Explaining the **len** function
- Explaining the **cap** function
- Explaining the **append** function
- Explaining the **make** function
- Arrays/Slices addressing
- Arrays/Slices looping
- Maps
- Maps looping
- Maps - comma, ok syntax
- Nil map
- Maps - delete function
- Map implementation (Bucket)
- Zero values
- Control structures (if, else, switch, for) overview
- if/else statement
- switch statement
- for statement
- Assignment - Problem
- Assignment - Solution
- Chapter 2 recap

## Chapter 3 - Language basics (II)

- Small overview on Chapter 3
- Strucs
- Structs - exported / unexported
- Structs - field promotion
- Functions
- Function - (data, error) pattern
- Closures
- Custom types
- Methods (function receivers)
- Pointers
- Pointer vs value (Everything is passed by value)
- Marshalling vs Unmarshalling (Overview)
- Structs JSON
- Arrays/Slices JSON
- Maps JSON
- JSON Unmarshalling
- Channels
- Buffered channels
- Send only / receive only channels
- Interfaces - general overview
- Interfaces in Std Lib
- Interface{} type
- Implement some STD lib interfaces & use them
- Interfaces - Type Assertion
- Assignment - Problem
- Assignment - Solution
- Chapter 3 recap

## Chapter 4 - Notes CLI App

- Introduction to Notes App
- Create a CLI interface
- Read OS args
- Create read command
- Create write command
- Exploring the net package
- Create a memory cache database
- Connect the memory cache
- Introduction to Mutexes (Overview & examples)
- Fine tune the app - (add mutexes)
- Notes backup improvement (explanation)
- Write the backup utility
- Introduction to `os.Signal`
- Make backup before program exits
- Chapter 4 recap

## Chapter 5 - Errors

- Small overview of Chapter 5
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
- Chapter 5 recap

## Chapter 6 - Testing
- Small overview on Chapter 6
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
- Chapter 6 recap

## Chapter 7 - Testing & Refactoring Notes App

- Small overview on Chapter 7
- Code smells & bad architecture design & solutions
- SOLID - explanation
- Explain how the app refactoring
- Replace implementations with abstract interfaces
- Extract big functionalities in function helpers
- Create separate types & unit test them individually
- Unit test each component
- Write docs for exported symbols
- Integration test the whole CLI app
- Chapter 7 recap

## Chapter 8 - Expenses REST API

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
- Chapter 7 recap

## Chapter 9 - Expenses REST API Auth

- Small overview on Chapter 9
- Introduction to router middleware
- Create Auth middleware
- Guard private routes with Auth middleware
- Introduction to JWT (Overview)
- Adjust Auth middleware to use JWT
- Adjust all failing tests, after implementing Auth middleware
- Testing all routes using Postman and Auth bearer token
- Assignment (add auth middleware & tests for some private routes)
- Assignment solution
- Chapter 9 recap

## Chapter 10 - Package management

- Small overview on Chapter 10
- Explaining packages in depth
- Package management & go get
- Go mod (vgo)
- Migration to Go mod
- Migrate Expenses REST API to Go modules
- Chapter 10 recap

## Chapter 11 - Concurrency

- Small overview on Chapter 11
- Concurrency is not Parallelism. 1 CPU concurrency vs multi CPU concurrency
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
- Chapter 8 recap

## Chapter 12 - Real Time Chat

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

## Chapter 13 - Deploying Go apps

- Small overview on Chapter 13
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
- Chapter 13 recap

## Chapter 14 - Go best practices

- Small overview on Chapter 13
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
- Chapter 14 recap

## Farewell & Resources

- Resources, social media contacts (GopherTuts) & Good bye (fare well)
