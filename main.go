package main

import (
	"github.com/SuperCodeM/course"
)

func main() {
	Go.course.new("Go desde cero", 0, false)
	
	Go.UserIDs = []uint{15, 45, 77},
	Go.Name = map[uint]string{
			1: "introduccion",
			2: "Estructuras",
			3: "Mapas",
		},
	}
	//Go.printclases()
	//(Go).Changeprice(67.2)
	Go.Printclases()
	Go.Changeprice()
	//fmt.Println("El nuevo precio del curso: "+Go.Name,Go.Price)
}
