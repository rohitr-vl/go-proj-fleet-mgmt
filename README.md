# go-proj-fleet-mgmt
Golang practice project Fleet tracking and mgmt application
# initiate project
1. cd practice-projects/go-proj-fleet-mgmt/
2. go mod init fleetmgmt/web
3.  cmd/api/main.go file initializes an HTTP webserver on a specific port & a service initializer and dependency injection of other modules that the service depends on.
    3.a. The cmd directory has the packages responsible for external accesses.
    3.b. The api package that implements the REST/go-chi layer
4. config/config.go contains environment variables and parameters externally mapped to the project
5. internal/db: create sqlc.yaml, schema.sql & query.sql. Since sqlc is already installed, Execute command "sqlc generate" 