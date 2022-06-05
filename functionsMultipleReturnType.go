package main

func multipleReturn(username string)(string, int){
	userInfo := make(map[string]int)
	userInfo["swopnil"] = 1
	userInfo["randomGuy"] = 2

	for key, value := range userInfo{
		if key == username{
			return key, value
		}
	}
	return "userNotFound", 0
}