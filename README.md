# Go For Professionals - Workout Tracker API

This project is a backend API for a workout tracking application, built with Go. It allows users to register, authenticate, and manage their workouts and workout entries.

## Description

The "Go For Professionals" project is a RESTful API designed to serve as the backend for a fitness application. It provides functionalities for user management (registration, login via tokens) and workout management (creating, retrieving, updating, and deleting workouts and their specific entries). The application uses a PostgreSQL database to store user and workout data.

## Learning Project Disclaimer

> [!CAUTION]
> **Please Note:**
> This is a personal learning project. Its primary goal is to explore and understand various concepts in Go backend development, API design, and database interactions.

For the sake of simplicity and to focus on core learning objectives, some production-level best practices have been intentionally omitted or simplified. For example:

- **Configuration Management**: Some configurations, such as database connection strings (e.g., in `internal/store/database.go` and `internal/store/workout_store_test.go`), are hardcoded. In a production environment, these would typically be managed through environment variables or configuration files.
- **Error Handling & Logging**: While basic logging and error handling are implemented, they might not be as comprehensive or robust as required for a production system.
- **Security**: Security considerations are addressed at a basic level (e.g., password hashing, token authentication). However, a production application would require more in-depth security measures, audits, and potentially more complex authentication/authorization mechanisms.

This approach allows the project to remain focused on the specific learning goals without the added complexity of full production readiness.

## Features

- **User Authentication**: Secure user registration and token-based authentication.
- **Workout Management**:
  - Create new workouts with details like title, description, duration, and calories burned.
  - Add detailed entries to workouts, including exercise name, sets, reps or duration, weight, and notes.
  - View existing workouts by ID.
  - Update existing workouts.
  - Delete workouts.
- **Database Migrations**: Uses `goose` for managing database schema migrations.
- **Dockerized Environment**: Includes `docker-compose.yml` for easy setup of PostgreSQL databases for development and testing.
- **Structured Logging**: Implements logging for application events and errors.
- **Configuration**: Application configuration (e.g., port) can be set via command-line flags.
- **Testing**: Includes unit tests for store functionalities (e.g., `workout_store_test.go`).

## Technologies Used

- **Language**: Go
- **Web Framework**: Chi (used for routing)
- **Database**: PostgreSQL
- **Database Driver**: `pgx`
- **Migrations**: `pressly/goose`
- **Password Hashing**: `bcrypt`
- **Containerization**: Docker (via `docker-compose`)
- **Testing**: Testify (for assertions and requirements in tests)
