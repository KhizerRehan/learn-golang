package user

import (
	"fmt"
	"learn-golang/go-with-maximilian/structs"
)

func GetUserFromExportedStruct() {

	user, err := structs.New("Khizer", 20, "khizerrehan92@gmail.com")
	if err != nil {
		fmt.Println("Error:", err)
		return
	} 
	fmt.Println(user)


	// Empty User
	user2, err := structs.New("", 0, "jane@example.com")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(user2)

}


func GetAdminFromExportedStruct() {
	admin, err := structs.NewAdmin("Khizer", 20, "khizerrehan92@gmail.com", "admin")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(admin)

	fmt.Println(admin.GetRole())
	fmt.Println(admin.GetName())
	fmt.Println(admin.GetAge())
	fmt.Println(admin.GetEmail())


	admin2, err := structs.NewAdmin("Khizer", 20, "khizerrehan92@gmail.com", "")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(admin2)
}

func GetStaticAdminFromExportedStruct() {
	staticAdmin := structs.NewAdminStatic()
	fmt.Println(staticAdmin)

	fmt.Println(staticAdmin.GetRole())
	fmt.Println(staticAdmin.GetName())
	fmt.Println(staticAdmin.GetAge())
	fmt.Println(staticAdmin.GetEmail())
}