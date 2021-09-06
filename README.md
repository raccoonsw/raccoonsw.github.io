# Описание проекта
Краткое описание проделанной работы для тестового задания — REST API для работы с товарами электронной коммерции.

* [Структура проекта](#project-structure)
* [Возможности проекта](#project-features)
    * [REST API](#project-features-rest-api)
    * [GraphQL](#project-features-graphql)
    * [gRPC](#project-features-grpc)
* [Локальный запуск](#run-local-app)
    * [Структура](#run-local-app-no-docker)
    * [Структура](#run-local-app-docker)
* [Локальный запуск тестов](#test-local-app)

<a name="project-structure"></a>
## Структура проекта
Проект состоит из следующих директорий и файлов:
* `cmd`:
    * `grpc` — сервис по отправке письма о заказе товара. Содержит файлы `.env` и `main.go`.
    * `rest_api` — приложение для работы с товарами. Содержит файлы `.env`, `main.go` и `main_test.go`.
* `database` — методы по работе с базой данных.
* `grpc_email_client` — передача данных из приложения в сервис `grpc`.
* `grpc_email_server` — подготовка письма и его отправка по протоколу SMTP.
* `handlers` — методы REST API и GraphQL.
* `models` — объекты базы данных и приложения.
* `validation` — методы валидации входных параметров методов REST API.
* `vendor`, `go.mod` и `go.sum` — зависимости, необходимые для работы приложения.
* `Makefile` — команды для локального запуска приложения и тестов.
* `docker-compose.yml`, `Dockerfile.grpc` и `Dockerfile.web` — работа с докером.
* `index.html` и `swagger.yaml` — документация REST API.

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
POST /api/graphql HTTP/1.1

mutation {
    insertItem(sku: "blabla", name: "wow bla", type: "virtual_good", cost: 1.05) {
        id
    }
}
```

* Получение информации о товаре по его идентификатору  
  Пример запроса:
```graphql
POST /api/graphql HTTP/1.1

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
<a name="project-features-grpc"></a>
### gRPC
Реализован сервис по отправке письма о заказе товара. Endpoint: `/api/orders`.

Чтобы использовать сервис, укажите в файле `/cmd/grpc/.env` значение переменных:
* `GRPC_USEREMAIL` — email-адрес пользователя, от чъего имени отправляется письмо.
* `GRPC_USERPASSWORD` — пароль для указанного email-адреса.

Работа сервиса проверялась с помощью почтового сервиса Gmail. Для корректной работы через Gmail может потребоваться настройка:
1. Войдите в учетную запись Gmail.
2. Перейдите по ссылке `https://www.google.com/settings/security/lesssecureapps`.
3. Установите переключатель **Небезопасные приложения разрешены** в активное положение.
4. Если это не решило проблему, можно попробовать [другие предложенные решения](https://serverfault.com/questions/635139/how-to-fix-send-mail-authorization-failed-534-5-7-14).

Пример запроса:
```http request
POST /api/orders HTTP/1.1
Host: localhost
Content-Type: application/json

{
    "item_id": 1,
    "email": "email.to.this.user@gmail.com"
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
