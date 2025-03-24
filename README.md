# go-solid-principles

This repository demonstrates each of the five SOLID principles using the Go programming language.  
Each principle is placed in its own folder (`SRP`, `OCP`, `LSP`, `ISP`, `DIP`) and contains a focused example.

---

## 📘 Index

1. [SRP – Single Responsibility Principle](#-srp--single-responsibility-principle)
2. [OCP – Open/Closed Principle](#-ocp--openclosed-principle)
3. [LSP – Liskov Substitution Principle](#-lsp--liskov-substitution-principle)
4. [ISP – Interface Segregation Principle](#-isp--interface-segregation-principle)
5. [DIP – Dependency Inversion Principle](#-dip--dependency-inversion-principle)

---

## 🔹 SRP – Single Responsibility Principle

This example demonstrates the **Single Responsibility Principle (SRP)** using Go.

Each component of the logging system is designed to perform a **single, clearly defined task**:

- `FileLogger`: Responsible for writing logs to a local file.
- `HTTPLogger`: Simulates sending log entries to a remote HTTP server.
- `LogManager`: Manages the coordination of logging operations without performing the logging itself.
- `LogEntry`: A simple struct that represents a log message with a timestamp.

The application uses **goroutines and channels** to simulate a concurrent logging system with a worker pool.  
Each worker processes log entries independently, sending them both to the file and to the simulated server.

### 🖨️ Sample Output
[Worker-1] Processing log...
SENT TO SERVER: [2025-03-24T10:15:00Z] Event #1 happened
[Worker-2] Processing log...
SENT TO SERVER: [2025-03-24T10:15:00Z] Event #2 happened
...
All loggers shut down gracefully.

          +------------+      +-----------------+
          |  main.go   | ---> |   LogManager    |
          +------------+      +--------+--------+
                                        |
        +-------------------------------+-------------------------------+
        |                                                               |
+---------------+                                         +----------------------+
|  FileLogger   |                                         |     HTTPLogger       |
+---------------+                                         +----------------------+
| Writes logs to file  📝 |                              | Sends logs to server 🌐 |
+------------------------+                              +------------------------+


---

## 🔹 OCP – Open/Closed Principle

> ✍️ Coming soon...

---

## 🔹 LSP – Liskov Substitution Principle

> ✍️ Coming soon...

---

## 🔹 ISP – Interface Segregation Principle

> ✍️ Coming soon...

---

## 🔹 DIP – Dependency Inversion Principle

> ✍️ Coming soon...


