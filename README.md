# go-grpc-api-gateway

### Set up Postgres database

Firs you will need to pull the lates postgres image from docker 

````docker pull postgres:latest```

The you will need to start a postgres instance, run the following command below to do so.

```docker run --name <container-name> -e POSTGRES_PASSWORD=<password> -d postgres```


### Create the database 

Then ssh into the container 
```docker exec -it <container-name> sh```

Then run the following commands to connect to postgres

```psql -u postgres```

Lets create the authentication service database 

```CREATE DATABASE auth_svc;```


Lets create the order service database 

```CREATE DATABASE order_svc;```


Let create the product service database 

```CREATE DATABASE product_svc;```

Lets list all the databases that exist 

```\l```

![databases created](./docs/images/Screenshot%202022-04-01%20at%2018.48.32.png)

To quit postgres we use this command 

```\q```