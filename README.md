# Go Lang Practice Projects 🚀

A collection of small, focused Go projects designed to help you learn **Go fundamentals**, **CLI development**, and **core standard library features** in a practical, hands-on way.

Each project lives in its own folder with an independent `go.mod`, making them easy to run, understand, and extend.

---

## 📁 Repository Structure

```text
.
├── go_calc/
│   ├── calculator.go
│   └── go.mod
│
├── go_guess/
│   ├── numbGuess.go
│   └── go.mod
│
└── README.md
```

---

## ✅ Prerequisites

Before getting started, make sure you have:

- **Go 1.20+** installed  
  👉 https://go.dev/dl/
- **Git**
- Basic terminal / command-line knowledge

**Verify installation:**

```sh
go version
```

---

## 🚀 Getting Started

**Clone the repository:**

```sh
git clone <repo-url>
cd <repo-name>
```

Each project is self-contained, so you run them individually from their folder.

---

## ▶️ Running a Project

### Example: CLI Calculator

```sh
cd go_calc
go run calculator.go
```

### Example: Number Guessing Game

```sh
cd go_guess
go run numbGuess.go
```

---

## 🏗️ Build a Project (Optional)

You can build a binary instead of running directly:

```sh
go build
```

This produces an executable you can run:

```sh
./go_calc
```

---

## 📚 Learning Goals

This repository focuses on learning Go through small, incremental projects, covering:

- Idiomatic Go syntax
- CLI applications
- Error handling
- Structs, slices, and maps
- File I/O
- Concurrency with goroutines & channels
- Networking and HTTP servers

---

## 🧩 Project List & Roadmap

### 1. CLI Calculator ✅

**Status:** Implemented (`go_calc`)

**Features:**
- Supports: `+ - * /`
- Handles invalid input
- Prevents division by zero

**Concepts Learned:**
- Variables & data types
- `fmt.Scan`, `fmt.Printf`
- Conditionals and loops
- Basic error handling

---

### 2. Number Guessing Game 🎮

**Status:** Implemented (`go_guess`)

**Features:**
- Random number generation
- Repeated user guesses
- Attempt counter

**Concepts Learned:**
- `math/rand`
- Loops
- User input validation
- Program flow control

---

### 3. To-Do List CLI App 📝

**Status:** Planned

**What to Build:**
- Add / list / delete tasks
- Persist tasks using JSON file storage

**Concepts:**
- Structs
- Slices & maps
- File I/O
- JSON encoding / decoding

---

### 4. Word Count Tool 📄

**Status:** Planned

**What to Build:**
- Read `.txt` files
- Count words, lines, and characters

**Concepts:**
- `bufio.Scanner`
- String manipulation
- Maps for counting

---

### 5. Password Generator 🔐

**Status:** Planned

**What to Build:**
- Secure password generation
- Configurable length & symbols

**Concepts:**
- `crypto/rand`
- Byte slices
- CLI flags (`flag` package)

---

### 6. REST API (CRUD) 🚀

**Status:** Planned

**What to Build:**
- Basic REST API (users or products)
- GET, POST, PUT, DELETE

**Concepts:**
- `net/http`
- JSON APIs
- Struct tags
- Request / response lifecycle

---

### 7. URL Shortener 🌐

**Status:** Planned

**What to Build:**
- Short URL generation
- Redirect handling
- In-memory or Redis storage

**Concepts:**
- HTTP routing
- Hashing / ID generation
- Concurrency safety
- Middleware concepts

---

### 8. Concurrent Web Scraper 🕷️

**Status:** Planned

**What to Build:**
- Fetch page titles concurrently
- Handle multiple URLs efficiently

**Concepts:**
- Goroutines
- Channels
- `sync.WaitGroup`
- Concurrency patterns

---

### 9. Log Analyzer 📊

**Status:** Planned

**What to Build:**
- Parse server log files
- Count status codes & IPs

**Concepts:**
- File streaming
- Maps
- Sorting
- Performance considerations

---

### 10. Mini Chat Server 💬

**Status:** Planned (Intermediate)

**What to Build:**
- TCP or WebSocket chat server
- Multiple connected clients

**Concepts:**
- Networking (`net`)
- Goroutines per client
- Channels for broadcasting
- Graceful shutdowns

---

## 🧪 Testing (Future)

Each project may include tests using:

```sh
go test ./...
```

---

## 🤝 Contributing

Contributions are welcome!

- Open an issue or pull request
- Follow `gofmt`
- Add tests for new functionality
- Keep projects simple and educational

---

## 📜 License

This project is licensed under the MIT License (or your preferred license).  
Add a `LICENSE` file at the root of the repository.

---

## ✨ Notes

- This repository is intentionally beginner-focused and practical.
- The goal is clarity and learning — not perfection.

**Happy coding with Go!** 🐹🔥