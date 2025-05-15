include .env

runserver:
	go run cmd/server/main.go
	
url-request:
	curl -X POST 'http://localhost:8000/url/short/' -H "Content-Type: application/json" -d '{"longUrl": "helloworld"}'
	
github-push:
	@echo "pushing to GitHub..."
	@echo "" 
	@git push https://${GITHUB_USERNAME}:${GITHUB_PAT}@github.com/${GITHUB_USERNAME}/${GITHUB_REPO_NAME} master