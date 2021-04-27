package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Vehicle struct {
	Id    string `json:"id" binding:"required"`
	Model string `json:"model" binding:"required"`
	Make  string `json:"make" binding:"required"`
	Year  int    `json:"year"`
}

var length = 0 //to start at least 0 length and increase overtime
var storeData = make([]Vehicle, length)

func main() {

	router := gin.Default()

	//GET '/'  --> all cars
	router.GET("/cars", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":  "All Cars",
			"vehicles": storeData,
		})
	})

	//POST '/cars'  --> create cars
	router.POST("/cars", func(c *gin.Context) {
		var car Vehicle

		err := c.ShouldBindJSON(&car) //binds the input data into 'car' var
		if err != nil {

			fmt.Println(err)
			c.JSON(422, gin.H{"message": "Car Not Created"})
			return
		}

		storeData = append(storeData, car)

		c.JSON(200, gin.H{
			"message": "Successful post",
			"id":      car.Id,    //write in 'id' --> `json:"id"`
			"make":    car.Make,  //write in 'title' --> `json:"title" binding:"required"`
			"model":   car.Model, //write in 'type' --> `json:"type" binding:"required"`
			"year":    car.Year,  //write in 'madeIn' --> `json:"madeIn"`
		})
	})

	//GET '/cars/:carid'  --> get single car
	router.GET("/cars/:carid", func(c *gin.Context) {
		carid := c.Param("carid")
		var car Vehicle

		for i := 0; i < len(storeData); i++ {
			//at the end and id not found
			if i == len(storeData)-1 && storeData[i].Id != carid {
				c.JSON(404, gin.H{
					"message": "car id not found",
				})
				return
			}
			if storeData[i].Id == carid {
				car = storeData[i]
			}
		}
		c.JSON(200, gin.H{
			"message": "Single Car Found",
			"car":     car,
		})
	})

	//PUT '/cars/:carid'  --> modify that single car
	router.PUT("/cars/:carid", func(c *gin.Context) {

		var car Vehicle
		err := c.ShouldBindJSON(&car) //binds the input data into 'car' var
		if err != nil {
			c.JSON(400, gin.H{"message": "Could not update Car"})
			return
		}

		for i := 0; i < len(storeData); i++ {
			if storeData[i].Id != car.Id && err == nil {
				c.JSON(404, gin.H{
					"message": "car not found",
				})
				return
			} else {
				storeData[i] = car
				c.JSON(200, gin.H{
					"car": storeData[i],
				})
			}
		}

	})

	//DELETE '/cars/:carid'  --> delete that single car
	router.DELETE("/cars/:carid", func(c *gin.Context) {
		carid := c.Param("carid")

		for i := 0; i < len(storeData); i++ {

			//at the end and id not found
			if i == len(storeData)-1 && storeData[i].Id != carid {
				c.JSON(404, gin.H{
					"message": "car not found",
				})
				return
			}
			if storeData[i].Id == carid {
				storeData = append(storeData[:i], storeData[i+1:]...)
			}
		}
		c.JSON(200, gin.H{
			"message":  "Deleted",
			"vehicles": storeData,
		})
	})

	router.Run()
}
