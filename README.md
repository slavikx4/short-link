# short-link
Сборка и запуск
С помощью docker-compose

Запуск с in-memory хранилищем:
docker-compose build app_im
docker-compose run -p 127.0.0.1:8080:8080 -d app_im

Запуск с postgresql хранилищем:
docker-compose build app_db
docker-compose run -p 127.0.0.1:8080:8080 -d app_db

Из-за того, что долго провозился с докером - не успел написать тесты.
