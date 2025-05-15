runserver:
	go run cmd/server/main.go
	
url-request:
	curl -X POST 'http://localhost:8000/url/short/' -H "Content-Type: application/json" -d '{"longUrl": "helloworld"}'