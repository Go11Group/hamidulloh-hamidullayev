package main

import (
	"Modul/funcgo"
	"Modul/modul"
	"fmt"
)

func main() {

	db, err := funcgo.ConnectDB()
	if err != nil {
		panic(err)
	}
	dbUser := funcgo.NewUserRepo(db)
	dbproducts := funcgo.NewProductsRepo(db)

	var a int

	for true {
		fmt.Println("1 => User ")
		fmt.Println("2 => Products ")
		fmt.Println("1 yokida 2 ")
		fmt.Scan(&a)
		if a == 1 || a == 2 {
			break
		}
	}
	switch a {
	case 1:
		fmt.Println("User va siz nima qilasiz ")
		fmt.Println("1 => Create => User qoshish")
		fmt.Println("2 => Delete => User ochirish")
		fmt.Println("3 => Update => User ni ozgartirish")
		fmt.Println("4 => Select => Userlarni korsatish")
		fmt.Println("5 => Qidirish")
		var a1 int
		for true {
			fmt.Scan(&a1)
			if a1 == 1 || a1 == 2 || a1 == 3 || a1 == 4 || a1 == 5 {
				break
			}
		}
		switch a1 {
		case 1:
			var name string
			var email string
			var password string
			fmt.Println("UserName kirirting:")
			fmt.Scan(&name)
			fmt.Println("Email kirirting:")
			fmt.Scan(&email)
			fmt.Println("Password kirirting:")
			fmt.Scan(&password)
			user := modul.User{
				UserName: name,
				Email:    email,
				Password: password,
			}
			dbUser.UserCreate(user)
			fmt.Println("Mofaqiyatli qoshildi")
		case 2:
			var id uint
			fmt.Println("User id kirirting ochirish uchun:")
			fmt.Scan(&id)
			dbUser.UserDelete(id)
			fmt.Println("Mofaqiyatli Ochirildi")
		case 3:
			var name string
			var email string
			var password string
			fmt.Println("UserName kirirting ozgartiradigan:")
			fmt.Scan(&name)
			fmt.Println("Email kirirting ozgartiradigan:")
			fmt.Scan(&email)
			fmt.Println("Password kirirting ozgartiradigan:")
			fmt.Scan(&password)
			user := modul.User{
				UserName: name,
				Email:    email,
				Password: password,
			}
			var id uint
			fmt.Println("User id kirirting ozgartiradigan userni :")
			fmt.Scan(&id)
			dbUser.UserUpdate(id,user)
			fmt.Println("Mofaqiyatli ozgartirildi")
		case 4:
			c, err := dbUser.UserSelect()
			if err != nil{
				panic(err)
			}
			fmt.Println(c)
		case 5:
			var id uint
			fmt.Println("User id kirirting qidiryatgan:")
			fmt.Scan(&id)
			user, err := dbUser.UserID(id)
			if err != nil{
				panic(err)
			}
			fmt.Println(user)
		}

	case 2:
		fmt.Println("Products va siz nima qilasiz ")
		fmt.Println("1 => Create => Products qoshish")
		fmt.Println("2 => Delete => Products ochirish")
		fmt.Println("3 => Update => Products ni ozgartirish")
		fmt.Println("4 => Select => Productlarni korsatish")
		fmt.Println("5 => Qidirish")
		var a1 int
		for true {
			fmt.Scan(&a1)
			if a1 == 1 || a1 == 2 || a1 == 3 || a1 == 4 || a1 == 5 {
				break
			}
		}
		
		switch a1 {
		case 1:
			var name string
			var description string
			var price float32
			var stock_quantity int 
			fmt.Println("Name kirirting:")
			fmt.Scan(&name)
			fmt.Println("description kirirting:")
			fmt.Scan(&description)
			fmt.Println("price kirirting:")
			fmt.Scan(&price)
			fmt.Println("stock_quantity kirirting:")
			fmt.Scan(&stock_quantity)
			produc := modul.Products{
				Name: name,
				Description: description,
				Price: price,
				Stock_quantity: stock_quantity,
			}
			dbproducts.ProductsCreate(produc)
			fmt.Println("Mofaqiyatli qoshildi")
		case 2:
			var id uint
			fmt.Println("Product id kirirting ochirish uchun:")
			fmt.Scan(&id)
			dbproducts.ProductsDelete(id)
			fmt.Println("Mofaqiyatli Ochirildi")
		case 3:
			var name string
			var description string
			var price float32
			var stock_quantity int 
			fmt.Println("Name kirirting:")
			fmt.Scan(&name)
			fmt.Println("description kirirting:")
			fmt.Scan(&description)
			fmt.Println("price kirirting:")
			fmt.Scan(&price)
			fmt.Println("stock_quantity kirirting:")
			fmt.Scan(&stock_quantity)
			produc := modul.Products{
				Name: name,
				Description: description,
				Price: price,
				Stock_quantity: stock_quantity,
			}
			var id uint
			fmt.Println("Product id kirirting ozgartiradigan Productni :")
			fmt.Scan(&id)
			dbproducts.ProductsUpdate(id,produc)
			fmt.Println("Mofaqiyatli ozgartirildi")
		case 4:
			c, err := dbproducts.ProductsSelect()
			if err != nil{
				panic(err)
			}
			fmt.Println(c)
		case 5:
			var id uint
			fmt.Println("Product id kirirting qidiryatgan:")
			fmt.Scan(&id)
			user, err := dbproducts.ProductsID(id)
			if err != nil{
				panic(err)
			}
			fmt.Println(user)
		}
	}
}
