# How to Run

1. `docker build --tag receipt-project-fetch .`

2. `docker run --rm -p 1323:1323 receipt-project-fetch`

# Testing

`go test -v`

# Usage

You can use the endpoints exactly as specified in the documentation! Using a tool such as postman or insomnia will make it much easier, and that is how this was tested and validated.

#### Examples

- `GET http://localhost:1323/receipts/{id}/points` - returns the points for the receipt with the passed in id (if it exists).
- `POST http://localhost:1323/receipts/process` - returns the id of the receipt that was processed if it was successful, or an error if it was not.
- `GET http://localhost:1323/receipts/` - an extra endpoint that returns all the receipts in the database that was helpful for troubleshooting (this would not be in a production environment). This can be helpful for you to verify the JSON representation.

#### Challenge

https://github.com/fetch-rewards/receipt-processor-challenge

#### Future Improvements

- Break code into modules to improve organization
- Create more tight coupling between the provided yml spec and the structs to encode and decode JSON payloads (I don't have experience doing this in Go and chose to leave it as-is for now)
