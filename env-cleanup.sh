#!/bin/bash

read -p "Очистить все volume файлы окружения? Опасность утери данных. [y/N]: " ans

if [ "$ans" = "y" ] || [ "$ans" = "Y" ]; then
    docker compose down todoapp-postgres && \
    rm -rf out/pgdata
    echo "Файлы окружения очищены"
else
    echo "Очистка окружения отменена"
fi