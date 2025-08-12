# weather_service
http server that serves the current weather using given lat/long coordinates. serves a short forecast for that area for Today/Tonight

i.e Hot, Cold, Moderate depending on a floridans sense of temperature. 

the service is driven by National Weather Service API; specifically /points & /forecast endpoints

# how to run
1. clone this repo using git clone
   1. git clone https://github.com/jsonballadares/weather_service.git
   2. cd weather_service
   3. go version (make sure go is installed)
2. go mod download (download dependencies)
3. go build -o weather_service (**optional**)
   1. ./weather_service
4. go run ./
# some happy path test cases
1. curl "http://localhost:8080/shortForecast?latitude=38.8894&longitude=-77.0352"
2. curl "http://localhost:8080/shortForecast?latitude=40.730610&longitude=-73.935242"
3. curl "http://localhost:8080/shortForecast?latitude=25.7617&longitude=-80.1918"
# some rainy day test cases
1. curl "http://localhost:8080/shortForecast?longitude=-80.1918"
2. curl "http://localhost:8080/shortForecast?latitude=25.7617"
3. curl "http://localhost:8080/shortForecast?latitude=abc&longitude=xyz"
