package main

type Users struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Age int64 `json:"age"`
}

func main () {
	// read json file 

	jsonFile, err := os.Open("test3.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	jsonByte, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}

	var user Users
	json.Unmarshal(jsonByte, &user)

	user.Email = "johndoe@example.com"
	user.Age++

	userJson, _ := json.Marshal(user)
	err = ioutil.WriteFile("test3.json", userJson, 0644)


}