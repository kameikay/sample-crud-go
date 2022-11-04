
# CRUD - Go

Um simples CRUD em Go.


## Stack utilizada

**Back-end:** Go
**Banco de dados:** MySQL
**Container:** Docker


## Rodando localmente

Clone o projeto

```bash
  git clone https://github.com/kameikay/sample-crud-go.git
```

Entre no diretório do projeto

```bash
  cd sample-crud-go
```

Inicie o docker com docker-compose

```bash
  docker-compose up -d
```

Execute o bash do MySQL dentro do container

```bash
  docker-compose exec mysql bash
```

Dentro do container, acesse o banco de dados pelo bash

```bash
  mysql -uroot -p goexpert
```

```bash
  CREATE TABLE products(id varchar(255), name varchar(80), price decimar(10,2), primary key (id))
```