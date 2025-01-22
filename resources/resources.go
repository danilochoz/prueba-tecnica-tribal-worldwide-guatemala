package resources

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"pt-twwgt/models"
	"slices"
)

func RandomUser(w http.ResponseWriter, r *http.Request) {
	config := models.Config{
		URL: "https://randomuser.me/api/?results=5000",
	}
	users := []models.User{}
	for {
		response, err := http.Get(config.URL)
		if err != nil {
			log.Fatal(err)
		}
		responseData, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		result := models.Result{}
		err2 := json.Unmarshal(responseData, &result)
		if err2 != nil {
			log.Fatal(err2)
		}
		for _, user := range result.Results {
			exists := slices.ContainsFunc(users, func(usr models.User) bool {
				return usr.Login.UUID == user.Login.UUID
			})
			if !exists {
				users = append(users, user)
			}
		}
		if len(users) > 15000 {
			break
		}
	}
	result := models.Result{
		Results: users,
	}
	json.NewEncoder(w).Encode(result)
}
