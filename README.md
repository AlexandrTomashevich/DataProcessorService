## DataProcessorService

### Описание:
**Задача 1**

Написать SQL-запросы для ClickHouse:

- Выборки всех уникальных eventType у которых более 1000 событий.
- Выборки событий которые произошли в первый день каждого месяца.
- Выборки пользователей которые совершили более 3 различных eventType.

**Решение:**
- Файл part_1.sql

**Задача 2**

Реализовать на GO:

1) Вставку тестовых данных в таблицу events.
2) Вывод событий по заданному eventType и временному диапазону.

**Решение:**

1) В файле part_2.go происходит загрузка конфигурации, инициализация подключения к бд, чтение данных из файла **testdata.json**, вставка данных в таблицу бд, запуск и передача аргументов в функцию sortEvents.
В файле testdata.json записаны тестовые записи в формате json
2) В функции sortEvents происходит передача аргументов в запросе к бд, чтение ответа и вывод результата в терминале.

**Задача 3**

Реализовать на GO JSON API с методом отправки события в ClickHouse. Пример запроса:

POST /api/event
{
"eventType": "login",
"userID": 1,
"eventTime": "2023-04-09 13:00:00",
"payload": "{\"some_field\":\"some_value\"}"
}
**Решение** 
Сервис состоит из:
[main.go] - главная файл из которого запускается приложение
[connection.go] - в нем логика подключения к бд
[server.go] - в нем описывается логика работы нашего обработчиика (хендлера)
[config.go] - в нем описывается логика загрузка наших конфигураций
[models.go] - структура ответа пользователя
[config.yaml] - файл конфигурации
