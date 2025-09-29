# Go Programming Tutorial

A comprehensive tutorial for learning Go (Golang) from the ground up. This guide covers everything from basic syntax to advanced concurrency patterns with practical, runnable examples.

## 📚 Core Topics Covered

### Fundamentals
- **Program Structure** - Understanding packages, imports, and the main function
- **Variables & Data Types** - Declarations, type inference, constants, and basic types
- **Control Structures** - If statements, switch cases, and loops
- **Functions** - Parameters, return values, variadic functions, and named returns

### Data Structures
- **Arrays** - Fixed-size collections with type safety
- **Slices** - Dynamic arrays with powerful built-in operations
- **Maps** - Key-value pairs for efficient lookups
- **Structs** - Custom data types and composition

### Object-Oriented Concepts
- **Methods** - Attaching behavior to types
- **Interfaces** - Defining contracts and polymorphism
- **Embedded Structs** - Code reuse through composition
- **Pointer Receivers** - Modifying values in methods

### Advanced Features
- **Error Handling** - Idiomatic error patterns and custom errors
- **Goroutines** - Lightweight concurrent execution
- **Channels** - Safe communication between goroutines
- **Select Statement** - Multiplexing channel operations
- **WaitGroups** - Synchronizing concurrent operations

### Code Organization
- **Packages** - Creating and using custom packages
- **Modules** - Dependency management with go.mod
- **Exported vs Unexported** - Controlling API visibility
- **External Libraries** - Using third-party packages

## ✨ Key Features

- ✅ **Complete Examples** - Every concept includes working code you can run immediately
- ✅ **Progressive Learning** - Builds from simple to complex topics logically
- ✅ **Best Practices** - Follows idiomatic Go conventions and patterns
- ✅ **Real-World Patterns** - Practical examples you'll use in actual projects
- ✅ **Concurrent Programming** - Deep dive into Go's powerful concurrency model
- ✅ **Hands-On Approach** - Learn by doing with executable code samples
- ✅ **Clear Explanations** - Commented code with detailed descriptions
- ✅ **Next Steps** - Guidance on advanced topics and project ideas

## 🚀 Getting Started

### Prerequisites

1. **Install Go** - Download from [golang.org/dl](https://golang.org/dl)
2. **Verify Installation**
   ```bash
   go version
   ```
   You should see output like: `go version go1.21.0 linux/amd64`

### Setting Up Your Project

1. **Create a Project Directory**
   ```bash
   mkdir go-tutorial
   cd go-tutorial
   ```

2. **Initialize a Go Module**
   ```bash
   go mod init go-tutorial
   ```
   This creates a `go.mod` file for dependency management.

3. **Create Your First Program**
   ```bash
   touch main.go
   ```

## 🏃 How to Run the Code

### Running a Single File

For most examples in this tutorial, you'll create a `main.go` file:

```bash
go run main.go
```

**Example:**
```go
// main.go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

Run it:
```bash
go run main.go
```

Output:
```
Hello, Go!
```

### Building an Executable

To compile your program into a binary:

```bash
go build main.go
```

This creates an executable file. Run it:

```bash
# On Linux/Mac
./main

# On Windows
main.exe
```

### Running with Build Flags

Build with optimizations:
```bash
go build -ldflags="-s -w" main.go
```

### Running Multiple Files

If your program spans multiple files:

```bash
go run .
```

Or specify all files:
```bash
go run main.go utils.go
```

### Working with Packages

**Project structure:**
```
go-tutorial/
├── go.mod
├── main.go
└── mathutils/
    └── mathutils.go
```

**Run the main program:**
```bash
go run main.go
```

Go automatically handles local package imports based on your `go.mod`.

### Installing External Dependencies

When using external packages:

```bash
# Install a specific package
go get github.com/gorilla/mux

# Install all dependencies in go.mod
go mod download

# Tidy up dependencies
go mod tidy
```

Then run normally:
```bash
go run main.go
```

## 📖 Using This Tutorial

### Step-by-Step Learning Path

1. **Start with Basics** (Day 1)
   - Your First Go Program
   - Variables and Data Types
   - Control Structures

2. **Functions and Data Structures** (Day 2)
   - Functions
   - Arrays and Slices
   - Maps

3. **Object-Oriented Concepts** (Day 3)
   - Structs and Methods
   - Interfaces
   - Error Handling

4. **Concurrency** (Day 4)
   - Goroutines
   - Channels
   - Select and WaitGroups

5. **Code Organization** (Day 5)
   - Packages and Modules
   - Working with External Libraries

### How to Practice

1. **Read the Concept** - Understand the explanation
2. **Type the Code** - Don't copy-paste, type it yourself
3. **Run the Example** - Use `go run main.go`
4. **Modify and Experiment** - Change values, add features
5. **Break Things** - Intentionally make errors to learn debugging

### Quick Tips

- 💡 Use `go fmt` to automatically format your code
- 💡 Run `go vet` to catch common mistakes
- 💡 Try code in the [Go Playground](https://play.golang.org/) for quick experiments
- 💡 Read error messages carefully - Go's errors are usually clear and helpful

## 🛠️ Common Commands Reference

| Command | Description |
|---------|-------------|
| `go run main.go` | Compile and run a Go program |
| `go build` | Compile the current package |
| `go install` | Compile and install packages |
| `go test` | Run tests |
| `go fmt` | Format Go source code |
| `go vet` | Report suspicious constructs |
| `go mod init` | Initialize a new module |
| `go mod tidy` | Add missing and remove unused modules |
| `go get <package>` | Download and install packages |
| `go doc <package>` | Show documentation for a package |

## 📂 Project Structure Example

```
go-tutorial/
├── go.mod                  # Module definition and dependencies
├── go.sum                  # Dependency checksums
├── main.go                 # Main application entry point
├── README.md               # This file
│
├── mathutils/              # Custom package
│   └── mathutils.go
│
├── models/                 # Data models
│   ├── user.go
│   └── product.go
│
└── handlers/               # HTTP handlers or business logic
    ├── user_handler.go
    └── product_handler.go
```

## 🎯 Next Steps

After completing this tutorial, explore:

### Advanced Topics
- Testing with the `testing` package
- Benchmarking and profiling
- Context for cancellation and deadlines
- Reflection and type assertions
- Custom error types with stack traces

### Web Development
- Build REST APIs with Gin or Echo
- Create web applications with HTML templates
- Work with databases using GORM
- Implement authentication and middleware

### Systems Programming
- File I/O and OS interactions
- Network programming with TCP/UDP
- Command-line tools with Cobra
- System monitoring and metrics

### Project Ideas
1. **TODO CLI App** - Command-line task manager
2. **REST API** - User management service with database
3. **Web Scraper** - Concurrent website data collector
4. **Chat Server** - Real-time messaging with WebSockets
5. **Microservice** - Small focused service with API endpoints

## 📚 Additional Resources

- [Official Go Documentation](https://golang.org/doc/)
- [Effective Go](https://golang.org/doc/effective_go.html) - Style guide
- [Go by Example](https://gobyexample.com/) - Annotated example programs
- [Go Playground](https://play.golang.org/) - Run Go code in your browser
- [Go Standard Library](https://pkg.go.dev/std) - Complete stdlib documentation

## 🤝 Contributing

Feel free to improve this tutorial by:
- Adding more examples
- Fixing errors or typos
- Suggesting better explanations
- Contributing advanced topics

## 📄 License

This tutorial is open source and available for learning purposes.

---

**Happy Coding with Go! 🚀**

Start with `go run main.go` and enjoy the journey!