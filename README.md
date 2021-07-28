### Запуск

Для запуска/сборки проекта, необходимо скомпилировать `wallet-core`

*З.Ы. долго компилируется. сори, я не стал добавлять это в репу, файлы слишком толстые в
сборке, тем более я компилил под Mac OS*

```bash
make wallet-core
```

По умолчанию запускается RPC и REST сервера на портах 8081, 8080 соответственно

Примеры запросов находятся в `tests.sh`, так же можно запустить их с помощью команды

```bash
make sign_test
```

Создадутся HTTP запросы и точно такие же по RPC (файл `testing/main.go`)