package cmd_variabel

import (
	"fmt"
	"os"
	"strings"
)

func AddSrcVariabel(arg string) {
	arg = strings.Title(arg)

	f, err := os.OpenFile("router/declaration.go", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	if _, err = f.WriteString(
		"\nvar " + strings.ToLower(arg) + "_repository repositories.I" + arg + "Repository = repositories.New" + arg + "Repository(db)\n" +
			"var " + strings.ToLower(arg) + "_service services.I" + arg + "Service = services.New" + arg + "Service(" + strings.ToLower(arg) + "_repository)\n" +
			"var " + strings.ToLower(arg) + "_controller controllers.I" + arg + "Controller = controllers.New" + arg + "Controller(" + strings.ToLower(arg) + "_service)\n"); err != nil {
		panic(err)
	}
	fmt.Println("Add Variable success")
}

func AddControllerVariabel(arg string) {
	arg = strings.Title(arg)

	f, err := os.OpenFile("router/declaration.go", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	if _, err = f.WriteString(
		"var " + strings.ToLower(arg) + "_controller controllers.I" + arg + "Controller = controllers.New" + arg + "Controller(" + strings.ToLower(arg) + "_service)\n"); err != nil {
		panic(err)
	}
	fmt.Println("Add Variable success")
}
