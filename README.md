# DEGO
Это учебный проект.
## Description
Сервис позволяет добавить человека только по имени. Остальные поля, такие как возраст, пол заполняются с помощью публичных API.
Данные человека добавляются в PostgreSQL
## Getting Started
Запускаем PostgreSQL в докере. 
```
docker pull postgres
docker run --name postgres -p 5432:5432 -e POSTGRES_PASSWORD={PASSWORD} -e POSTGRES_USER={USERNAME} -d postgres
```
В корневой папке проекта лежит .env файл, в нем нужно прописать параметры:
```
DBHOST=postgres://{USERNAME}:{PASSWORD}@{DBHOST}:5432/postgres
```
REST API методы:
```POST: /create``` - создает пользователя с именем Vasya и заполняет пустые поля данными из публичных API

Have:
```
{
    "name":"Vasya",
}
```

Got:
```
{
    "name":"Vasya",
    "surname":"",
    "patronymic":"",
    "age":49,
    "gender":"male",
}
```

### TODO {

```POST: /deleteById``` - удаляет пользователя с id=1 из БД
```
{
    "id": 1
}
```

```POST: /editById/1``` - вносит изменения в БД с id=1
Have:
```
{
    "name":"Vasya",
    "surname":"",
    "patronymic":"",
    "age":49,
    "gender":"male",
}
```

Want to change:
```
{
    "name":"",
    "surname":"Petrov",
    "patronymic":"Ivanovich",
    "age":56,
    "gender":"",
}
```

Got:
```
{
    "name":"Vasya",
    "surname":"Petrov",
    "patronymic":"Ivanovich",
    "age":56,
    "gender":"male",
}
```

```POST: /filter/name/Vasya``` - получает данные из бд с фильтром
```
{
    "name":"Vasya",
    "surname":"Petrov",
    "patronymic":"Ivanovich",
    "age":56,
    "gender":"male",
}
```

}
