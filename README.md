# DEGO
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
    "surname": "Petrov"
}
```

Got:
```
{
    "id": "4",
        "name": "Vasya",
        "surname": "Petrov",
        "patronymic": "",
        "age": 50,
        "gender": "male",
        "country": [
            {
                "country_id": "UA",
                "probability": 0.471
            },
            {
                "country_id": "RU",
                "probability": 0.202
            },
            {
                "country_id": "BY",
                "probability": 0.083
            },
            {
                "country_id": "BG",
                "probability": 0.04
            },
            {
                "country_id": "KZ",
                "probability": 0.029
            }
        ]
}
```

```GET: /all``` - получает всех добавленных людей из БД

Got:
```
{
    {
            "id": "3",
            "name": "Michal",
            "surname": "Michalov",
            "patronymic": "",
            "age": 48,
            "gender": "male",
            "country": [
                {
                    "country_id": "CZ",
                    "probability": 0.372
                },
                {
                    "country_id": "IE",
                    "probability": 0.015
                },
                {
                    "country_id": "IL",
                    "probability": 0.153
                },
                {
                    "country_id": "PL",
                    "probability": 0.079
                },
                {
                    "country_id": "SK",
                    "probability": 0.305
                }
            ]
        },
        {
            "id": "4",
            "name": "Vasya",
            "surname": "Petrov",
            "patronymic": "",
            "age": 50,
            "gender": "male",
            "country": [
                {
                    "country_id": "BG",
                    "probability": 0.04
                },
                {
                    "country_id": "BY",
                    "probability": 0.083
                },
                {
                    "country_id": "KZ",
                    "probability": 0.029
                },
                {
                    "country_id": "RU",
                    "probability": 0.202
                },
                {
                    "country_id": "UA",
                    "probability": 0.471
                }
            ]
        }
    ]
}
```




```GET: /delete/1``` - удаляет пользователя с id=1 из БД

```POST: /update/1``` - вносит изменения в БД с id=1
Have:
```
{
    "name":"Vasya",
    "surname":"",
    "patronymic":"",
    "age":49,
    "gender":"male"
}
```

Want to change:
```
{
    "name":"",
    "surname":"Petrov",
    "patronymic":"Ivanovich",
    "age":56,
    "gender":""
}
```

Got:
```
{
    "name":"Vasya",
    "surname":"Petrov",
    "patronymic":"Ivanovich",
    "age":56,
    "gender":"male"
}
```

### TODO {
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
