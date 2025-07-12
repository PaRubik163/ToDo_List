📝 ToDo List на Go + Gin + SQLite3
Простой ToDo‑список с REST API, реализованный на Go c использованием Gin и SQLite3.

❓ Описание:  
Проект представляет собой backend‑сервис, позволяющий:  

добавлять задачи,

получать список всех задач,

обновлять статус выполнения,

удалять задачи.  

Используется:  

Go — язык реализации;

Gin — фреймворк для HTTP API;

database/sql с драйвером github.com/mattn/go-sqlite3 для работы с базой данных;

SQLite3 — лёгкая встроенная БД.

⚙️ Установка и запуск
<pre> bash 
git clone https://github.com/PaRubik163/ToDo_List 
cd ToDo_List/todoapp
go mod tidy 
go run cmd/main.go
</pre>
— после этого сервер будет доступен по адресу http://localhost:8060/tasks

| Метод    | Путь             | Описание                  | Тело запроса                        | Ответ                             |
| -------- | ---------------- | ------------------------- | ----------------------------------- | --------------------------------- |
| `POST`   | `/api/tasks`     | Создать новую задачу      | `{ "title": "...", "done": false }` | Созданная задача                  |
| `GET`    | `/api/tasks`     | Получить все задачи       | —                                   | `[{ "id": 1, ... }, ...]`         |
| `PUT`    | `/api/tasks/:id` | Обновить cтатус задачи    | `{ "title": "...", "done": true }`  | Обновлённая задача                |
| `DELETE` | `/api/tasks/:id` | Удалить задачу по ID      | —                                   | HTTP 200 OK или пустой JSON       |


💡 Основные разделы проекта
main.go
Создаёт экземпляр gin.Engine (gin.Default()).

Инициализирует подключение к SQLite3 через sql.Open("sqlite3", "tasks.db").

Регистрирует маршруты: POST, GET, PUT, DELETE.

Запускает HTTP‑сервер

models.go 
<pre>
type Task struct {
    Id    int    `json:"id"`
    Title string `json:"title"`
    Done  bool   `json:"done"`
}
</pre>
— структура задачи соответствует полям в базе и сериализуется в JSON.

handlers.go
Обработка входящего JSON через c.BindJSON(&task)

Вставка через DB.Exec(...) или DB.Prepare(...).Exec(...)

Обработка ошибок с c.JSON(http.StatusInternalServerError, gin.H{...})

Ответ с c.JSON(http.StatusOK, task) или списком задач.
