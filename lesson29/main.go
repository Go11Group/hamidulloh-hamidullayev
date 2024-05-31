package main

import (
	"Modul/funcgo"
	// "Modul/modul"
	"fmt"
	// "os/user"

	// "gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	// "gorm.io/gorm"
	_ "gorm.io/gorm"
)

func main() {

	dbUser, err := funcgo.NewDBUser()
	if err != nil {
		panic(err)
	}

	//bunda foydalanuchi qoshish ===============================================================
	// user := modul.User{FirstName: "Zaaaza", LastName: "Xamidulloyev1",Email: "nssd",Password: "1111",Age: 23,Fiald: "G011",Gander: "m",IsEmployee: true}
	// dbUser.Create1(user)
	// fmt.Println("User created")

	//bunda berilgan userni malumotlar kiritib beroilgan aydi boyicha ozgartirib bertadi============================================================================
	// user := modul.User{FirstName: "Zaa", LastName: "Xamidullev1",Email: "nssd",Password: "1111",Age: 23,Fiald: "G011",Gander: "m",IsEmployee: true}
	// var a1 string
	// fmt.Println("Ochirish uchun foydalanuchi aydisini kiriting ")
	// fmt.Scan(&a1)
	// if err := dbUser.Update1(user,a1); err != nil {
	// 	panic(err)
	// }
	// fmt.Println("User updated")

	var ism string
	var son int
	fmt.Println("1 id boyicha qidirish\n2 Firstname boyicha qidirish\n3 email boyicha qidirish\n4 lastname boyicha qidirish\nson kiritin 1 4 oraligida")
	fmt.Scan(&son)
	fmt.Println("qidiryabgan narsani kiritng")
	fmt.Scan(&ism)
	if son > 0 && son < 5 {
		use, _ := dbUser.FindByID(ism, son)
		fmt.Println(use)
	}


	//bunda berilhgan aydini ochirib tashl,aydi=========================================================
	// var a uint
	// fmt.Println("Ochirish uchun foydalanuchi aydisini kiriting ")
	// fmt.Scan(&a)
	// dbUser.Delete1(a)
	// fmt.Println("User deleted")
}
