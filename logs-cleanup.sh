#!/bin/bash

read -p "Очистить все log файлы? Опасность утери логов. [y/N]: " ans

if [ "$ans" = "y" ] || [ "$ans" = "Y" ]; then
    rm -rf out/logs
    echo "Файлы логов очищены"
else
    echo "Очистка логов отменена"
fi