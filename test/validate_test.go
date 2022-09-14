package test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"padepokan/helpers/validation"
	"testing"
)

type dtoForTest struct {
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password,omitempty" validate:"required,max=20,min=6"`
	ConfirmPassword string `json:"confirm_password,omitempty" validate:"required,max=20,min=6,eqfield=Password"`
}

func TestValidationNegativCaseForDTO(t *testing.T) {
	objRegisterNegativeCase := dtoForTest{
		Email:           "hariansyah434@gmail.com",
		Password:        "admin1234",
		ConfirmPassword: "admin123",
	}

	objRegisterPositifCase := dtoForTest{
		Email:           "hariansyah434@gmail.com",
		Password:        "admin1234",
		ConfirmPassword: "admin1234",
	}
	t.Run("negetive case", func(t *testing.T) {
		getErrMesg, errDTO := validation.ValidationForDTO(objRegisterNegativeCase)
		if errDTO == nil {
			data, _ := json.Marshal(getErrMesg)
			assert.Equal(t, string(data), "not success")
			t.FailNow()
		}

	})

	t.Run("positif case", func(t *testing.T) {
		getErrMesg, errDTO := validation.ValidationForDTO(objRegisterPositifCase)
		if errDTO != nil {
			data, _ := json.Marshal(getErrMesg)
			assert.Equal(t, string(data), "success")
			t.FailNow()
		}

	})

}
