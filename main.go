package main

import (
    "net/http"
    "os"

    "github.com/gin-gonic/gin"
    "fmt"
)

type getPacksRequest struct {
    NumberOfItems int `json:"numberOfItems"`
    PackSizes  []int  `json:"packSizes"`
}

func main() {
    port := os.Getenv("PORT")
    host := ""

    if port == "" {
        port = "8080";
        host = "localhost"
    }

    router := gin.Default()
    
    router.POST("/pack-sizes", setPackSizesHandler)
    router.GET("/pack-sizes", getPackSizesHandler)
    router.GET("/packs-old", getPacksOldHandler)
    router.GET("/packs/all-in-one", getPacksAllInOneHandler)

    // This is the main entry point 
    router.POST("/packs", getPacksHandler)

    router.Run(host + ":" + port)
}

func setPackSizesHandler(c *gin.Context) {
    var request []int

    if err := c.BindJSON(&request); err != nil {
        fmt.Println("Error parsing request.")
        c.Writer.WriteHeader(http.StatusBadRequest)
        return
    }

    if len(request) == 0 {
        fmt.Println("Requires at least 1 pack size to be passed.")
        c.Writer.WriteHeader(http.StatusBadRequest)
        return
    }

    var updatedPackSizes = setPackSizes(request)

	c.IndentedJSON(http.StatusOK, updatedPackSizes)
}

func getPackSizesHandler(c *gin.Context) {
    var packSizes = getPackSizes()
    c.IndentedJSON(http.StatusOK, packSizes)
}

func getPacksOldHandler(c *gin.Context) {
    var request int

    if err := c.BindJSON(&request); err != nil {
        fmt.Println("Error parsing request.")
        c.Writer.WriteHeader(http.StatusBadRequest)
        return
    }

    if  request <= 0  {
        fmt.Println("Invalid number of items in request.")
        c.Writer.WriteHeader(http.StatusBadRequest)
        return
    }

    if len(getPackSizes()) == 0 {
        fmt.Println("Setting pack sizes required before calculating packs for order.")
        c.Writer.WriteHeader(http.StatusBadRequest)
        return
    }

    var packs = getPacksOld(request)

	c.IndentedJSON(http.StatusOK, packs)
}

func getPacksAllInOneHandler(c *gin.Context) {
    var request getPacksRequest

    if err := c.BindJSON(&request); err != nil {
        fmt.Println("Error parsing request.")
        c.Writer.WriteHeader(http.StatusBadRequest)
        return
    }

    if  request.NumberOfItems <= 0  {
        fmt.Println("Invalid number of items in request.")
        c.Writer.WriteHeader(http.StatusBadRequest)
        return
    }

    if len(request.PackSizes) < 2 {
        fmt.Println("Invalid pack sizes in request.")
        c.Writer.WriteHeader(http.StatusBadRequest)
        return
    }

    var packs = getPacksFromSizes(request.NumberOfItems, request.PackSizes)

	c.IndentedJSON(http.StatusOK, packs)
}

func getPacksHandler(c *gin.Context) {
    c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
    var request getPacksRequest

    if err := c.BindJSON(&request); err != nil {
        fmt.Println("Error parsing request.")
        c.Writer.WriteHeader(http.StatusBadRequest)
        return
    }

    if  request.NumberOfItems <= 0  {
        fmt.Println("Invalid number of items in request.")
        c.Writer.WriteHeader(http.StatusBadRequest)
        return
    }

    if len(request.PackSizes) < 2  {
        fmt.Println("Invalid pack sizes in request.")
        c.Writer.WriteHeader(http.StatusBadRequest)
        return
    }

    for _, packSize := range request.PackSizes {
		if !(packSize > 0) {
			fmt.Println("Invalid pack sizes in request.")
            c.Writer.WriteHeader(http.StatusBadRequest)
            return
		}
	}

    var packs = getPacks(request.NumberOfItems, request.PackSizes)

	c.IndentedJSON(http.StatusOK, packs)
}