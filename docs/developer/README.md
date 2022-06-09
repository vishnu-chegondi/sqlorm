# Developer Documentation

Thanks for showing interest in contributing. ORM development is supported with the help of ormdrivers listed [below](#drivers). Drivers are developed in such a way that they will return statements which can be used for building queries with different data. 

In the current package, we have different [sub-packages](#sub-packages) to handle different functionalities.


## Sub Packages

* [client](#client)
* [migrations](#migrations)
* [models](#models)
* [queries](#queries)

### Client

> TODO: Client handles all the functionalities related to meta data querying and run ddl queries.

### Migrations

> TODO: Register tables to be migrated.

### Models

> TODO: Table Interface

### Queries

> TODO: Object Interface extending Table Interface

## Drivers

Currently the following are the databases supported. Pointing towards the orm drivers.

1. [MySql][mysql_driver]


[mysql_driver]:https://github.com/codeamenity/mysql_ormdriver