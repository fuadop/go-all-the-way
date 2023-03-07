package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

type Chef struct {
	Id                string `json:"id"`
	Name              string `json:"name"`
	Country           string `json:"country"`
	YearsOfExperience int8   `json:"yearsOfExperience"`
}

type ExtendedChef struct {
	Chef
	Recipes []Recipe `json:"recipes"`
}

func NewChefHandler(c *gin.Context) {
	var chef Chef
	if err := c.ShouldBindJSON(&chef); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	chef.Id = xid.New().String()
	chefs = append(chefs, chef)

	chefsIndex[chef.Id] = &chef // update index

	c.JSON(http.StatusCreated, chef)
}

func ListChefsHandler(c *gin.Context) {
	var shouldPopulate bool
	if populate := c.Query("populate"); populate != "" {
		if populate != "no" && populate != "false" {
			shouldPopulate = true
		}
	}

	if !shouldPopulate {
		c.JSON(http.StatusOK, chefs)
		return
	}

	populatedChefs := make([]ExtendedChef, len(chefs))
	// get all recipes of chef
	for i, c := range chefs {
		chefRecipes := []Recipe{}
		for _, r := range recipes {
			if r.Chef.Id == c.Id {
				chefRecipes = append(chefRecipes, r)
			}
		}

		populatedChefs[i] = ExtendedChef{
			Chef:    c,
			Recipes: chefRecipes,
		}
	}

	c.JSON(http.StatusOK, populatedChefs)
}

func UpdateChefHandler(c *gin.Context) {
	var chef Chef
	id := c.Param("chef-id")

	if err := c.ShouldBindJSON(&chef); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	index := -1
	for i, c := range chefs {
		if c.Id == id {
			index = i
		}
	}

	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Chef not found",
		})
		return
	}

	chef.Id = id
	chefs[index] = chef

	c.JSON(http.StatusOK, chef)
}

func DeleteChefHandler(c *gin.Context) {
	id := c.Param("chef-id")

	index := -1
	for i, c := range chefs {
		if c.Id == id {
			index = i
		}
	}

	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Chef not found",
		})
		return
	}

	if len(chefs) <= 1 {
		chefs = make([]Chef, 0)
	} else {
		chefs = append(chefs[:index], chefs[index+1:]...)
	}

	delete(chefsIndex, id) // delete chef from index
	c.JSON(http.StatusOK, gin.H{
		"message": "Chef deleted",
	})
}
