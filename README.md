# REST-API Golang

- [x] go get github.com/gin-gonic/gin
- [x] go get gorm.io/gorm
- [x] go get github.com/dgrijalva/jwt-go
- [x] go get gorm.io/driver/mysql
- [x] go get github.com/joho/godotenv
- [x] go get github.com/mashingan/smapping
- [x] go get golang.org/x/crypto/bcrypt

## Operações possíveis

1. Editar o livro
2. Listar os livros por id
3. Postar um livro
4. Listar todos os livros
5. Deletar um livro
6. Criar um usuário
7. Logar
8. Editar
9. Buscar Livros relacionados ao meu usuario

## Banco de dados escolhido
- *MySQL*

## Canal que ensina sobre Golang
```https://www.youtube.com/channel/UCDG7gCaRawAdyeA5omGRBRA```

## Observação
- Precisa criar um arquivo .env com esses dados:
``` 
DB_USER= <seu user>
DB_PASS= <sua senha>
DB_HOST=localhost
DB_NAME= <nome database>

JWT_SECRET= <sua senha> 
```