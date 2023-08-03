# test-task-makves-group

## Запустить проект можно командой:

```shell
make run
```
или
```shell
go run cmd/web/main.go
```

###

#### Получение элементов по ID через запятую

```shell
curl -X GET "localhost:8080/get-items?id=872,1048"
```

#### Получение элемента по одиночному ID

```shell
curl -X GET "localhost:8080/get-items?id=133"
```

#### Получение элементов по нескольким ID

```shell
curl -X GET "localhost:8080/get-items?id=133&id=520&id=728"
```
