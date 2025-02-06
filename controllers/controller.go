package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"num-api/utils"
	
)


// Function to fetch a fun fact from the Numbers API
func getFunFact(n int) (string, error) {
	url := fmt.Sprintf("http://numbersapi.com/%d?json", n)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}

	return result["text"].(string), nil
}

func ClassifyNumber(c *gin.Context) {
	// Get the number parameter
	numberStr := c.DefaultQuery("number", "")
	if numberStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
		})
		return
	}

	// Try to convert the number to an integer
	number, err := strconv.Atoi(numberStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"number":"alphabet",
			"error": true,
		})
		return
	}

	// Generate the classification properties
	properties := []string{}
	if utils.IsArmstrong(number) {
		properties = append(properties, "armstrong")
	}
	if number%2 == 0 {
		properties = append(properties, "even")
	} else {
		properties = append(properties, "odd")
	}

	// Get the fun fact from the Numbers API
	funFact, err := getFunFact(number)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
		})
		return
	}

	// Build the JSON response
	response := gin.H{
		"number":    number,
		"is_prime":  utils.IsPrime(number),
		"is_perfect": utils.IsPerfect(number),
		"properties": properties,
		"digit_sum": utils.DigitSum(number),
		"fun_fact":  funFact,
	}

	// Return the response
	c.JSON(http.StatusOK, response)
}