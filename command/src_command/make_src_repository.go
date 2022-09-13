package src_command

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func MakeSrcRepository(arg string) {
	arg = strings.Title(arg)

	/*cek apakah file sudah ada*/
	if _, err := os.Stat("app/repositories/" + strings.ToLower(arg) + "_repo.go"); !errors.Is(err, os.ErrNotExist) {
		panic("repositories sudah ada")
	}

	f, err := os.OpenFile("app/repositories/"+strings.ToLower(arg)+"_repo.go", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(
		"package repositories \n" +
			"\n" +
			"import (\n" +
			"	\"gorm.io/gorm\"\n" +
			")\n\n" +
			"type I" + arg + "Repository interface{\n" +
			"\n}\n\n" +
			"type " + strings.ToLower(arg) + "Repository struct {\n" +
			"	db *gorm.DB" +
			"\n}\n\n" +
			"func New" + arg + "Repository(db *gorm.DB) *" + strings.ToLower(arg) + "Repository {\n" +
			"	return &" + strings.ToLower(arg) + "Repository{db}" +
			"\n}\n\n"); err != nil {
		panic(err)
	}

	fmt.Println("create repositories success")
}
