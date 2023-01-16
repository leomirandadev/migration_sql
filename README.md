# MIGRATION SQL

To install you have to just run:
```shell
$ go install github.com/leomirandadev/migration_sql/cmd/migration_sql@latest
```

Support following SQL Databases:
- `mysql`
- `postgres`


## Create
```shell
    $ migration_sql --create migration_name
```
If you want to inform the dir, you can do:
```shell
    $ migration_sql --create migration_name --dir "./migrations_assets"
```

## Run
```shell
    $ migration_sql --up mysql --conn "root:root@tcp(127.0.0.1:3306)/your_database_name"
```

If you want to inform the dir, you can do:
```shell
    $ migration_sql --up mysql --conn "root:root@tcp(127.0.0.1:3306)/your_database_name" --dir "./migrations_assets"
```

## Rollback
```shell
    $ migration_sql --down mysql --conn "root:root@tcp(127.0.0.1:3306)/your_database_name"
```

If you want to inform the dir, you can do:
```shell
    $ migration_sql --down mysql --conn "root:root@tcp(127.0.0.1:3306)/your_database_name" --dir "./migrations_assets"
```