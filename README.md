# grattor
Go - Blog Aggregator

Postgres is used for the database 
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

I usually setup a pod/container for the application which you'll find it setup for below.
To initilize a db manually you can do it as below. Keep in mind the volume, I setup a custom volume pg-data for this project.
```
podman run -d \
  --name postgres-test \
  --network [network] \
  -e POSTGRES_PASSWORD=[password] \
  -p 5432:5432 \
  -v [volume]:/var/lib/postgresql/data \
  postgres
```

Ex.
```
podman run -d --name postgres-test \
        --network bridge \
        -e POSTGRES_PASSWORD=123 \
        -p 5432:5432 \
        -v pg-data:/var/lib/postgresql \
        postgres
```

After the initialization of the container you can connect to like so
```
podman exec -it <container_name> psql -U postgres
```
or if installed psql
```
psql -h localhost -p 5432 -U postgres -d postgres
```
The default should work just fine, just make sure to check the password and that the database exists

To handle any database migrations the project has been setup to use Goose and sqlc used to generate sql queries for us
```
go install github.com/pressly/goose/v3/cmd/goose@latest
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

If it is your first time you'll need to create the database in question `CREATE DATABASE gator`
Once that is I recommend connecting to it and verifying `\c gator` and then run `SELECT version();` for the sanity check

# Reindexing :
The Postgres database can get out of wack due to upgrade done to the system. You can use  goose as mention, but sometimes you need to refresh the instance which is done below.
```
ALTER DATABASE postgres REFRESH COLLATION VERSION;
REINDEX DATABASE postgres;
```

Connection String : Is place in a file called .gatorjson at root of the project
```
{"db_url": ["postgres://", "?sslmode=disable"],"current_user_name":"postgres", "DBString":"", "password": ""}
```
postgres://postgres:123@localhost:5432/database?sslmode=disable

# Example goose:

Goose for database migrations, setups are in SQL folder 
goose postgres <connection_string> up
```
- goose postgres "postgres://postgres:@localhost:5432/gator" up
- goose postgres "postgres://postgres:@localhost:5432/gator" down
```

The connection string protocal follow this :
```
- goose postgres "postgres://postgres:123@localhost:5432/gator" up
- protocol://username:password@host:port/database?sslmode=disable
```

go build
./grattor login Timmonthy 123


