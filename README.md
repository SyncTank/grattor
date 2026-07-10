# grattor
go - Blog Aggregator

postgres is used for the database 

If installed locally
```
sudo apt update
sudo apt install postgresql postgresql-contrib
sudo -u postgres psql
```

Your instance may need to be started
```
systemctl status postgresql
sudo systemctl start postgresql
```

I usually setup a pod for the application which you'll find it setup for.
```
podman exec -it <container_name> psql -U postgres
```

or if installed psql
```
psql -h localhost -p 5432 -U postgres -d postgres
```

To handle any database migrations the project has been setup to use Goose
```
go install github.com/pressly/goose/v3/cmd/goose@latest
```

One more tool to note is sqlc used to generate sql queries for us
```
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

For reindexing :
ALTER DATABASE postgres REFRESH COLLATION VERSION;
REINDEX DATABASE postgres;

Connection String : file called .gatorjson
{"db_url": ["postgres://", "?sslmode=disable"],"current_user_name":"postgres", "DBString":"", "password": ""}

Goose for database migrations, setups are in SQL folder 
goose postgres <connection_string> up
# example:
```
- goose postgres "postgres://postgres:@localhost:5432/gator" up
- goose postgres "postgres://postgres:@localhost:5432/gator" down
```

The connection string protocal follow this :
```
- goose postgres "postgres://postgres:123@localhost:5432/gator" up
- protocol://username:password@host:port/database?sslmode=disable
```

postgres://postgres:123@localhost:5432/database?sslmode=disable

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

