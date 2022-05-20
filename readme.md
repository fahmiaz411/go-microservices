docker build -t go-app-normal:latest
docker run -d -p 80:80 --name web go-app-normal:latest
docker logs -f web
docker rm -f web

<!-- microservices -->

docker-compose build
docker-compose up -d
docker compose down
