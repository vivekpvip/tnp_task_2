# CERTIFICATE1.GO

# APPROACH TO THE SOLUTION :
MY main goal of this project is to create a system where certificates can be stored,updated, and sent to users via email.
The system allows you to manage certificates through an API and send them via email individually or in bulk

1. ** FIRST APPROACH ** : MY FIRST APPROACH IS TO SETUP DATABASE (POSTGRE) FOR CREATING CONTROLLER. SO I LEARNED ABOUT HOW TO SETUP postgreSQL DATABASE.
I INSTALL THE REQUIRED PACKAGES. The system connects to the database and automatically creates the necessary table to store certificates
After this i create a  controller , In the context of this code Controller are implemented as hendler functions like: GetCertificateByID, CreateCertificate, etc.
2. ** SECOND APPROACH ** : my second approach is to send these certificate via email and to create bulk messaging system, 
I Created a SendCertificate Function to send certificate to our client
The system can send certificates via email to specific users. We provide an email address, and the system sends the certificate to that email.
Additionally, we can send the same certificate to multiple people in one go (bulk email).
3. ** THIRD APPROACH ** : my third approach is to set middlewares 
I set a middleware for security purpose To make sure only authorized users can use the system, there's a simple security check. 
Each request needs to include a special token (like a password), which the system verifies before allowing access.

# PROBLEMS ENCOUNTERED DURING WORKING ON TASK :
I have faces many problems during working on task and solve them with the help of chatgpt,youtube and self efforts.
my first problem/confusion is that, Should i modify these features to my existing code or create a new go file including these features,
but i modified these features to my existing code by updating code

1. ** Database Connection Issue ** : Initially, I had trouble connecting to the PostgreSQL database using GORM. 
I had to ensure that the dsn (Data Source Name) was configured correctly with the right username, password, database name, and port. 
Additionally, I had to install and configure the required PostgreSQL driver.
2. ** Email Sending Issue ** : When attempting to send emails through Gmail's SMTP server, I encountered difficulties related to Gmail’s security settings. 
To resolve this, I had to enable "Less secure apps" or use an app-specific password in Gmail to allow the Go application to send emails successfully.
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





