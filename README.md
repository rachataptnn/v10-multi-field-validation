# v10-multi-field-validation

## Overview
This repository serves as an illustrative example to showcase two distinct methods of handling error responses in a Go API:

1. **Simple Function**: Demonstrates handling error responses using a straightforward function approach.
2. **v10 Custom Validate Function**: Highlights the utilization of custom validation functions from the `v10` validator package to manage error responses.

## Usage

### Getting Started

1. **Import Postman Collection**: Import the provided Postman collection to swiftly access and test the API endpoints.

2. **Start Server**: Launch the Go server to begin handling API requests. You can do this by executing the command `go run main.go`.

3. **Run Example Requests**: Utilize the pre-configured requests within the imported Postman collection to interact with the API endpoints effectively.

### Exploring Error Handling Approaches

#### 1. Simple Function
The "Simple Function" approach is showcased in the `/normal-way` endpoint. It features a conventional error-handling strategy using standard Go functions. When querying this endpoint, the server employs a function named `normalValidate` to validate the provided parameters. This function's error response is a boolean value and an accompanying error message.

#### 2. v10 Custom Validate Function
The "v10 Custom Validate Function" approach is demonstrated in the `/v10-custom-func` endpoint. This method employs the `v10` validator package to create custom validation functions. The `customValidation` function evaluates the parameters and provides tailored error messages. These messages are then automatically managed by the `v10` validator, streamlining the error response process.

## Steps to Follow

1. Clone this repository to your local machine.

2. Import the Postman collection (`v10-multi-field-validation.postman_collection.json`) into your Postman application.

3. Open a terminal and navigate to the repository directory.

4. Start the Go server by running the command:
   ```
   go run main.go
   ```

5. Utilize the imported Postman collection to interact with the provided endpoints, observing the different error response mechanisms in action.

## Conclusion

By exploring these two methods of handling error responses, this repository aims to provide a clear understanding of effective error management strategies within a Go API.