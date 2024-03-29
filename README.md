<h1 align="center">Задание №1</h1>
<h3 align="left">Необходимо создать сервис реализующий авторизацию через Google OAuth2.

Сервис должен выполнять следующий сценарий:
1.	Пользователь инициирует авторизацию перейдя на некий путь, например /auth
2.	Сервис перенаправляет пользователя на авторизацию через Google, где пользователь авторизуется введя свой логин и пароль
3.	Google перенаправляет пользователя обратно на сервис, подставляя в query данные для получения информации об авторизовавшемся пользователе
4.	Сервис должен получить данные о пользователе и авторизовать его в системе, автоматически создавая нового пользователя при его отсутствии в БД
5.	В результате авторизации сервис должен создать сессию пользователя, которую можно использовать через Cookie или заголовок Authorization
6.	При корректной работе механизма, пользователь должен иметь возможность запросить метод /me, который вернет информацию о текущем авторизованном пользователе или ошибку 401 при отсутствии авторизации

Требования:
1.	Использовать PostgreSQL15-16 версии в качестве базы данных. И данный драйвер для работы с PostgeSQL в GoLang - https://github.com/jackc/pgx (v5)
</h3>

<h1 align="center">Развертка</h1>

- Склонировать репозиторий
```
git clone https://github.com/Yury132/Golang-Task-1.git
```
- Установить PostgreSQL в Docker контейнер, используя docker-compose.yml файл из проекта
  
1. Скопировать docker-compose.yml в новую папку "postgresql"
  
2. Выполнить в терминале команду
```
docker compose up
```
- Подключиться к базе данных PostgreSQL (Например, через DBeaver)

POSTGRES_DB: mydb

POSTGRES_USER: root

POSTGRES_PASSWORD: mydbpass

Port: 5432

Host: localhost

- Скопировать полученный файл .env по пути Golang-Task-1/internal/config

- Запустить веб-приложение командой
```
go run cmd/main.go
```

<h1 align="center">Тестирование</h1>

- Перейти в браузер

```
http://localhost:8080
```

- Кнопка "Авторизация" - авторизация пользователя в системе через Google OAuth2.

![alt text](https://github.com/Yury132/Golang-Task-1/blob/main/forREADME/1.PNG?raw=true)

  Для нового пользователя создается запись в БД.

  Также создается сессия.

![alt text](https://github.com/Yury132/Golang-Task-1/blob/main/forREADME/2.PNG?raw=true)

- Кнопка "Информация обо мне" - получение имени пользователя и адреса электронной почты.

![alt text](https://github.com/Yury132/Golang-Task-1/blob/main/forREADME/3.PNG?raw=true)

  Если пользователь неавторизован - выдается сообщение об ошибке.
- Кнопка "Выход из системы" - окончание сессии, переход на главную страницу.
- Получение текущего списка пользователей в БД по адресу:

```
http://localhost:8080/users-list
```

![alt text](https://github.com/Yury132/Golang-Task-1/blob/main/forREADME/4.PNG?raw=true)

