package main

import (
	"context"
	"fmt"

	pb "dodi/genproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var names = []string{
	"Abbos", "Azizbek", "Bekzod", "Diyorbek", "Faxriddin", "Fazliddin", 
	"Hamidjon", "Hamidulloh", "Ibrohim", "Jamshidbek", "Javohir", 
	"Muhammadaziz", "Muhammadjon", "Nurmuhammad", "Ozodjon", "Sanjarbek", 
	"Bobur", "Firdavs", "Ozodbek", "Abdulaziz", "Intizor",
}
var nameToFamily = map[string]string{
	"Abbos":                      "Qambarov",
	"Azizbek":                    "Qobulov",
	"Bekzod":                     "Qo'chqarov",
	"Diyorbek":                   "Nematov Dadajon o'g'li",
	"Faxriddin":                  "Raximberdiyev Farxodjon o'g'li",
	"Fazliddin":                  "Xayrullayev",
	"Hamidjon":                   "Nuriddinov",
	"Hamidulloh":                 "Hamidullayev",
	"Ibrohim":                    "Umarov Fazliddin o'g'li",
	"Jamshidbek":                 "Hatamov Erkin o'g'li",
	"Javohir":                    "Abdusamatov",
	"Muhammadaziz":               "Yoqubov",
	"Muhammadjon":                "Ko'palov",
	"Nurmuhammad":                "",
	"Ozodjon":                    "A'zamjonov",
	"Sanjarbek":                  "Abduraxmonov",
	"Bobur":                      "Yusupov",
	"Firdavs":                    "",
	"Ozodbek":                    "",
	"Abdulaziz":                  "Xoliqulov",
	"Intizor":                    "opa",
}

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	gen := pb.NewGeneratorClient(conn)
	req := &pb.Request{
		Limit:     1,
		Names:     names,
	}
	resp, err := gen.RandomPicker(context.Background(), req)
	if err != nil {
		panic(err)
	}
	fmt.Println("Seat - Name")
	for k, v := range resp.Result {
		for d, _ := range nameToFamily {
			if v == {
				fmt.Println(v)
			}
		}
	}
}