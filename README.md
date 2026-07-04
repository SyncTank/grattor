# grattor
go - Blog Aggregator

Initial Project was setup using these tools 
psql for Postgres
systemctl status postgresql
sudo systemctl start postgresql

For reindexing :
ALTER DATABASE postgres REFRESH COLLATION VERSION;
REINDEX DATABASE postgres;

Connection String : file called .gatorjson
{"db_url": ["postgres://", "?sslmode=disable"],"current_user_name":"postgres", "DBString":"", "password": ""}

Goose for database migrations, setups are in SQL folder 
goose postgres <connection_string> up
# example:
- goose postgres "postgres://postgres:@localhost:5432/gator" up
- goose postgres "postgres://postgres:@localhost:5432/gator" down

The connection string protocal follow this :
goose postgres < postgres://postgres:postgres@localhost:5432/gator > up
- protocol://username:password@host:port/database?sslmode=disable

Additional arguments can be passed for the password after compile
go build
./grattor login Timmonthy 123

For a continer build you can use this for the db

podman run -d \
  --name postgres-test \
  --network [network] \
  -e POSTGRES_PASSWORD=[password] \
  -p 5432:5432 \
  -v [volume]:/var/lib/postgresql/data \
  postgres

Ex.
podman run -d --name postgres-test \
        --network bridge \
        -e POSTGRES_PASSWORD=123 \
        -p 5432:5432 \
        -v pg-data:/var/lib/postgresql \
        postgres

If setup like so you can use podman / docker or psql to hop in. 

podman exec -it postgres-test psql -U postgres
or 
psql -h 127.0.0.1 -p 5432 -U postgres -d postgres

