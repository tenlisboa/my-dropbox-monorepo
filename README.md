# My own Dropbox API

### Local dev containers
#### For the first time
```bash
docker run --name my-dropbox-postgres -v $(pwd):/tmp -e POSTGRES_PASSWORD=1234 -d postgres
docker run -d --hostname my-dropbox-q --name my-dropbox-rabbit -p 5672:5672 rabbitmq:3
```
#### From twice on
```bash
docker start my-dropbox-postgres
docker start my-dropbox-rabbit
```

To create the structure, you'll need to get into the pgsql container and follow the shell bellow:

```bash
docker exec -it my-dropbox-postgres bash

cd /tmp/backend/scripts/database

psql -U postgres
```

So, create the database
```bash
create database mydropbox;

\q
```

Now, we want to insert the SQL scripts into our brand new DB:
```bash
psql -U postgres mydropbox < users.sql
psql -U postgres mydropbox < folders.sql
psql -U postgres mydropbox < files.sql
```