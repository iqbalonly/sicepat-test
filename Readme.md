# sicepat BackEnd Test
Name : Muhammad Iqbal Rofikurrahman

Email : miqbalrofik@gmail.com

## Installation
To completely run this project you must have [docker & docker-compose](https://docker.com) installed on your PC, 
you'll also need `make` command to help the development.

First let's
```sh
make up
```
When all the containers are up, let's move on to the migration process. 
```sh
make migrate
```

And this  project is ready to execute.
[The postman](postman.json) documentation is available in this repository. 

If you are have more data to test, use your favorite database client to connect with `db-sicepat`.

MySQL instance `port = 3606` `user=root pass=test_pass`