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
3) Активируем виртуальное окружение и запускаем скомпилированный проект
```sourse {PYTHON_VE} && cd out_folder && python your_script.pyc```

##Флаги:
project.path - путь до проекта(абсолютный или относительный). По умолчанию == "."
out.path - путь куда сохраняется результат. По умолчанию = "./out_folder"


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
3) Activate the virtual environment and run the compiled project
`` `sourse {PYTHON_VE} && cd out_folder && python your_script.pyc```

## Flags:
project.path - path to the project (absolute or relative). Default == "."
out.path - the path where the result is present. Default = "./out_folder"