# MIGRATION SQL

To install you have to just run:
```shell
$ go install github.com/leomirandadev/migration_sql/cmd@latest
```

Support following SQL Databases:
- `mysql`
- `postgres`


## CREATE MIGRATION
```shell
    $ migration_sql create data2
```
If you want to inform the dir, you can do:
```shell
    $ migration_sql create data2 "./migrations_assets"
```

## RUN MIGRATION
```shell
    $ migration_sql up mysql "root:root@tcp(127.0.0.1:3306)/your_database_name"
```

If you want to inform the dir, you can do:
```shell
    $ migration_sql up mysql "root:root@tcp(127.0.0.1:3306)/your_database_name" "./migrations_assets"
```

## ROLLBACK MIGRATION
```shell
    $ migration_sql down mysql "root:root@tcp(127.0.0.1:3306)/your_database_name"
```

If you want to inform the dir, you can do:
```shell
    $ migration_sql down mysql "root:root@tcp(127.0.0.1:3306)/your_database_name" "./migrations_assets"
```