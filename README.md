# Projeto de Gerenciamento de Usuário com Golang e GORM

Este projeto é uma aplicação que utiliza a linguagem de programação Go, juntamente com os frameworks Gin e GORM, para criar um serviço de gerenciamento de usuários. Os dados dos usuários são persistidos em um banco de dados PostgreSQL, utilizando o serviço de banco de dados como serviço ElephantSQL. Além disso, este projeto também permite o armazenamento de objetos do usuário, como imagens, em buckets S3 utilizando o LocalStack.

---
![Golang](https://img.shields.io/badge/-Golang-00ADD8?style=flat-square&logo=go&logoColor=ffffff) ![PostgreSQL](https://custom-icon-badges.herokuapp.com/badge/PostgreSQL-025E8C.svg?logo=postgresql&logoColor=white) ![LocalStack](https://img.shields.io/badge/-LocalStack-00ADD8?style=flat-square&logo=amazon-aws&logoColor=white)

## Recursos

- :construction: Cadastro de usuários: permite o cadastro de novos usuários, incluindo informações como nome, e-mail e senha.
- :construction: Autenticação: oferece funcionalidades de autenticação e autorização, permitindo que usuários autenticados acessem recursos restritos.
- :construction: Gerenciamento de objetos do usuário: possibilita o armazenamento de objetos do usuário, como imagens, em buckets S3.
- Banco de dados: utiliza o GORM para interagir com o banco de dados Postgres hospedado no ElephantSQL, garantindo a persistência dos dados dos usuários.

## Configuração

Antes de executar o projeto, certifique-se de ter as seguintes dependências instaladas:

- Golang (versão 1.20.5): [https://golang.org/dl/](https://golang.org/dl/)
- GORM: Execute o seguinte comando para instalar o GORM:
  ```
  go get -u gorm.io/gorm
  ```
- PostgreSQL: Configure uma instância do PostgreSQL e obtenha as informações de conexão (endereço, porta, nome do banco de dados, usuário e senha).
- LocalStack: Instale e configure o LocalStack para simular o ambiente S3 localmente.

Depois de configurar as dependências e obter as informações de conexão, siga as etapas abaixo para executar o projeto:

1. Clone o repositório para o seu ambiente local:
   ```
   git clone https://github.com/guisteglich/go-user-app
   ```
2. Acesse o diretório do projeto:
   ```
   cd go-user-app
   ```
3. Configure variáveis de ambiente coomo informações de conexão com o banco de dados Postgres e o serviço S3 no arquivo `.env`.
4. Execute o comando para baixar as dependências do projeto:
   ```
   go mod download
   ```
5. Execute o seguinte comando para iniciar a aplicação:
   ```
   go run main.go
   ```
6. A aplicação será iniciada e estará disponível no endereço `http://localhost:8080`.

<!-- ## Endpoints -->
<!-- 
A aplicação disponibiliza os seguintes endpoints para interação:

- `POST /api/signup`: Cria um novo usuário com base nos dados fornecidos no corpo da solicitação.
- `POST /api/login`: Realiza a autenticação do usuário com base nas credenciais fornecidas.
- `GET /api/users`: Retorna a lista de usuários cadastrados (requer autenticação).
- `GET /api/users/{id}`: Retorna os detalhes de um usuário específico (requer autenticação).
- `POST /api/users/{id}/objects`: Faz upload de um objeto (imagem) associado a um usuário (requer autenticação).
- `GET /api/users/{id}/objects`: Retorna a lista de objetos (imagens) associados a um usuário (requ

er autenticação). -->

## Contribuição

Contribuições são bem-vindas! Se você encontrar algum problema, tiver ideias de melhorias ou quiser adicionar novos recursos, fique à vontade para abrir uma issue ou enviar um pull request.

## Licença

Este projeto está licenciado sob a [MIT License](LICENSE).


## Contato

Para entrar em contato, envie um email para guilherme.steglich16@gmail.com.

---
