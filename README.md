# Roman To Decimal Converter

A simple REST API to Convert Roman Numeral to an Decimal Value

[![Go](https://github.com/Bhupesh-V/roman-to-no-api/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/Bhupesh-V/roman-to-no-api/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/Bhupesh-V/roman-to-no-api/branch/main/graph/badge.svg?token=MK95H1VJVL)](https://codecov.io/gh/Bhupesh-V/roman-to-no-api)

---

NOTE:
this little project was given as a coding assignment for a Golang role(junior). I didn't get selected, might help someone

## Demo

![demo-api](https://user-images.githubusercontent.com/34342551/109393971-94726d80-794a-11eb-9632-ef1dbf33a941.gif)


## End Points

1. **POST**: `127.0.0.1:8080/roman`

   Post Body
   ```
   {"roman": "MCMXCIV"}
   ```

   Response
   ```
   {"integer":1994,"roman":"MCMXCIV"}
   ```

## Installation

1. Clone the repository
   ```bash
   git clone https://github.com/Bhupesh-V/roman-to-no-api
   cd roman-to-no-api
   ```

2. Install dependencies
   ```bash
   go get ./...
   ```

3. Run Tests
   ```bash
   go test ./... -v -coverpkg=./... -coverprofile=coverage.out
   go tool cover -html=coverage.out
   ```

4. Run Gin Development server
   ```bash
   go run main.go
   ```


## License

Copyright © 2021 [Bhupesh Varshney](https://github.com/Bhupesh-V).<br />
This project is [MIT](https://github.com/Bhupesh-V/areyouok/blob/master/LICENSE) licensed.

