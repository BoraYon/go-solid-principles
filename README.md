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

🧱 LSP – Liskov Substitution Principle

This example demonstrates the Liskov Substitution Principle (LSP) using Go.

"Objects of a superclass should be replaceable with objects of its subclasses without affecting the correctness of the program."— Barbara Liskov

🧠 Concept

LSP ensures that a derived type (or a struct implementing an interface in Go) can replace its base type without altering the correct behavior of the system.

In Go, we apply LSP using:

Interfaces: Ensuring consistent behavior across different types.

Struct Implementations: Each struct adheres to the defined contract.

Avoiding Misuse: Structs should not implement interfaces if they cannot fulfill their responsibilities.

💳 Example – Payment Processing with Fraud Detection

In this example, we have a payment processing system where multiple payment methods can be used:

Credit Card Payment

Crypto Payment

Additionally, we introduce a Fraud Detection Service to detect fraudulent credit card transactions.

The correct use of LSP ensures that:

Payment processors are interchangeable.

Fraud detection is handled by a dedicated service without misusing the interface.

📦 Components

File

Responsibility

payment.go

Defines the PaymentProcessor interface

credit_card.go

Implements payments using credit card

crypto_payment.go

Implements payments using cryptocurrency

fraud_detection.go

Contains the fraud detection logic (LSP compliant)

main.go

Simulates payment processing and fraud detection

🚀 Example Code Execution

When the application runs, it processes payments and checks for fraud. Fraudulent transactions are blocked.

Sample Output

Processing payments...
Paid 150.00 using Credit Card ending with 3456
Paid 0.015 BTC to wallet 1BitcoinWalletExample123
Performing fraud check...
Transaction blocked: Fraud detected on card ending with 3456

✅ LSP Compliance Explanation

CreditCardPayment and CryptoPayment both implement the PaymentProcessor interface.

The processPayment function works with any valid payment processor without modification.

The FraudDetectionService does not implement PaymentProcessor, adhering to LSP by not misusing the interface.

Changes in fraud detection or new payment methods will not affect the existing code.

By applying LSP, the system remains scalable, maintainable, and error-free.

🔎 Conclusion

This project effectively demonstrates how to apply the Liskov Substitution Principle in Go by designing robust and flexible software systems. Each component serves a clear purpose, and future expansions can be added without disrupting existing functionality.



---

## 🔹 ISP – Interface Segregation Principle

> ✍️ Coming soon...

---

## 🔹 DIP – Dependency Inversion Principle

> ✍️ Coming soon...


