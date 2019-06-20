package main

import "fmt"

func main() {
	myMap := map[string][]string{
		`bond_james`: []string {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`: []string{`Being evil`, `Ice cream`, `Sunsets`}}
	myMap["avigail"] = []string{`aa`, `bb`, `cc`}
	delete(myMap, `no_dr`)
	for i := range myMap {
		fmt.Println(i, myMap[i])
		for j, v := range myMap[i] {
			fmt.Println(j, v)
		}

	}

}
