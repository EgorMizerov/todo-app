# todo-app

## Пакеты
1. Веб-фреймворк - <a href="https://github.com/gin-gonic">gin</a>
2. Логгер - <a href="https://github.com/sirupsen/logrus">logrus</a>
3. Работа с базой данных - <a href="https://github.com/jmoiron/sqlx">sqlx</a>
4. Аутентификация - <a href="https://github.com/dgrijalva/jwt-go">jwt-go</a>
5. Парсер конфигураций - <a href="https://github.com/spf13/viper">viper</a>

## Технологии
1. Rest API
2. Postgres
3. Clean Architecture
## API
```
POST /auth/sign-up - Регистрация
POST /auth/sign-in - Авторизация

GET /lists        - Список списков
GET /lists/:id    - Получить список
POST /lists       - Создать список
PUT /lists/:id    - Изменить список
DELETE /lists/:id - Удалить список

GET /lists/:id/items        - Список задач
POST /lists/:id/items       - Создать задачу
PUT /lists/:id/items/:id    - Изменить задачу
GET /lists/:id/items/:id    - Получить задачу
DELETE /lists/:id/items/:id - Удалить задачу
```
