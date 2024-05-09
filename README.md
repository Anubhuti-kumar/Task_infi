To run this project please follow following steps

step 1. initialise the project ,run "go mod init infilon_test"
step 2. run, "go mod tidy"
step 3. go run .

Note please add "username:password" for mysql into config.go file.

// To get the output for GET-POST follow below steps

GET Request:
Open Postman.
Set the request type to GET.
Enter the URL of your endpoint (e.g., http://localhost:8080/person/1/info).
Click on the "Send" button.
You should receive a response containing the person's information.

POST Request:
Set the request type to POST.
Enter the URL of your endpoint (e.g., http://localhost:8080/person/create).
Go to the "Body" tab.
Select "raw" and choose JSON from the dropdown menu.
Enter the JSON data for creating a new person in the request body (e.g., {"name": "John", "age": 30, "phone_number": "123-456-7890", ...}).
Click on the "Send" button.

You should receive a response indicating that the person was created successfully.
