#!/bin/bash

# Конфигурация базы данных
DB_HOST=postgres
DB_PORT=5432
DB_SCHEMA=anomaly
DB_USER=postgres 
DB_PASSWORD=root
DB_SSLMODE=disable

# Конфигурация сервера-передатчика
SERVER_TRANSMITTER_HOST=server_generator
SERVER_TRANSMITTER_PORT=8080


docker run --rm -it --network anomaly_detection_transmitter_network \
    -e SERVER_TRANSMITTER_HOST=$SERVER_TRANSMITTER_HOST \
    -e SERVER_TRANSMITTER_PORT=$SERVER_TRANSMITTER_PORT \
    -e DB_HOST=$DB_HOST \
    -e DB_PORT=$DB_PORT \
    -e DB_SCHEMA=$DB_SCHEMA \
    -e DB_USER=$DB_USER \
    -e DB_PASSWORD=$DB_PASSWORD \
    -e DB_SSLMODE=$DB_SSLMODE \
    -w $(pwd) test sh