Startup

Running the backend:

cd /Users/marcosviniciusdepaivacelio/go_projects/interview-accountapi && docker-compose up

Cli app

Tests

1. cd /Users/marcosviniciusdepaivacelio/go_projects/cli-app/github.com/marcoscelio/cli-app

2. Create an account
   go test cli_test.go cli_create_test.go

3. Fetch all accounts
   go test cli_test.go cli_fetch_test.go

4. Delete an account

   go test cli_test.go cli_delete_test.go
