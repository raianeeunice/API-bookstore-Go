package middleware

import (
	"livraria-go/helper"
	"livraria-go/service"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// AuthorizeJWT valida o token fornecido pelo usuário, retorna 401 se não for válido
func AuthorizeJWT(jwtService service.JWTService) gin.HandlerFunc{
	return func (c *gin.Context)  {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response := helper.BuildErrorResponse("Falha no processo de requisição", "Token não encontrado", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		token, err := jwtService.ValidateToken(authHeader)
		if token.Valid{
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[user_id]: ", claims["user_id"])
			log.Println("Claim[issuer]: ", claims["issuer"])
		}else{
			log.Println(err)
			response := helper.BuildErrorResponse("Token não validado", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}