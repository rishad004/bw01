# **User Management and Method Handling Using Go**

## **Overview**
This project implements two Go microservices with gRPC communication:  
- **Api Gateway**: Handles all client requests, routing them to the appropriate microservices.  
- **Microservice 01**: Handles user management (CRUD operations) with PostgreSQL and Redis.  
- **Microservice 02**: Executes parallel and sequential methods with database interactions.

---

## **Features**
### **Microservice 01**:  
- Create, retrieve, update, and delete user records.  
- Caches user data in Redis for faster access.  

### **Microservice 02**:  
- Executes **Method 01** sequentially and **Method 02** in parallel.  
- Communicates with **Microservice 01** via gRPC to fetch user data.

---

## **Prerequisites**
- **Go** (1.19 or higher)  
- **Docker** and **Docker Compose**  
- **PostgreSQL**  
- **Redis**   

---

## **Quick Setup**

### **Step 1: Clone Repository**
```bash
git clone https://github.com/rishad004/bw01.git
cd micro-user-method
```

### **Step 2: Run with Docker Compose**
Use the provided `docker-compose.yml` to spin up all services with one command:  
```bash
docker-compose up --build
```

This will start the following services:  
- **Microservice 01**: `localhost:50051`  HTTP port
- **Microservice 02**: `localhost:50052` Grpc port 
- **PostgreSQL**: Running on `localhost:5432`  
- **Redis**: Running on `localhost:6379`

---

## **API Endpoints**

### **Api Gateway**: User Management API
The User Management provides the following REST endpoints:

| Method | Endpoint         | Description                   |
|--------|------------------|-------------------------------|
| POST   | `/user/create`   | Create a new user.            |
| GET    | `/user/fetch`    | Retrieve user by ID.          |
| PUT    | `/user/update`   | Update user details.          |
| DELETE | `/user/delete`   | Delete user by ID.            |

---

### **Api Gateway**: Method Execution API
The Method Execution supports the following operations:

| Method | Endpoint | Description                                                 |
|--------|----------|-------------------------------------------------------------|
| GET    |`/method`| Accepts a `method` (1 or 2) and `wait` (seconds).       |

#### **Behavior**:
- **Method 01**: Executes sequentially.  
- **Method 02**: Executes in parallel.

---

## **Testing the APIs**

1. **Create a User**  
   - Endpoint: `POST /user/create`  
   - Sample Request:  
     ```json
     {
       "name": "Jack",
       "email": "Jack@example.com",
     }
     ```

2. **Retrieve a User by ID**  
   - Endpoint: `GET /user/fetch`  
    - Sample Request:  
      ```json
      {
        "id": 1,
      }
      ```

3. **Update User Details**  
   - Endpoint: `PUT /user/update/:id`
   - Sample Request:
   - Endpoint: `PUT /user/update/1`
     ```json
     {
       "name": "Jane Doe",
       "email": "jane.doe@example.com",
     }
     ```

4. **Delete a User by ID**  
   - Endpoint: `DELETE /user/delete`
   - Sample Request:  
     ```json
     {
       "id": 1,
     }
     ```

5. **Execute a Method**  
   - Endpoint: `POST /method`  
   - Sample Request:  
     ```json
     {
       "method": 1,
       "wait": 10
     }
     ```
   - Methods:
     - **Method 01**: Executes tasks sequentially.  
     - **Method 02**: Executes tasks in parallel.  


---
