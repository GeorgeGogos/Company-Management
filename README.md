# Companies Microservice

This is a Golang-based microservice for managing companies. It supports basic CRUD (Create, Read, Update, Delete) operations on company data and includes user authentication using JWT (JSON Web Tokens). Only authenticated users can create, update, or delete companies, while any user can retrieve company details.

## Features

- **User Registration**: Users can register with an email and password.
- **User Authentication**: Users can log in and receive a JWT token to authenticate further requests.
- **CRUD Operations**: Users can create, update, delete, and retrieve company details.
- **Password Security**: Passwords are securely hashed using `bcrypt`.
- **JWT Authentication**: Protects routes, ensuring only authenticated users can access certain endpoints.
- **PostgreSQL Integration**: Uses PostgreSQL as the data storage layer.
- **Dockerized Setup**: Easily deployable using Docker and Docker Compose.

## Table of Contents

- [Technologies](#technologies)
- [Project Structure](#project-structure)
- [Getting Started](#getting-started)

## Technologies

- **Go (Golang)**: Backend service language.
- **PostgreSQL**: Relational database to store company and user information.
- **JWT**: Token-based authentication system.
- **Docker**: Containerization platform to run the application and its dependencies.
- **bcrypt**: Library for hashing passwords securely.

## Project Structure

The project is structured to separate different concerns such as controllers, models, services, and repositories.


## Getting Started

### Prerequisites

- **Go** (1.16 or higher) installed on your local machine.
- **Docker** installed for containerization.
- **Postman** or `curl` for testing API requests.

### Setting Up the Project

1. **Clone the repository**:

   ```bash
   git clone https://github.com/GeorgeGogos/Company-Management.git
   cd Company-Management

2. **Run Service**:

    ```bash
    docker-compose up --build -d 
