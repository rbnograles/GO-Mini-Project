# Go Project

Small, idiomatic Go repository template with commands for building, testing, linting, and releasing.

## Prerequisites

- Go 1.20+ installed
- git

## Getting started

1. Clone
    git clone <repo-url>
2. Enter
    cd <repo-dir>
3. Initialize module (if not present)
    go mod init github.com/username/project-name

## Build

- Build all packages:
  go build ./...
- Build a command in `cmd/<name>`:
  go build -o bin/<name> ./cmd/<name>

## Run

- Run a single package:
  go run ./cmd/<name>
- Run tests:
  go test ./... -v

## Projects

1. CLI Calculator ✅ (Beginner)

What to build

A command-line calculator that supports + - * /

Handle invalid input and division by zero

You’ll learn

Variables & types

fmt.Scan, fmt.Printf

Conditionals & loops

Basic error handling

2. Number Guessing Game 🎮

What to build

Program generates a random number

User keeps guessing until correct

Track attempts

You’ll learn

math/rand

Loops

User input validation

Simple program flow

3. To-Do List CLI App 📝

What to build

Add / list / delete tasks

Store tasks in memory or a file (JSON)

You’ll learn

Structs

Slices and maps

File I/O

JSON encoding/decoding

4. Word Count Tool 📄

What to build

Read a .txt file

Count words, lines, and characters

You’ll learn

File handling

String manipulation

bufio.Scanner

Maps for counting

5. Simple Password Generator 🔐

What to build

Generate secure random passwords

Options: length, include symbols, numbers

You’ll learn

crypto/rand

Working with byte slices

Flags (flag package)

6. REST API (CRUD) 🚀

What to build

REST API for users or products

Endpoints: GET, POST, PUT, DELETE

You’ll learn

net/http

JSON APIs

Struct tags

Basic request/response lifecycle

Tip: Start with in-memory data before adding a DB.

7. URL Shortener 🌐

What to build

Convert long URLs to short keys

Redirect on visit

Store mapping in memory or Redis

You’ll learn

HTTP routing

Hashing / ID generation

Maps & concurrency safety

Middleware basics

8. Concurrent Web Scraper 🕷️

What to build

Given a list of URLs, fetch titles in parallel

You’ll learn

Goroutines

Channels

WaitGroups

Real-world concurrency patterns

9. Log Analyzer 📊

What to build

Parse server logs

Count status codes, top IPs, busiest times

You’ll learn

File streaming

Regex (optional)

Maps & sorting

Performance thinking

10. Mini Chat Server 💬 (Advanced Beginner / Intermediate)

What to build

TCP or WebSocket-based chat server

Multiple clients can send messages

You’ll learn

Networking (net)

Goroutines per client

Channels for message broadcasting

Graceful shutdowns

## Contributing

- Open issues or PRs.
- Follow gofmt and run tests locally before submitting.
- Add tests for any new behavior.

## License

Add a LICENSE file (e.g., MIT, Apache 2.0) and reference it here.

---

Replace placeholders (repo URL, command names) with project-specific values.