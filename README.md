[No content]
# Go Server

A RESTful server built with Go for managing users and chirps (short messages), featuring authentication, metrics, and readiness endpoints.

## Features
- User registration and login
- JWT-based authentication and refresh tokens
- CRUD operations for chirps
- Metrics and readiness endpoints
- SQL database integration
- Secure password hashing

## Folder Structure
```
Go-Server/
├── assets/                # Static assets (e.g., logo)
├── internal/
│   ├── auth/              # Authentication logic (JWT, password hashing)
│   ├── database/          # Database models and queries
├── sql/
│   ├── queries/           # SQL query files
│   ├── schema/            # Database schema migrations
├── handler_*.go           # HTTP handlers for API endpoints
├── main.go                # Server entry point
├── metrics.go             # Metrics endpoint
├── readiness.go           # Readiness endpoint
├── reset.go               # Reset endpoint
├── index.html             # Main HTML page
├── go.mod, go.sum         # Go modules
```

## Setup Instructions
1. **Clone the repository:**
	```sh
	git clone https://github.com/Mayank3299/Go-Server.git
	cd Go-Server
	```
2. **Install dependencies:**
	```sh
	go mod tidy
	```
3. **Configure the database:**
	 - Update connection details in `internal/database/db.go` as needed.
	 - Run migrations in `sql/schema/` to set up the database.
	 - **Run Goose for migrations:**
		 - Install Goose if you don't have it:
			 ```sh
			 go install github.com/pressly/goose/v3/cmd/goose@latest
			 ```
		 - Run migrations:
			 ```sh
			 goose -dir sql/schema postgres <your_postgres_connection_string> up
			 ```

4. **Create a `.env` file:**
	The server expects environment variables for database connection, platform, secrets, and API keys. Create a `.env` file in the project root with the following format (replace values with your own):
	```env
	DB_URL="postgres://<username>:<password>@localhost:5432/<database>?sslmode=disable"
	PLATFORM="dev"
	SECRET="<your_secret_key>"
	POLKA_KEY="<your_polka_api_key>"
	```

5. **Build and run the server:**
	```sh
	go build -o out && ./out
	```

## Usage
The server exposes RESTful endpoints for user and chirp management. Use tools like `curl` or Postman to interact with the API.

## API Endpoints
### User
- `POST /users` - Register a new user
- `POST /login` - Login and receive JWT
- `POST /refresh` - Refresh JWT
- `POST /revoke` - Revoke refresh token
- `PUT /users/{id}` - Update user

### Chirps
- `POST /chirps` - Create a chirp
- `GET /chirps` - List chirps
- `GET /chirps/{chirpID}` - Get chirp by ID
- `DELETE /chirps/{chirpID}` - Delete chirp

### Other
- `GET /metrics` - Metrics endpoint
- `GET /readiness` - Readiness endpoint
- `POST /reset` - Reset server state

## Author

Mayank Agarwal

## Contact
For questions or support, reach out via email: [mayank.agarwal0903@gmail.com](mailto:mayank.agarwal0903@gmail.com)
