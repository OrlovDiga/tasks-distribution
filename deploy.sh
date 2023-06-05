#!/bin/bash

# Обновите URL до вашего репозитория или zip-файла и имя исполняемого файла
REPO_URL="https://github.com/yourusername/yourproject.git"
EXECUTABLE="yourapp"

# Убедитесь, что git установлен
if ! command -v git &> /dev/null
then
    echo "git could not be found"
    exit
fi

# Убедитесь, что go установлен
if ! command -v go &> /dev/null
then
    echo "go could not be found"
    exit
fi

# Клонируем репозиторий
git clone $REPO_URL

# Переходим в директорию проекта
cd yourproject

# Собираем проект
go build -o $EXECUTABLE

# Запускаем проект
./$EXECUTABLE