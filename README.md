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
│           main.go                     # for run start server
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
│   ├───repository                      # Handles database operations
│   │       token_repository.go
│   │       user_repository.go
│   ├───server                          # starting server
│   │       grpc.go
│   └───service                         # Business logic service
│           auth_service.go
│           user_service.go
├───proto                               # gRPC API definition
│       auth.pb.go
│       auth.proto
│       auth_grpc.pb.go
│       user.pb.go
│       user.proto
│       user_grpc.pb.go
└───utils                               # Helper functions
│       hash.go
│       jwt.go
│       jwt_interceptor.go
│       validate.go
│   .env
│   .gitignore
│   go.mod
│   go.sum
│   README.md
```
## Run Project
### 1. Create .env and Setting (Example)
```env
MONGO_URI=mongodb://localhost:27017
MONGO_DB=auth_db
GRPC_PORT=:50051
JWT_SECRET=your-super-secret
```

### 2. Install Dependency 
```bash
go mod tidy
```

### 3. Run the gRPC Server
```bash
go run cd cmd/server/main.go
```

## ทดสอบด้วย Postman
### 1. ติดตั้ง Postman (Desktop App) 

### 2. กด New และเลือก gRPC

### 3. ใส่ `url` ตาม port ที่ตั้งไว้ใน .env (Ex. localhost:50051)

### 4. ช่อง `select method` เลือก import a .proto file แล้วเลือกไฟล์นามสกุล .proto (โฟลเดอร์ proto) โดยจะมี 
#### 4.1. auth.proto `(Register, Login, Logout)`
#### 4.1. user.proto `(GetProfile, updateProfile, deleteProfile, ListUsers)`

### 5. import แล้ว select method จากไฟล์ที่นำเข้ามาได้ แล้วมาที่ `Message` กดที่ `Use Example Message` เพื่อกรอกข้อมูลตามที่เตรียมไว้

### 6. กดปุ่ม invoke เพื่อทดสอบ API
