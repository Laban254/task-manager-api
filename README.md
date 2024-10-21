# Task Management API 🚀

## Overview 🌟

The Task Management API is a RESTful API built using the Gin framework in Go. It provides a robust backend for managing tasks and projects, featuring user authentication, role-based access control, and integration with Google OAuth for enhanced security and user experience.

Getting Started 🛠️
Refer to the  [Project Setup](./docs/project_setup.md) section in the original documentation for instructions on setting up the environment and dependencies.

## Features ⚙️

-   **User Authentication**:
    -   Supports JWT-based authentication for secure user sessions. 
    -   Implements OAuth 2.0 for Google authentication. 
-   **Role-Based Access Control (RBAC)**:
    -   Differentiate user roles (admin and regular users) with specific permissions for managing projects and tasks. 👥
-   **CRUD Operations**:
    -   Full support for creating, reading, updating, and deleting tasks and projects. 
    -   API endpoints include:
        -   `/login` - User login and JWT issuance. 
        -   `/register` - New user registration. 
        -   `/projects` - Manage projects (create, read, update, delete). 
        -   `/tasks` - Manage tasks (create, read, update, delete). 
-   **Input Validation and Error Handling**:
    -   Ensures incoming data is validated and errors are handled gracefully. 
-   **Middleware Support**:
    -   Integrated middleware for logging, CORS handling, and request parsing. 
    -   Custom middleware for JWT validation and rate-limiting. 
-   **Secure Communication**:
    -   All endpoints are protected and require authentication. 

## Directory Structure 📁

bash

Copy code

`/cmd            - Entry points for the application (main.go).
/pkg            - Reusable components (authentication, routes, models).
/config         - Configuration files (environment variables).
/internal       - Application-specific logic (services, handlers).
/scripts        - Deployment scripts (Docker, Kubernetes).` 

## Technologies Used 💻

-   **Go**: Programming language for building the API.
-   **Gin**: Web framework for building the RESTful API.
-   **GORM**: ORM for interacting with the database.
-   **PostgreSQL**: Database for storing user and project data.
-   **JWT**: For user authentication and session management.

