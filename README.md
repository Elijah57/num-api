# num-api (Go)

A simple API built with Go that classifies numbers and returns interesting mathematical properties about them, along with a fun fact.

## Features
- **Prime Check**: Determines if the number is prime.
- **Perfect Check**: Checks if the number is a perfect number.
- **Odd/Even Check**: Determines if the number is odd or even.
- **Sum of Digits**: Calculates the sum of digits of the number.
- **Fun Fact**: Fetches a fun fact about the number from the Numbers API.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/Elijah57/num-api.git
   cd num-api
   ```

2. Install Go dependencies (if using modules):

   ```bash
   go mod tidy
   ```

3. Build the application:

   ```bash
   go build
   ```

4. Run the application:

   ```bash
   ./number-api-go
   ```

   The server will start running at `http://localhost:5080`.

## API Endpoint

### `GET /api/classify-number?number=<number>`

Classifies the given number and returns its properties along with a fun fact.

#### Example Request:

```
GET http://localhost:8080/api/classify-number?number=371
```

#### Example Response:

```json
{
    "number": 371,
    "is_prime": false,
    "is_perfect": false,
    "properties": ["odd"],
    "digit_sum": 11,
    "fun_fact": "371 is an Armstrong number because 3^3 + 7^3 + 1^3 = 371."
}
```

## Folder Structure

```
/number-api-go
├── src/
│   ├── controllers/       # Contains logic for handling requests
│   ├── utils/             # Contains utility functions for number checks (prime, perfect, etc.)
│   └── main.go            # Entry point of the application
├── go.mod                 # Go module file for dependencies
└── README.md              # Project documentation
```

## License

MIT License.
