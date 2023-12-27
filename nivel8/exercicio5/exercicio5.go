/*- Partindo do código abaixo, ordene os []user por idade e sobrenome.
    - https://play.golang.org/p/BVRZTdlUZ_
- Os valores no campo Sayings devem ser ordenados tambem, e demonstrados de maneira esteticamente harmoniosa.
- Solução: https://play.golang.org/p/3wgW4BDasu */

package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type UserSortByAge []user

func (u UserSortByAge) Len() int           { return len(u) }
func (u UserSortByAge) Swap(i, j int)      { u[i], u[j] = u[j], u[i] }
func (u UserSortByAge) Less(i, j int) bool { return u[i].Age < u[j].Age }

type UserSortByLast []user

func (u UserSortByLast) Len() int           { return len(u) }
func (u UserSortByLast) Swap(i, j int)      { u[i], u[j] = u[j], u[i] }
func (u UserSortByLast) Less(i, j int) bool { return u[i].Last < u[j].Last }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	// your code goes here

	//Ordenação por idade
	var userSortByAge UserSortByAge = users

	sort.Sort(userSortByAge)

	fmt.Println(userSortByAge)

	//Ordenação por sobrenome
	var userSortByLast UserSortByLast = users

	sort.Sort(userSortByLast)

	fmt.Println(userSortByLast)

}
