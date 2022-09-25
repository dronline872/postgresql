# postgresql
Postgresql

# Миграции
- Создание новых миграций
```
$ bin/goose -dir migrations postgres "postgres://user:1234567@localhost:5432/lesson sslmode=disable" create name_migration sql
```

- Применить миграции
```
$ bin/goose -dir migrations postgres "postgres://user:1234567@localhost:5432/lesson" up
```

- Откат одной миграции
```
$ bin/goose -dir migrations postgres "postgres://user:1234567@localhost:5432/lesson" down
```