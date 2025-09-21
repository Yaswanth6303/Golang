# Go Learning Playlist

A comprehensive collection of Go programming examples covering fundamental concepts, data structures, concurrency, and advanced features. This repository serves as a practical learning resource for mastering Go programming language.

## 📚 Topics Covered

### 1. **Getting Started**

-   [Hello World](1_hello_world/) - Basic Go program structure
-   [Simple Values](2_simple_values/) - Understanding basic data types

### 2. **Variables and Constants**

-   [Variables](3_variables/) - Variable declaration and usage
-   [Constants](4_constants/) - Constant declaration and usage

### 3. **Control Structures**

-   [For Loops](5_for/) - Loop constructs in Go
-   [If-Else](6_if_else/) - Conditional statements
-   [Switch](7_switch/) - Switch statements

### 4. **Data Structures**

-   [Arrays](8_arrays/) - Fixed-size arrays
-   [Slices](9_slices/) - Dynamic arrays
-   [Maps](10_maps/) - Key-value pairs
-   [Structs](16_structs/) - Custom data types

### 5. **Functions and Methods**

-   [Functions](12_functions/) - Function declaration and usage
-   [Variadic Functions](13_variadic_functions/) - Functions with variable arguments
-   [Closures](14_closures/) - Function closures
-   [Pointers](15_pointers/) - Memory addresses and pointers

### 6. **Advanced Features**

-   [Interfaces](17_interfaces/) - Interface implementation
-   [Enums](18_enums/) - Enumeration patterns
-   [Generics](19_generics/) - Generic programming

### 7. **Concurrency**

-   [Goroutines](20_goroutines/) - Lightweight threads
-   [Channels](21_channels/) - Communication between goroutines
-   [Mutex](22_mutex/) - Mutual exclusion for shared resources

### 8. **File Operations**

-   [Files](23_files/) - File reading and writing

### 9. **Packages and Modules**

-   [Packages](24_packages/) - Package organization and modules

## 🚀 Getting Started

### Prerequisites

-   Go 1.22.5 or higher
-   Basic understanding of programming concepts

### Installation

1. Clone this repository:

```bash
git clone <repository-url>
cd golang-playlist
```

2. Verify Go installation:

```bash
go version
```

## 📖 How to Run Examples

Each directory contains a self-contained Go example. To run any example:

```bash
# Navigate to the example directory
cd 1_hello_world

# Run the example
go run main.go
```

### Example Outputs

**Hello World:**

```bash
$ cd 1_hello_world && go run main.go
hello world
```

**Goroutines:**

```bash
$ cd 20_goroutines && go run goroutine.go
doing task 0
doing task 1
doing task 2
...
```

## 🏗️ Project Structure

```
golang-playlist/
├── 1_hello_world/          # Basic Go program
├── 2_simple_values/        # Data types
├── 3_variables/           # Variable declaration
├── 4_constants/           # Constants
├── 5_for/                 # Loops
├── 6_if_else/             # Conditionals
├── 7_switch/              # Switch statements
├── 8_arrays/              # Arrays
├── 9_slices/              # Slices
├── 10_maps/               # Maps
├── 11_range/              # Range keyword
├── 12_functions/          # Functions
├── 13_variadic_functions/ # Variadic functions
├── 14_closures/           # Closures
├── 15_pointers/           # Pointers
├── 16_structs/            # Structs
├── 17_interfaces/         # Interfaces
├── 18_enums/              # Enums
├── 19_generics/           # Generics
├── 20_goroutines/         # Concurrency
├── 21_channels/           # Channels
├── 22_mutex/              # Mutex
├── 23_files/              # File operations
└── 24_packages/           # Packages and modules
    ├── auth/              # Authentication package
    ├── user/              # User package
    ├── go.mod             # Module definition
    └── main.go            # Main application
```

## 🔧 Dependencies

Most examples use only the Go standard library. The packages example includes external dependencies:

-   `github.com/fatih/color` - For colored terminal output

To install dependencies for the packages example:

```bash
cd 24_packages
go mod tidy
```

## 📝 Learning Path

For beginners, follow this recommended order:

1. **Start with basics**: `1_hello_world` → `2_simple_values` → `3_variables` → `4_constants`
2. **Control flow**: `5_for` → `6_if_else` → `7_switch`
3. **Data structures**: `8_arrays` → `9_slices` → `10_maps` → `16_structs`
4. **Functions**: `12_functions` → `13_variadic_functions` → `14_closures` → `15_pointers`
5. **Advanced concepts**: `17_interfaces` → `18_enums` → `19_generics`
6. **Concurrency**: `20_goroutines` → `21_channels` → `22_mutex`
7. **Practical skills**: `23_files` → `24_packages`

## 🎯 Key Learning Objectives

-   **Syntax and Structure**: Understand Go's clean and simple syntax
-   **Data Types**: Master Go's type system and data structures
-   **Functions**: Learn function declaration, closures, and pointers
-   **Concurrency**: Understand goroutines, channels, and synchronization
-   **Packages**: Organize code into reusable modules
-   **Best Practices**: Follow Go idioms and conventions

## 🤝 Contributing

Feel free to contribute by:

-   Adding new examples
-   Improving existing code
-   Fixing bugs
-   Enhancing documentation

## 📄 License

This project is open source and available under the [MIT License](LICENSE).

## 🔗 Resources

-   [Go Official Documentation](https://golang.org/doc/)
-   [Go by Example](https://gobyexample.com/)
-   [Effective Go](https://golang.org/doc/effective_go.html)
-   [Go Tour](https://tour.golang.org/)

---

Happy coding! 🚀
