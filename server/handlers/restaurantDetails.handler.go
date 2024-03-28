package handlers

import (
	"serveur/server/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RestaurantDetailsHandler(c *gin.Context) {
	var restaurantIdAsStr = c.Param("restaurantId")

	restaurantID, err := strconv.Atoi(restaurantIdAsStr)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid restaurant id"})
		return
	}
	restaurantDetails := services.GetRestaurantDetails(restaurantID)

	c.JSON(200, gin.H{"restaurant": restaurantDetails})
}

/*
Function that update the  affluence of the restaurant according to votes
*/
func UpdateAffluenceWithClientVoteHandler(c *gin.Context) {
	var restaurantId = c.Param("restaurantId")
	var vote = c.Param("level")

	restaurantID, err := strconv.ParseUint(restaurantId, 10, 64)

	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid restaurant id"})
		return
	}
	affluence, err := services.SubmitClientVoteForAffluence(restaurantID, vote)

	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid  Request"})
		return
	}
	c.JSON(200, gin.H{"Affluence": affluence})
}

func UpdateAffluenceWithRestaurantVoteHandler(c *gin.Context) {
	var restaurantId = c.Param("restaurantId")
	var vote = c.Param("level")

	restaurantID, err := strconv.Atoi(restaurantId)

	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid restaurant id"})
		return
	}
	affluence, err := services.UpdateAffluenceForRestaurantVote(restaurantID, vote)

	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid  Request"})
		return
	}
	c.JSON(200, gin.H{"Affluence": affluence})

}

/*
fUNCTION THAT gets the affluence of the restaurant (low | high | medium)
*/
func GetAffluenceHandler(c *gin.Context) {
	var restaurantId = c.Param("restaurantId")

	restaurantID, err := strconv.ParseUint(restaurantId, 10, 64)

	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid restaurant id"})
		return
	}
	affluence_level, err := services.GetAffluence(restaurantID)

	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid  Request"})
		return
	}

	c.JSON(200, gin.H{"Affluence": affluence_level})
}
