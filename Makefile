include .env

runserver:
	go run cmd/server/main.go
	
apply-migrations:
	@echo "applying the migrations..."
	@echo ""
	@goose -dir ./migrations postgres "postgres://${DB_USERNAME}:${DB_PASSWORD}@localhost:5432/url_shortener?sslmode=disable" up

github-push:
	@echo "pushing to GitHub..."
	@echo "" 
	@git push https://${GITHUB_USERNAME}:${GITHUB_PAT}@github.com/${GITHUB_USERNAME}/${GITHUB_REPO_NAME} master
	@echo ""

short-url:
	curl -X POST 'http://localhost:8000/url/short/' -H "Content-Type: application/json" -d '{"longUrl": "https://music.youtube.com/watch?v=NFODb2JRXBI"}'

get-long-url:
	@echo "Enter the short URL:"
	@sh -c 'read -r shortUrl; curl -X GET "http://localhost:8000/$$shortUrl"'
	