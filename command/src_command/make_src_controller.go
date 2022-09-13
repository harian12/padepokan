package src_command

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func MakeSrcController(arg, app_name_cmd string) {

	arg = strings.Title(arg)

	/*cek apakah file sudah ada*/
	if _, err := os.Stat("app/controllers/" + strings.ToLower(arg) + "_controller.go"); !errors.Is(err, os.ErrNotExist) {
		panic("controllers sudah ada")
	}

	f, err := os.OpenFile("app/controllers/"+strings.ToLower(arg)+"_controller.go", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	if _, err = f.WriteString(
		"package controllers \n" +
			"\n" +
			"import (\n" +
			"\"" + app_name_cmd + "/app/services\"\n" +
			")\n\n" +
			"type I" + arg + "Controller interface{\n" +
			"\n}\n\n" +
			"type " + strings.ToLower(arg) + "Controller struct {\n" +
			strings.ToLower(arg) + "Service services.I" + arg + "Service" +
			"\n}\n\n" +
			"func New" + arg + "Controller(" + strings.ToLower(arg) + "Service services.I" + arg + "Service) *" + strings.ToLower(arg) + "Controller {\n" +
			"	return &" + strings.ToLower(arg) + "Controller{" + strings.ToLower(arg) + "Service: " + strings.ToLower(arg) + "Service}" +
			"\n}\n\n"); err != nil {
		panic(err)
	}
	fmt.Println("create controllers success")

}
