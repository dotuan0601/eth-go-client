Work with ETH using golang

How to run?
Step 1:
- run: go mod vendor
- change eth config in main.go
- run: go run main.go

How to test?
- call api via postman:
example: localhost:8080/api/v1/eth/get-balance?address=0xffbcd481c1330e180879b4d2b9b50642eea43c02