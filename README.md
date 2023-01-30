# task

1. Клиент сервис

    _go run tasks_client/main.go_

**POST** http://localhost:8080/task/
   {
   "method": "POST",
   "url": "https://google.com",
   "headers": {"Accept-Language":"kz"}
   }
**GET** http://localhost:8080/task/1


2. Сервис-краулер(работает по порту 50051 через gRPC)

   _go run tasks_service/main.go_
