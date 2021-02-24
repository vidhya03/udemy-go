package main

import "fmt"

func main() {

	nattukozhi := map[string]string{
		"aki": "1days",
		"moi": "2days",
	}

	nattukozhi["kum"] = "4days"
	nattukozhi["zari"] = "6days"
	delete(nattukozhi, "aki")

	//this statement will get ignored
	delete(nattukozhi, "akisfs")

	// fmt.Println(nattukozhi)
	printMap(nattukozhi)
}

func printMap(c map[string]string) {
	for kozhi, days := range c {
		fmt.Println(kozhi, " is ", days)
	}
}
