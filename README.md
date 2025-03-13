# Certificate Management API
This project implements a Certificate Management System using Go and PostgreSQL. It provides a RESTful API to create, retrieve, update, and manage certificates. Additionally, it includes the functionality to send certificates via email (both individual and bulk email).

# Approach to the Solution
The solution follows a structured approach to handle various operations related to certificates. The main components are as follows:

 # Database Setup:
A PostgreSQL database is used to store certificates. The gorm package is used as an ORM to interact with the database.
A Certificate model is defined with fields such as ID, Name, Content, Owner, and Date.

 # API Endpoints:

• GET /certificates/{id}: Retrieves a certificate by its ID.
• POST /certificates: Creates a new certificate.
• GET /certificates: Retrieves all certificates.
• PUT /certificates/{id}: Updates a certificate by ID.
• POST /send/{id}: Sends a specific certificate to an email.
• POST /send_bulk: Sends bulk emails with a custom message/content.

 # Authentication:
A basic authentication middleware (AuthMiddleware) is implemented using Bearer token to ensure that only authorized users can access the endpoints.

 # Email Sending:
The gomail package is used to send certificates via email.
Individual certificate emails are sent using the SendCertificate handler, and bulk emails are handled by the SendBulkEmail handler.

# PROBLEMS ENCOUNTERED DURING WORKING ON TASK :
I have faces many problems during working on task and solve them with the help of chatgpt,youtube and self efforts.
my first problem/confusion is that, Should i modify these features to my existing code or create a new go file including these features,
but i modified these features to my existing code by updating code

1. ** Database Connection Issue ** : Initially, I had trouble connecting to the PostgreSQL database using GORM. 
[solution] : I had to ensure that the dsn (Data Source Name) was configured correctly with the right username, password, database name, and port. 
Additionally, I had to install and configure the required PostgreSQL driver.
2. ** Email Sending Issue ** : When attempting to send emails through Gmail's SMTP server, I encountered difficulties related to Gmail’s security settings. 
[solution] : I had to enable "Less secure apps" or use an app-specific password in Gmail to allow the Go application to send emails successfully.
3. ** Error Handling in Database Operations** : Handling various database errors, such as "record not found" or "database connection issues", required a thorough understanding of GORM’s error handling mechanisms.
[solution] : Ensuring the API returns clear and informative error messages to the client was essential for debugging and user experience.
4. ** Additionaly most of the time i forget to update the modifications in handler's function and router too.

## What Makes This Solution Unique
This solution is unique because:
1. **[End to End Certificate Management]**: [The system offers a comprehensive solution for managing certificates, from creation and storage in a database to sending them via email]

2. **[Bulk Email Sending]**: [The inclusion of bulk email functionality sets this solution apart by making it easy to send certificates to multiple users at once]

3. **[Security with Authorization Middleware]**: [he integration of an Authorization Middleware ensures that only authorized users can access and interact with the system]
# Additionaly :
1. Integration with GORM for Database Handling
2. The system uses Gmail’s email service to send emails. When you want to send a certificate, it connects to Gmail’s SMTP server and sends the email to the specified address.

# CONCLUSION :
Overall, the project was an interesting challenge. I learned how to add security to our api and controller and most of new things.


# DIRECTORY STRUCTURE 
/certificate-app
  ├── go.mod
  ├── cmd/
  │   └── main.go
  ├── handlers/
  │   ├── certificate.go
  │   └── email.go
  |   |__ auth.go
  ├── models/
  │   └── certificate.go
  ├── utils/
  │   └── db.go
  





