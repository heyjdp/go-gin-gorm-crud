# CRUD app in golang using Gin and GORM

1. Fetch and install the Compile Daemon

https://github.com/githubnemo/CompileDaemon

```bash
go get github.com/githubnemo/CompileDaemon
go install github.com/githubnemo/CompileDaemon
```

2. Fetch GoDotEnv for environment variables

https://github.com/joho/godotenv

```bash
go get github.com/joho/godotenv
```

3. Add Gin web framework

https://github.com/gin-gonic/gin

```bash
go get -u github.com/gin-gonic/gin
```

4. Download and include GORM

https://gorm.io/

```bash
go get -u gorm.io/gorm
```

5. Add a `main.go` and run the Compile Daemon

```bash
cat main.go
package main

import "fmt"

func main() {
	fmt.Println("Hello world")
}
```

And in the same directory:

```bash
CompileDaemon -command="./go-gin-gorm-crud"
```

And it will auto build upon saved changes

6. Setup `.env` file

```bash
echo PORT=3000 > .env
cat .env
PORT=3000
DB_INFO="host=localhost user=heyjdp password=password123 dbname=go_crud_demo port=5432 sslmode=disable TimeZone=Asia/Nicosia"
```

And see usage examples here: https://github.com/joho/godotenv

7. Setup postgres

**NOTE:** Assuming you are on Debian Linux and installed via `apt`

First, create the role by issuing the following command:

```bash
sudo su - postgres -c "createuser heyjdp"
```

And make a password for the user:

```bash
sudo -u postgres psql -c "ALTER USER heyjdp PASSWORD 'password123';"
```

Next, create the database using the createdb command:

```bash
sudo su - postgres -c "createdb go_crud_demo"
```

To grant permissions to the user on the database, connect to the PostgreSQL shell:

```bash
sudo -u postgres psql
```

Run the following query:

```bash
GRANT ALL PRIVILEGES ON DATABASE go_crud_demo TO heyjdp;
```

And use `\q` to quit the postgres shell

**NOTE:** By default postgres is listening on port 5432 and is limited to connections from localhost. This is a good thing if you are going to make a dumbass password like above ;)

```bash
ss -nlt | grep 5432
LISTEN   0        128                 127.0.0.1:5432             0.0.0.0:*      
LISTEN   0        128                     [::1]:5432                [::]:*  
```

8. Migrate the models to the DB

```bash
go run migrate/migrate.go
```

And check the DB tables:

```bash
sudo -u postgres psql
postgres=# \c go_crud_demo \d
You are now connected to database "go_crud_demo" as user "postgres".
             List of relations
 Schema |     Name     |   Type   | Owner  
--------+--------------+----------+--------
 public | posts        | table    | heyjdp
 public | posts_id_seq | sequence | heyjdp
(2 rows)
```

**NOTE:** GUI software such as DBeaver (https://dbeaver.io/) or pgAdmin (https://www.pgadmin.org/https://www.pgadmin.org/) will do the same job

