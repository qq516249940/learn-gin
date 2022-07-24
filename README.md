# README.md

docker run --name some-postgres -e POSTGRES_PASSWORD=mysecretpassword -d postgres
docker run --name postgresql -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=root123456 -p 5432:5432 -d postgres
