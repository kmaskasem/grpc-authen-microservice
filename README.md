# gRPC Authentication Microservice

Authentication Microservice with gRPC + JWT + MongoDB

## Features

- Register
- Login (use JWT to generate token)
- Logout
- Get List Users
- Get User Profile
- Update User Profile
- Delete User Profile

## Project Structure
```text
project
├───cmd
│   ├───client                          # for client testing
│   └───server
│           main.go                     # for start server
├───config
│       config.go                       # loads environment variables from .env
├───internal
│   ├───database
│   │       mongodb.go                  # connecting MongoDB
│   ├───handler
│   │   └───grpc                        # gRPC implementation service
│   │           auth_handler.go
│   │           user_handler.go
│   ├───model                           # Data model
│   │       token.go
│   │       user.go
│   ├───repository                      # Handles  database operations
│   │       token_repository.go
│   │       user_repository.go
│   └───service                         # Business logic service
│           auth_service.go
│           user_service.go
├───proto                               # gRPC API definition
│       auth.pb.go
│       auth.proto
│       auth_grpc.pb.go
│       user.proto
└───utils                               # Helper functions
│       hash.go
│       jwt.go
│       validate.go
│   .env
│   .gitignore
│   go.mod
│   go.sum
```
## Run Project
### 1. Create .env and Setting (Example)
```env
MONGO_URI=mongodb://localhost:27017
MONGO_DB=auth_db
JWT_SECRET=your-super-secret
```

### 2. Run the gRPC Server
```bash
go run cd cmd/server/main.go
```

## Testing with Postman

