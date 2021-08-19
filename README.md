# Описание проекта
Краткое описание проделанной работы для тестового задания — REST API для работы с товарами электронной коммерции.

* [Структура проекта](#project-structure)
* [Возможности проекта](#project-features)
    * [REST API](#project-features-rest-api)
    * [GraphQL](#project-features-graphql)
* [Локальный запуск](#run-local-app)
    * [Структура](#run-local-app-no-docker)
    * [Структура](#run-local-app-docker)
* [Локальный запуск тестов](#test-local-app)

<a name="project-structure"></a>
## Структура проекта
Проект состоит из следующих директорий и файлов:
* `cmd`.
* `database` — методы по работе с базой данных.
* `handlers` — методы API.
* `models` — объекты базы данных и приложения.
* `validation` — методы валидации входных параметров методов API.
* `vendor`, `go.mod` и `go.sum` — зависимости, необходимые для работы приложения.
* `.env` — переменные окружения.
* `Makefile` — команды для локального запуска приложения и тестов.
* `main.go` — точка входа в приложение.
* `main_test.go` — тесты для проверки работоспособности методов API.
* `Dockerfile` и `docker-compose.yml` — работа с докером.
* `index.html` и `swagger.yaml` — документация API.

<a name="project-features"></a>
## Возможности проекта
<a name="project-features-rest-api"></a>
### REST API
Список реализованных методов и их описания можно узнать в [документации API](https://raccoonsw.github.io/).

<a name="project-features-graphql"></a>
### GraphQL
Реализованы методы по работе с данными приложения с помощью GraphQL. Endpoint: `/api/graphql`.

* Добавление товара в каталог  
Пример запроса:
```graphql
mutation {
    insertItem(sku: "blabla", name: "wow bla", type: "virtual_good", cost: 1.05) {
        id
    }
}
```

* Получение информации о товаре по его идентификатору  
Пример запроса:
```graphql
query {
    getItem(id: 1) {
        id
        sku
        name
        type
        cost
    }
}
```

<a name="run-local-app"></a>
## Локальный запуск
<a name="run-local-app-no-docker"></a>
### Без докера
Чтобы локально запустить приложение:
1. Создайте пользователя и базу данных с такими параметрами, которые указаны в файле `.env`.
2. В командной строке в проекте выполнините `make run`.

Измените файл `.env`, если необходимо задать другие значения для запуска.

<a name="run-local-app-docker"></a>
### С докером
Чтобы запустить приложение в докере, в командной строке в проекте выполнините `docker-compose up`.  
Измените файлы `Dockerfile` и `docker-compose.yml`, если необходимо задать другие значения для запуска.

<a name="test-local-app"></a>
## Локальный запуск тестов
Чтобы локально запустить написанные тесты, в командной строке в проекте выполнините `make test`.
