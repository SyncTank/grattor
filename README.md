# grattor
go - Blog Aggregator

Initial Project was setup using these tools 
psql for Postgres


Goose for database migrations, setups are in SQL folder 
goose postgres <connection_string> up
# example:
# goose postgres "postgres://postgres:@localhost:5432/gator" up
# goose postgres "postgres://postgres:@localhost:5432/gator" down

The connection string protocal follow this :
# protocol://username:password@host:port/database?sslmode=disable

Additional arguments can be passed for the password after compile
go build
./grattor login Timmonthy 123
