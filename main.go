package main

import (
	// sesuaikan dengan struktuk folder project masing2

	"fmt"
	model "go_proto/model"
	"os"

	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	var user1 = &model.User{
		Id:       "u001",
		Name:     "Raden Ario Damar",
		Password: "12345678",
		Gender:   model.UserGender_FEMALE,
	}

	var userList = &model.UserList{
		List: []*model.User{
			user1,
		},
	}

	var garage1 = &model.Garage{
		Id:   "g001",
		Name: "Sette Desert",
		Coordinate: &model.GarageCoordinate{
			Latitude:  23.2212847,
			Longitude: 53.22033123,
		},
	}

	var garageList = &model.GarageList{
		List: []*model.Garage{
			garage1,
		},
	}

	var garageListByUser = &model.GarageListByUser{
		List: map[string]*model.GarageList{
			user1.Id: garageList,
		},
	}

	fmt.Printf("# ====== original\n %#v\n\n", user1)
	fmt.Printf("# ====== original\n %#v\n\n", userList)
	fmt.Printf("# ====== original\n %#v\n\n", garage1)
	fmt.Printf("# ====== original\n %#v\n\n", garageList)
	fmt.Printf("# ====== original\n %#v\n\n", garageListByUser)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Printf("# ====== As String\n %#v\n\n", user1.String())
	fmt.Printf("# ====== As String\n %#v\n\n", userList.String())
	fmt.Printf("# ====== As String\n %#v\n\n", garage1.String())
	fmt.Printf("# ====== As String\n %#v\n\n", garageList.String())
	fmt.Printf("# ====== As String\n %#v\n\n", garageListByUser.String())

	// convert proto object to json string
	jsonBytes, err1 := protojson.Marshal(user1)
	if err1 != nil {
		fmt.Println(err1.Error())
		os.Exit(0)
	}
	jsonString := string(jsonBytes)
	fmt.Printf("# ==== As Proto Object\n%v", jsonString)

	// convert json string to proto object
	err2 := protojson.Unmarshal([]byte(jsonString), user1)
	if err2 != nil {
		fmt.Println(err2.Error())
		os.Exit(0)
	}

	fmt.Printf("# ====== As String\n %#v\n\n", garageList)
}
