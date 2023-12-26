# Project README

This README provides an overview of the structure, design decisions, and libraries used in the HTTP server project. The goal of this project is to create an HTTP server that facilitates communication with third-party services based on tasks submitted by clients.

## Project Structure

The project is organized into three main directories:

### 1. cmd

The `cmd` directory houses the main entry point of the application. The `main.go` file, located in this directory, serves as the starting point for launching the application.

### 2. internal

The `internal` directory contains all the essential logic required for the application to function. This includes controllers and services responsible for handling client requests and managing background task execution.

- **Controllers**: These handle incoming HTTP requests and interact with the services to process tasks.

- **Services**: These encapsulate the business logic of the application and manage the execution of tasks in the background.

### 3. pkg

The `pkg` directory houses utility files that are not directly related to the business logic but provide essential functions for the application.

## Design Decisions

The decision to split the project into these three directories is driven by the separation of concerns:

- **cmd**: Focuses on the application's entry point, making it easy to identify where the execution begins.

- **internal**: Contains the core business logic, separating it from auxiliary functionalities and providing a clean and modular structure.

- **pkg**: Holds utility functions that are not specific to the business logic, promoting reusability.

## Libraries Used

The project leverages the following libraries:

1. [Viper](https://github.com/spf13/viper): For configuration management.
2. [Zerolog](https://github.com/rs/zerolog): A fast and flexible logging library.
3. [Echo](https://echo.labstack.com): A web framework for Go.
4. [Validator](https://github.com/go-playground/validator): For request validation.

## Project Functionality

The HTTP server is designed to handle tasks submitted by clients in JSON format. The client initiates a task by sending an HTTP POST request to `/task`, describing the task details. The server responds with a generated unique ID, and the task begins execution in the background.

### API Endpoints

1. **Submit Task:**

    - **Request:**
      ```
      POST /task
      {
        "method": "GET",
        "url": "http://google.com",
        "headers": {
          "Authentication": "Basic bG9naW46cGFzc3dvcmQ=",
          // ...
        }
      }
      ```

    - **Response:**
      ```
      200 OK
      {
        "id": <generated unique id>
      }
      ```

2. **Check Task Status:**

    - **Request:**
      ```
      GET /task/<taskId>
      ```

    - **Response:**
      ```
      200 OK
      {
        "id": <unique id>,
        "status": "done/in_process/error/new",
        "httpStatusCode": <HTTP status of 3rd-party service response>,
        "headers": {
          // headers array from 3rd-party service response
        },
        "length": <content length of 3rd-party service response>
      }
      ```

This project is geared towards creating an efficient and scalable HTTP server for managing asynchronous tasks with third-party services.