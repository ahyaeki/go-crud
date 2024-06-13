# CRUD Project

Application to test the flow of CRUD logic.

## Features
- Create new records
- Read existing records
- Update records
- Delete records
- Login and logout with authorization key

## Table of Contents
- [Installation](#installation)
- [Database Setup](#database-setup)
- [Usage](#usage)
- [API Documentation](#api-documentation)

## Installation

Download the ZIP file and extract it in a directory.

### Prerequisites
- [Go](https://golang.org/)
- [PostgreSQL](https://www.postgresql.org/)
- [Postman](https://www.postman.com/)

### Steps

1. **Clone the repository**
    ```sh
    git clone https://github.com/ahyaeki/go-crud.git
    cd go-crud
    ```

2. **Install dependencies**
    ```sh
    go mod tidy
    ```

3. **Set up environment variables**

    Create a `.env` file in the root directory and add the necessary environment variables:
    ```env
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=postgres
    DB_PASSWORD=Ahyaeki12
    DB_NAME=mydb
    ```

4. **Run the application**
    ```sh
    go run main.go
    ```

5. **Access the application**

    The server will start on `http://localhost:8080`.

## Database Setup

### Create Database and Tables

1. **Create the database**
    ```sql
    CREATE DATABASE mydb;
    ```

2. **Create tables**

    Connect to your database and run the following SQL commands to create the necessary tables:
    ```sql
      CREATE TABLE IF NOT EXISTS public.items
      (
        id integer NOT NULL DEFAULT nextval('items_id_seq'::regclass),
        name character varying(100) COLLATE pg_catalog."default",
        description text COLLATE pg_catalog."default",
        CONSTRAINT items_pkey PRIMARY KEY (id)
      )

      TABLESPACE pg_default;

      ALTER TABLE IF EXISTS public.items
        OWNER to postgres;

      CREATE TABLE IF NOT EXISTS public.users
      (
        id integer NOT NULL DEFAULT nextval('items_id_seq'::regclass),
        username character varying(100) COLLATE pg_catalog."default",
        password character varying(100) COLLATE pg_catalog."default",
        session_token character varying(255) COLLATE pg_catalog."default",
        CONSTRAINT users_pkey PRIMARY KEY (id)
      )

      TABLESPACE pg_default;

      ALTER TABLE IF EXISTS public.users
        OWNER to postgres;
    ```

3. **Insert initial data**
    ```sql
    INSERT INTO users (username, password) VALUES ('Ahya', '123');
    INSERT INTO users (username, password) VALUES ('Eki', '456');
    ```

4. **Run database migrations** (if using a migration tool)
    ```sh
    go run migrate.go
    ```

## Usage

### Running the application
- To start the server: `go run main.go`
- Start operating in Postman

### Using Postman

1. **Import the Postman collection** provided in the repository.
2. **Send requests to the API**:
    - **Create**: Send a POST request to `/login` or `/logout` and `/items`
    - **Read**: Send a GET request to `/items` or `/items/{id}` 
    - **Update**: Send a PUT request to `/items` or `/items/{id}`
    - **Delete**: Send a DELETE request to `/items/{id}`

## API Documentation

### Endpoints

- **GET /items**: Retrieve a list of items
- **POST /login**: Logging in and receive a session_key 
- **POST /logout**: Logging out and delete a session_key
- **POST /items**: Create a new item
- **GET /items/{id}**: Retrieve an item by ID
- **PUT /items/{id}**: Update an item by ID
- **DELETE /items/{id}**: Delete an item by ID

### Example Requests

(In order to create, read, update, and delete items, it is necessary to login first by using the method of POST in `/login` with existing user's information. After logged in, session_token will be generated, which will be manually entered inside the header for Authorization as its key name and session_token's value as its key value.)

#### Create an item
- **URL**: `/items`
- **Method**: POST
- **Body**: 
    ```json
    {
        "name": "Milk",
        "description": "Description of milk"
    }
    ```

#### Get All Items
- **URL**: `/items`
- **Method**: GET

#### Get an item by ID
- **URL**: `/items/1`
- **Method**: GET

#### Update an item
- **URL**: `/users/1`
- **Method**: PUT
- **Body**: 
    ```json
    {
        "name": "Cocoa",
        "email": "Description of Cocoa"
    }
    ```

#### Delete an item
- **URL**: `/items/1`
- **Method**: DELETE

### Logging in
- **URL**: `/login`
- **Method**: POST
- **Body**:
    ```json
    {
        "username": "Ahya",
        "password": "123"
    }
    ```
- After entering the method POST, session_token will be generated. Put the session_token inside Authorization's value in order to access the /items, otherwise it will return Forbidden/Unauthorized.

### Logging out
- **URL**: `/logout`
- **Method**: POST
- **Body**:
    ```json
    {
        "username": "Ahya",
        "password": "123"
    }
    ```
- Make sure to have the same session_token that was generated during login process, otherwise it will return Forbidden/Unathorized.

## Contributing

Contributions are welcome! Please follow these steps to contribute:
1. Fork the repository
2. Create a new branch (`git checkout -b feature/your-feature-name`)
3. Make your changes
4. Commit your changes (`git commit -m 'Add some feature'`)
5. Push to the branch (`git push origin feature/your-feature-name`)
6. Open a pull request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
