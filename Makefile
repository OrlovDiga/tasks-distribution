.PHONY: build run clean

# Имя исполняемого файла
EXECUTABLE := ./bin/tgbot

# Путь к исходному файлу
MAIN_SRC_FILE := ./cmd/tgbot/main.go

# Команда для сборки приложения
build:
	@echo "Building..."
	go build -o $(EXECUTABLE) $(MAIN_SRC_FILE)

# Команда для запуска приложения
run: build
	@echo "Running..."
	./$(EXECUTABLE)

# Команда для очистки: удаляет исполняемый файл
clean:
	@echo "Cleaning..."
	go clean
	rm -f $(EXECUTABLE)
