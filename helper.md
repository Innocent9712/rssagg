goose:
    Helps handle migration for dbs.
Installation: 
    go install github.com/pressly/goose/v3/cmd/goose@latest
Usage:
    from the directory of the migration files, run `goose [db_type (e.g postgres)] [connection_string] up|down`


sqlc:
    Helps handle transpiling "text definitions and sql code" into type safe go equivalent functions.
Installation:
    go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
Usage:
    - define an sqlc.yaml file in the root of your project
    - run `sqlc generate` from the root directory
    - use the generated go code in your application


When setting up the db connection to the psql db, you need to add the postgres driver to the go.mod file by running the following command:
go get -u github.com/lib/pq

then add this to the imports of the file where the connection was initialized:
	_ "github.com/lib/pq"


Extra:

I setup a Makefile to build, run and clean the project
