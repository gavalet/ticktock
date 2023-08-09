# TickTock API
TickTock API is a webservice that its unique perpuse is to generate timestamps between two times with a certain period. It is alligned with time zone and supports DST (daylight saving time). It does not need any authorization to perform a question to system. 

# How to execute and test
## EndPoints
End-points are described in swagger file. Swagger file is ./swagger.yaml

## Using Docker
```
sudo docker-compose up  -d --build
```
API will listen to  address "172.16.0.2" and port "8080".
You can get your first timestamps by executing 
```
curl --location 'http://172.16.0.2:8080/ptlist?period=1h&tz=Europe%2FAthens&t1=20210714T204603Z&t2=20210715T123456Z'
```
## Local maching (no Docker)
You can run API in your local machine by executing:
```
go run cmd/ticktock/main.go 127.0.0.1 8080
```
You can get your first timestamps by executing 
```
curl --location 'http://localhost:8080/ptlist?period=1h&tz=Europe%2FAthens&t1=20210714T204603Z&t2=20210715T123456Z'

```
### Test the API
Unit tests are write to validate the generated output.
Tests are executed by using Go's test package

```
go test ./models/ -v
```

# Extent the supported periods
API is written in a way for easy extension of supported period. Define the new periods and their functions
in the ./models/timestamps.go file at the funcMap variable.

# Contributing
We welcome contributions from the community! Whether you're a developer, designer, or enthusiast, your input is valuable and can help make this project even better. 



