# Python Packer
Сервис предназначен для преобразования открытого кода python в полузакрытый или зашифрованный для защиты кода.

##Требования:
ОС: Linux(Debian/Ubuntu)
Golang: 1.14

## Сборка и компиляция
```make```

## Алгоритм работы:
1) Скомпилируйте свой python проект
```python -m compileall .```
2) Запустите скрипт по переносу файлов
```go run cmd/main.go --project.path=/path/to/your/project```
Скрипт скопирует все файлы, а все .py файлы заменит скомпилированными .pyc файлами


# Python Packer
The service is designed to convert open python code to semi-closed or encrypted to protect the code.

## Requirements:
OS: Linux (Debian / Ubuntu)
Golang: 1.14

## Build and compile
```make```

## Work algorithm:
1) Compile your python project
`` `python -m compileall .```
2) Run the file transfer script
`` `go run cmd / main.go --project.path = /path/to/your/project```
The script copies all files, and all .py files are replaced by compiled .pyc files