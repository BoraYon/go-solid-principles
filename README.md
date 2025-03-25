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

# 🧱 OCP – Open/Closed Principle

This example demonstrates the **Open/Closed Principle (OCP)** using Go.

> **"Software entities should be open for extension, but closed for modification."**  
> — Bertrand Meyer

---

## 🧠 Concept

The idea is that we should be able to **add new functionality** to our system **without changing existing code**.  
This leads to more maintainable and stable codebases, especially in growing systems.

In Go, OCP is typically achieved through:
- Interfaces
- Composition
- Polymorphic behavior via struct-method patterns

---

## 💳 Example – Payment Processing System

We have a payment processor that can handle different payment methods, such as:

- **Credit Card**
- **Bank Transfer**
- **Crypto Payment** (e.g., Bitcoin)

Each payment method implements a common `PaymentMethod` interface with a `Pay()` function.

> New payment types can be added without modifying existing logic.  
> This is **true OCP** in action.

---

### 📦 Components

| File                | Responsibility                              |
|---------------------|----------------------------------------------|
| `payment.go`        | Declares the `PaymentMethod` interface       |
| `credit_card.go`    | Implements payment via credit card           |
| `bank_transfer.go`  | Implements payment via bank transfer         |
| `crypto_payment.go` | Implements payment via cryptocurrency        |
| `main.go`           | Executes payments through a generic function |

---

### 🖥️ Sample Output

### 🔍 OCP in Practice

In this project:

- We **did not modify** any existing files to add `CryptoPayment`.
- We simply created a new file `crypto_payment.go` that implements the same interface.
- The `processPayment()` function remained untouched.
- **This is exactly what the Open/Closed Principle promotes.**

---

## 🔹 LSP – Liskov Substitution Principle

> ✍️ Coming soon...

---

## 🔹 ISP – Interface Segregation Principle

> ✍️ Coming soon...

---

## 🔹 DIP – Dependency Inversion Principle

> ✍️ Coming soon...


