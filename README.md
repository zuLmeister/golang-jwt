Features

- Register User
- Login with JWT Access & Refresh Token
- Refresh Token Rotation (DB Stored)
- Logout (invalidate refresh token)
- Protected Route using Middleware
- MySQL Database Integration
- Clean Code & Scalable Folder Structure
- Cross-platform support (Windows/Linux/macOS)



Setup Project
Clone Repository
git clone https://github.com/zuLmeister/golang-jwt.git
cd golang-jwt

Install Dependencies
go mod tidy

Setup .env

Buat file .env:

DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASS=yourpassword
DB_NAME=myappdb

# Generate keys (PowerShell):

# [Convert]::ToBase64String((1..64 | ForEach-Object {Get-Random -Maximum 256}))

ACCESS_SECRET=GENERATED_TOKEN_HERE
REFRESH_SECRET=GENERATED_TOKEN_HERE

Running the App
OS Command
Windows go run main.go
Linux / Mac go run main.go

App berjalan default:
http://localhost:8080

API Endpoints
Method Endpoint
POST /register
POST /login
POST /refresh
POST /logout
GET /profile
GET /users
Authorization Header

Gunakan access token:

Authorization: Bearer <access_token>

API Collection

Postman collection included in repo:
postman_collection.json

Security Notes

Passwords encrypted using bcrypt

Refresh tokens stored in DB & rotated every login

On logout → refresh token removed from DB

Contribution

Pull Request selalu welcome!
Open issue kalau menemukan bug atau ingin request fitur baru 🚀

License

MIT © 2024 — Your Name
