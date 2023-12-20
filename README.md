# Directory Backend Project

This is a backend development mini-project made in `go` (golang) over the framework `gofiber` and the orm `gorm`, and `sqlite` for my GitHub portfolio.

## Features

- **User Management**: CRUD operations for managing user data.
- **SQLite Database**: Utilizes SQLite as the database for storing user information.
- **RESTful API**: Provides RESTful endpoints for user-related operations.
- **Fiber Framework**: Built using the Fiber framework in Go for handling HTTP requests.
- **Testing with REST Client**: Includes a `tests.rest` file for testing the API endpoints using Visual Studio Code's `REST client` extension.

## Usage

1. **Running the Application**:
    - Ensure Go is installed on your system.
    - Install necessary dependencies using `go mod tidy`.
    - Run the application using `go run main.go`.

2. **Testing Endpoints**:
    - Use the `tests.rest` file with Visual Studio Code's REST client extension.
    - Each request in `tests.rest` file targets specific endpoints to test CRUD operations.

## How to Run

1. Clone the Repository:
    ```bash
    git clone https://github.com/clrql/backend-directory.git
    ```

2. Install Dependencies:
    ```bash
    cd backend-directory
    go mod tidy
    ```

3. Start the Application:
    ```bash
    go run main.go
    ```

4. Test Endpoints:
    - Open the `tests.rest` file in Visual Studio Code.
    - Use the REST client extension to execute HTTP requests against the defined endpoints.

## API Endpoints

- **GET `/users`**: Retrieve all users.
- **POST `/users`**: Create a new user.
- **GET `/users/:id`**: Retrieve a specific user by ID.
- **PUT `/users/:id`**: Update a specific user by ID.
- **DELETE `/users/:id`**: Delete a specific user by ID.