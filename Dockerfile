# Используем официальный образ Go
FROM golang

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum для кэширования зависимостей
COPY go.mod ./
RUN go mod download

# Копируем остальные файлы
COPY . .

# Компилируем приложение
RUN go build -o server .

EXPOSE 8080

# Указываем команду для запуска приложения
CMD ["./server"]