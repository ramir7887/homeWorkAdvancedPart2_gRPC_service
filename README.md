# Домашняя работа по Advanced ч.2 gRPC service

Домашняя работа по Advanced ч.2 gRPC service
Реализация gRPC сервиса из домашнего задания. Задание: Сервис доставки

## Задание: Сервис заказов и сервис доставки

Реализован gRPC сервис доставки.
По умолчанию доступен: http://localhoct:8080

В папке uigrpc есть файлы проекта для Kreya.

protofile: https://gitlab.com/dive-into-golang-course/r_r_abdulov/homework_advanced_part2_proto

Команда для генирации:

    protoc --proto_path=. --go_out=. --go-grpc_out=. internal/pb/delivery.proto


