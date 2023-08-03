# test-task-makves-group

## Запустить проект:

```bash
make run
```

###

#### Получение элементов по ID через запятую

```bash
curl -X GET "localhost:8080/get-items?id=872,1048"
```

#### Получение элемента по одиночному ID

```bash
curl -X GET "localhost:8080/get-items?id=133"
```

#### Получение элементов по нескольким ID

```bash
curl -X GET "localhost:8080/get-items?id=133&id=520&id=728"
```
