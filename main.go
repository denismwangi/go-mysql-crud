package main

import (
	"go-crud/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

type APIResponse struct {
    Code    string `json:"code"`
    Message string `json:"message"`
    Data    interface{} `json:"data"`
}
type IndexAPIResponse struct {
    Code    string `json:"code"`
    Message string `json:"message"`
    Data    []string `json:"data"`
}
func main() {
	router := gin.Default()
	router.GET("/",testAPI)
	router.GET("/movies", listMoviesHandler)
	router.POST("/movies", createMovieHandler)
	router.GET("/movies/:id", getMovieById)
	router.Run("localhost:8005")
}

func listMoviesHandler(context *gin.Context) {
    movieItems := models.ListMoviesHandler()

    response := APIResponse{
        Code:    "404",
        Message: "No movies found",
        Data:    nil, 
	}
    if movieItems == nil || len(movieItems) == 0 {
        context.JSON(http.StatusNotFound, response)
    } else {
        response.Code = "200"
        response.Message = "Movies found"
        response.Data = movieItems

        context.JSON(http.StatusOK, response)
    }
}


func createMovieHandler(context *gin.Context){
	var movieItem models.Movie

	if err := context.BindJSON(&movieItem); err != nil {
		context.AbortWithStatus(http.StatusBadRequest)
	}else{
		models.CreateMovieHandler(movieItem)
		context.IndentedJSON(http.StatusOK, movieItem)
	}
}

func getMovieById(context *gin.Context){
	id := context.Param("id")

	response := APIResponse{
        Code:    "404",
        Message: "No movies found",
        Data:    nil, 
	}

	movieItem := models.GetMovieById(id)
	if movieItem == nil {
		context.AbortWithStatus(http.StatusNotFound,response)
	}else{
		response.Code = "200"
        response.Message = "Movie found"
        response.Data = movieItem
		context.IndentedJSON(http.StatusOK,response)
	}
}

func testAPI(context *gin.Context){
	response := IndexAPIResponse{
        Code:    "200",
        Message: "API running...",
		Data: []string{},
    }

	context.JSON(http.StatusOK, response)
}