# Capstone Project API

## Overview
This project is a backend API for a capstone project. It is built using Go and the Echo framework, and it connects to a MongoDB database.

## Features
- User authentication and authorization
- Thread management
- Category management
- Voting system
- Commenting system
- Reporting system
- Moderator management

## Requirements
- Go 1.16 or higher
- MongoDB
- Viper for configuration management

## Installation
1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/capstone-project-API.git
   ```
2. Install the dependencies:
   ```sh
    go mod download
    ```
3. Create a `.env` file in the root directory and add the following environment variables:
4. Run the application:
   ```sh
   go run main.go
   ```
## Configuration
Create a config.json file in the app directory with the following structure:
```json
{
  "database": {
    "username": "your_db_username",
    "password": "your_db_password",
    "host": "your_db_host",
    "port": "your_db_port",
    "name": "your_db_name"
  },
  "jwt": {
    "secret": "your_jwt_secret",
    "expired": 3600
  },
  "server": {
    "address": ":9090"
  },
  "context": {
    "timeout": 30
  }
}
```
