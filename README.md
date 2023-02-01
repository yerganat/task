# task
1. Клиент сервис

    _go run tasks_client/main.go_


JSON:

**POST** http://localhost:8080/task/
  
 {
   "method": "POST",
   "url": "https://google.com",
   "headers": {"Accept-Language":"kz"}
   }

**GET** http://localhost:8080/task/1


2. Сервис-краулер(работает по порту 50051 через gRPC)

   _go run tasks_service/main.go_

   _Описани сервиса  entity/tasks.proto_



- Структура проекта не соответствует Go Standard Project Layout
- Нет Graceful Shutdown сервера
- Нет очереди задач (требование ТЗ)
- Нет интерфейсов
- Нет юнит-тестов для проверки сервера
- Нет докерфайла и мейкфайла