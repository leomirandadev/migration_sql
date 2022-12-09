# MIGRATION SQL

## CREATE MIGRATION
```shell
    $ migration_sql create data2
```
If you want to inform the dir, you can do:
```shell
    $ migration_sql create data2 "./assets"
```

## RUN MIGRATION
```shell
    $ migration_sql up mysql "root:root@tcp(127.0.0.1:3306)/your_database_name"
```

If you want to inform the dir, you can do:
```shell
    $ migration_sql up mysql "root:root@tcp(127.0.0.1:3306)/your_database_name" "./assets"
```

## ROLLBACK MIGRATION
```shell
    $ migration_sql down mysql "root:root@tcp(127.0.0.1:3306)/your_database_name"
```

If you want to inform the dir, you can do:
```shell
    $ migration_sql down mysql "root:root@tcp(127.0.0.1:3306)/your_database_name" "./assets"
```