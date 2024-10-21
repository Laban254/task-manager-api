## Project Setup 🛠️

To set up the Task Management API project, follow these steps:

### Prerequisites 📋

-   **Go**: Ensure you have Go installed (version 1.20 or higher is recommended). You can download it from [the official Go website](https://golang.org/dl/).
-   **PostgreSQL**: Install PostgreSQL for database management. Make sure to set up a database for the application.
-   **Git**: Install Git for version control.

### Steps to Set Up the Project 🚀

1.  **Clone the Repository**: Open your terminal and run the following command to clone the repository:
    

    
    `git clone https://github.com/yourusername/task-management-api.git` 
    
2.  **Navigate to the Project Directory**: Change to the project directory:
    

    
    `cd task-management-api` 
    
3.  **Initialize Go Modules**: Initialize a new Go module (if not done already):
    

    `go mod init task_management_api` 
    
4.  **Install Gin Framework**: Install the Gin framework by running:

    `go get -u github.com/gin-gonic/gin` 
    
5.  **Install GORM and PostgreSQL Driver**: Run the following commands to install GORM and the PostgreSQL driver:

    `go get -u gorm.io/gorm
    go get -u gorm.io/driver/postgres` 
    
6.  **Configure Environment Variables**: Create a `.env` file in the root directory to store your environment variables. Here's an example configuration:
    

    `DB_USER=your_db_user
    DB_PASSWORD=your_db_password
    DB_HOST=localhost
    DB_PORT=5432
    DB_NAME=your_db_name
    JWT_SECRET=your_jwt_secret` 
    
7.  **Run Migrations**: Ensure your database is set up correctly and run any migrations if you have defined them in your application.
    
8.  **Start the Application**: You can start the application by running:
    

    `go run cmd/main.go` 
    
9.  **Access the API**: Once the server is running, you can access the API at `http://localhost:8080`.
    

### Testing the Setup ✅

After setting up, you can test the API using tools like Postman or cURL to ensure the endpoints are functioning correctly.