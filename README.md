# roman-to-no-api

A simple REST API to Convert Roman Numeral to an Integral Value

[![Go](https://github.com/Bhupesh-V/roman-to-no-api/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/Bhupesh-V/roman-to-no-api/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/Bhupesh-V/roman-to-no-api/branch/main/graph/badge.svg?token=MK95H1VJVL)](https://codecov.io/gh/Bhupesh-V/roman-to-no-api)

## End Points

1. **POST**: `127.0.0.1:8080/roman`

   Post Body
   ```
   {"roman": "MCMXCIV"}
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

4. Executing
   ```bash
   go run main.go
   ```

