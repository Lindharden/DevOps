# Lecture 2: Packaging applications, Containerization with Docker
February 07 and 10, 2023

## Step 1: Refactor ITU-MiniTwit to another language and technology of your choice

First of all we ran our test suite to see whether the tests pass or not. Here we have some tests that fail. To fix all of these tests, we had to change the string encoding. 

For our refactoring we decided to use the language GO. We choose this language as it's much faster than Python, and because we wanted to learn it. We installed an GO SQL package on the system, and imported it into our application. This allows us to query the database using `sql.open()`. Then we wanted to update our test suite to support our GO application. We had to update the flask HTML templates to work with GO. Here we had to change the expression declarations from using the following syntax: `{% expression %}` to `{{expression}}`. We then changed the templates to use some new controllers and routes we created for handling different tasks. We created controllers, routes and middleware as described in this blog by Jes Fink-Jensen: <https://betterprogramming.pub/how-to-create-a-simple-web-login-using-gin-for-golang-9ac46a5b0f89>.


## Step 2: Containerize ITU-MiniTwit with Docker

## Step 3: Describe Distributed Workflow
