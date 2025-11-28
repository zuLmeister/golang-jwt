---

## Setup & Installation JWT Golang GORM

### Clone Project
```bash
git clone https://github.com/zuLmeister/golang-jwt.git

cd golang-jwt
```

### Install Dependency
```bash 
go mod tidy
```

### Make mysql Database
```bash
CREATE DATABASE myappdb;
```

### Generate JWT Secret Key 
```bash
Windows (PowerShell):

[Convert]::ToBase64String((1..64 | ForEach-Object {Get-Random -Maximum 256}))


Linux / macOS:

openssl rand -base64 64
```

### Buat file .env
```bash
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASS=yourpassword
DB_NAME=myappdb


ACCESS_SECRET=fillwithyourgeneratejwt
REFRESH_SECRET=fillwithyourgeneratejwt
```

### Running the application

```bash
go run main.go
```

### There will be display like this
```
Server default:
http://localhost:8080
```


### API Endpoint
```bash
Method	Endpoint	
POST	/register	
POST	/login	
POST	/refresh	
GET	/profile	
```

### Collection JSON
Better yet you can import the collection json with the name file
```bash
auth-collection.json
```


Author

Developed by Zulmeister
GitHub: https://github.com/zuLmeister
