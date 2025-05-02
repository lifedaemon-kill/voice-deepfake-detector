# Используем официальный образ Python версии 3.11
FROM python:3.11 AS builder

# Устанавливаем необходимые пакеты, включая ffmpeg
RUN apt-get update
RUN apt-get install -y ffmpeg
RUN rm -rf /var/lib/apt/lists/*
RUN pip install --upgrade pip

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем requirements.txt для кэширования зависимостей
COPY requirements.txt ./

# Устанавливаем зависимости
RUN pip install --no-cache-dir -r requirements.txt

# Копируем исходный код
COPY . .

# Запускаем приложение
CMD ["python", "./cmd/main.py"]
