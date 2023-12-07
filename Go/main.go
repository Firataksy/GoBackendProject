package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/data", getUserData)
	err := http.ListenAndServe(":9000", mux)
	if err != nil {
		panic(err)
	}
}

/*func rediset() {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	jsonData, err := json.Marshal(usersign)
	if err != nil {
		fmt.Println("JSON error:", err)
		return
	}

	rediskey := "userdata"
	err = client.Set(rediskey, jsonData, 0).Err()
	if err != nil {
		fmt.Println("Redis error:", err)
		return
	}
	fmt.Println("Map successful registered redis")

	exists, err := client.Exists(rediskey).Result()
	if err != nil {
		log.Fatal(err)
	}

	if exists == 1 {
		getValue, _ := client.Get(rediskey).Result()
		fmt.Println(getValue)
	}
}
*/
