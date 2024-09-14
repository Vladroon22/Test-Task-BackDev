# Тестовое задание на позицию Junior Backend Developer

<h2>Configuration</h2>

`
sudo docker run --name=testDb -e POSTGRES_PASSWORD=11111 -p 5431:5432 -d postgres:16.2

`

`
export DB="user:11111@localhost:5431/postgres?sslmode=disable"
export JWT="imagine your own secret key"
export email="your email" (mail.ru)
export AppPass="your application password"

`
<h2>How to run</h2>

`
make run

`