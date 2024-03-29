package middlewares

import (
	"log"
	"net/http"
	"serveur/server/models"
	"serveur/server/services"
	JwtService "serveur/server/services/jwt"

	"github.com/gin-gonic/gin"
)

func VerifyOrderMiddleware(c *gin.Context) {
	token := extractToken(c)

	claims := JwtService.ParseRestaurantAccessToken(token)
	if claims == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"success": 0, "message": "The provided token is invalid"})
		c.Abort()
		return
	}

	orderID := c.Param("orderId")

	var order models.OrderDetails

	order = services.GetOrderDetails(orderID)

	if order.RestaurantId != claims.Id {
		c.JSON(http.StatusUnauthorized, gin.H{"success": 0, "message": "Unauthorized to update status"})
		c.Abort()
		return
	}
	c.Next()
}

func OrderAuth(c *gin.Context) {
	token := extractToken(c)
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"success": 0, "message": "Unauthorized to access this order"})
		c.Abort()
		return
	}

	claims := JwtService.ParseRestaurantAccessToken(token)

	if claims == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"success": 0, "message": "The provided token is invalid"})
		c.Abort()
		return
	}

	orderID := c.Param("orderId")

	var order models.OrderDetails

	order = services.GetOrderDetails(orderID)

	if order.RestaurantId != claims.Id && order.ClientId != claims.Id {
		c.JSON(http.StatusUnauthorized, gin.H{"success": 0, "message": "Unauthorized to access this order"})
		c.Abort()
		return
	}
	c.Next()
}

func OrderClientAuth(c *gin.Context) {
	token := extractToken(c)
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"success": 0, "message": "Unauthorized to access this order"})
		c.Abort()
		return
	}

	clientClaims := JwtService.ParseClientAccessToken(token)

	if clientClaims == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"success": 0, "message": "The provided token is invalid"})
		c.Abort()
		return
	}

	orderID := c.Param("orderId")

	var orderDetails models.OrderDetails

	orderDetails = services.GetOrderDetails(orderID)
	// log client and order ids
	log.Println("client id: ", clientClaims.Id)
	log.Println("order client id: ", orderDetails.ClientId)

	if orderDetails.ClientId != clientClaims.Id {
		c.JSON(http.StatusUnauthorized, gin.H{"success": 0, "message": "Unauthorized to access this order"})
		c.Abort()
		return
	}
	c.Next()
}
