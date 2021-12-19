package factory

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/ktakenaka/gosample2022/app/models"
	"github.com/ktakenaka/gosample2022/app/pkg/ulid"
)

var ErrInvalidAttribute = errors.New("invalid attribute")

// Attributes to overwrite default attributes
type Attributes map[string]interface{}

func (attrs Attributes) overWrite(ptr reflect.Value) error {
	for key, value := range attrs {
		v := ptr.Elem()
		f := v.FieldByName(key)
		if !f.CanSet() {
			return fmt.Errorf("%w: %s.%s is not writable or does not exist", ErrInvalidAttribute, v.Type().Name(), key)
		}

		newVal := reflect.ValueOf(value)
		if f.Type().Name() != newVal.Type().Name() {
			return fmt.Errorf("%w: %s.%s wants %s but %s is passed", ErrInvalidAttribute, v.Type().Name(), key, f.Type().Name(), newVal.Type().Name())
		}
		f.Set(newVal)
	}
	return nil
}

func MustBuildOffice(attrs Attributes) *models.Office {
	id, err := ulid.GenerateID()
	if err != nil {
		panic(err)
	}

	office := &models.Office{
		ID:   id,
		Name: "name",
	}

	if err := attrs.overWrite(reflect.ValueOf(office)); err != nil {
		panic(err)
	}

	return office
}

func MustBuildUser(attrs Attributes) *models.User {
	id, err := ulid.GenerateID()
	if err != nil {
		panic(err)
	}

	user := &models.User{
		ID:    id,
		Email: "example@dummy.com",
	}

	if err := attrs.overWrite(reflect.ValueOf(user)); err != nil {
		panic(err)
	}

	return user
}

func MustBuildOfficeUser(attrs Attributes) (
	*models.User, *models.Office, *models.OfficeUser,
) {
	user := MustBuildUser(nil)
	office := MustBuildOffice(nil)

	id, err := ulid.GenerateID()
	if err != nil {
		panic(err)
	}

	ou := &models.OfficeUser{
		ID:       id,
		UserID:   user.ID,
		OfficeID: office.ID,
	}

	if err := attrs.overWrite(reflect.ValueOf(user)); err != nil {
		panic(err)
	}

	return user, office, ou
}
