include .env

runserver:
	go run cmd/server/main.go
	
url-request:
	curl -X POST 'http://localhost:8000/url/short/' -H "Content-Type: application/json" -d '{"longUrl": "helloworld"}'

apply-migrations:
	@echo "applying the migrations..."
	@echo ""
	@goose -dir ./migrations postgres "postgres://${DB_USERNAME}:${DB_PASSWORD}@localhost:5432/url_shortener?sslmode=disable" up

github-push:
	@echo "pushing to GitHub..."
	@echo "" 
	@git push https://${GITHUB_USERNAME}:${GITHUB_PAT}@github.com/${GITHUB_USERNAME}/${GITHUB_REPO_NAME} master