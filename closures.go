package main

func closure() func() string{
	return func () string {
		return "Swopnil"
	}
}