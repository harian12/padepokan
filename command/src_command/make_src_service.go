package src_command

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func MakeSrcService(arg, app_name_cmd string) {

	arg = strings.Title(arg)

	/*cek apakah file sudah ada*/
	if _, err := os.Stat("app/services/" + strings.ToLower(arg) + "_service.go"); !errors.Is(err, os.ErrNotExist) {
		panic("services sudah ada")
	}

	f, err := os.OpenFile("app/services/"+strings.ToLower(arg)+"_service.go", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	if _, err = f.WriteString(
		"package services \n" +
			"\n" +
			"import (\n" +
			"\"" + app_name_cmd + "/app/repositories\"\n" +
			")\n\n" +
			"type I" + arg + "Service interface{\n" +
			"\n}\n\n" +
			"type " + strings.ToLower(arg) + "Service struct {\n" +
			strings.ToLower(arg) + "Repo repositories.I" + arg + "Repository" +
			"\n}\n\n" +
			"func New" + arg + "Service(" + strings.ToLower(arg) + "Repo repositories.I" + arg + "Repository) *" + strings.ToLower(arg) + "Service {\n" +
			"	return &" + strings.ToLower(arg) + "Service{" + strings.ToLower(arg) + "Repo: " + strings.ToLower(arg) + "Repo}" +
			"\n}\n\n"); err != nil {
		panic(err)
	}
	fmt.Println("create services success")

}
