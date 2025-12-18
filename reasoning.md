# Implementation Approach & Key Decisions

## Overview

For this assignment, I built a high-performance RESTful API using Go. My primary goal was to create a robust, type-safe, and maintainable backend that efficiently manages user data while adhering to modern Go best practices.

## Technology Stack & Rationale

### 1. Framework: GoFiber

I chose **Fiber** as my web framework because of its impressive performance benchmarks and its similarity to Express.js. Its zero-allocation philosophy and ease of use allowed me to set up routes and middleware quickly without sacrificing speed.

### 2. Database Access: SQLC with PostgreSQL

Instead of using a heavy ORM like GORM, I opted for **SQLC**.

- **Why?** I wanted the performance of raw SQL with the safety of compiled code.
- **How it works:** I write standard SQL queries in `db/queries/`, and `sqlc` generates type-safe Go code for me. This eliminates runtime SQL errors and ensures my database interactions are always in sync with my schema.

### 3. Logging: Uber Zap

I implemented **Zap** for structured logging. Standard library logging is insufficient for production debugging. Zap provides high-performance, structured (JSON) logs that are easy to parse and filter, which is crucial for tracking request flows and errors.

### 4. Validation: go-playground/validator

I used the `validator` library to enforce data integrity at the entry point. By using struct tags (e.g., `validate:"required,min=2"`), I keep my validation logic declarative and close to the data definition, preventing bad data from ever reaching my business logic.

## Architecture & Design

### Project Structure

I followed the **Standard Go Project Layout**:

- **`cmd/server`**: The entry point of the application.
- **`internal/`**: Contains the private application code (handlers, models, middleware) that shouldn't be imported by external projects.
- **`db/`**: Houses all database-related code, including migrations and the generated SQLC code.

### Error Handling

I implemented a consistent error handling strategy. Handlers return specific HTTP status codes (400 for bad input, 500 for server errors) and structured JSON error messages. This ensures the API client always knows exactly what went wrong.

### Dynamic Age Calculation

One of the specific requirements was to calculate age dynamically. I implemented a helper function in the `service` package that calculates the age based on the Date of Birth (DOB) and the current time, ensuring the age is always accurate when requested, rather than storing a static value that would become stale.

## Conclusion

This approach balances performance with developer productivity. By leveraging code generation (SQLC) and robust libraries (Fiber, Zap, Validator), I was able to build a reliable API that is easy to extend and maintain.
