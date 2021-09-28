package controller

import (
	"os"
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"tes_kitalulus/model"
	"github.com/gin-gonic/gin"
)

func ListUser(c *gin.Context){
	// get respon data from api
	response, err := http.Get("https://api.github.com/users")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// convert data from jsonstring to json object
	var userObj []model.ListUser
	err = json.Unmarshal(responseData, &userObj)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// dump data to temporary for become json string
	// dataJsonString := []model.ListUser{}
	// for i := 0; i < len(userObj); i++ {
	// 	dataJsonString = append(dataJsonString, model.ListUser{
	// 		Login : userObj[i].Login, 
	// 		Id : userObj[i].Id, 
	// 		Node_id : userObj[i].Node_id, 
	// 		Avatar_url : userObj[i].Avatar_url, 
	// 		Gravatar_id : userObj[i].Gravatar_id, 
	// 		Url : userObj[i].Url, 
	// 		Html_url : userObj[i].Html_url, 
	// 		Followers_url : userObj[i].Followers_url, 
	// 		Following_url : userObj[i].Following_url, 
	// 		Gists_url : userObj[i].Gists_url, 
	// 		Starred_url : userObj[i].Starred_url, 
	// 		Subscriptions_url : userObj[i].Subscriptions_url, 
	// 		Organizations_url : userObj[i].Organizations_url, 
	// 		Repos_url : userObj[i].Repos_url, Events_url : 
	// 		userObj[i].Events_url, 
	// 		Received_events_url : userObj[i].Received_events_url, 
	// 		Type : userObj[i].Type, Site_admin : userObj[i].Site_admin,
	// 	})
	// }
	// jsonData, _ := json.Marshal(dataJsonString)

	// // convert json string to json object
	// var data []model.ListUser

	// err = json.Unmarshal(jsonData, &data)
	// if err != nil {
	// 		fmt.Println(err.Error())
	// 		return
	// }

	//return response
	c.JSON(200, userObj)
}

func DetailUser(c *gin.Context){
	// get param login user
	login := c.Param("login")

	// get respon data from api
	response, err := http.Get("https://api.github.com/users/"+login)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// convert data from jsonstring to json object
	var userObj model.DetailUser
	err = json.Unmarshal(responseData, &userObj)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// direct to not found if user not found
	if userObj.Login == "" {
		var userObj model.NotFoundUser
		json.Unmarshal(responseData, &userObj)
		c.JSON(200, userObj)
		return
	}
	
	c.JSON(200, userObj)
}
