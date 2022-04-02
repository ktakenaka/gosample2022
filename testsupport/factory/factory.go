package factory

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/ktakenaka/gosample2022/app/domain/models"
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
	office := &models.Office{
		ID:   ulid.MustNew(),
		Name: "name",
	}

	if err := attrs.overWrite(reflect.ValueOf(office)); err != nil {
		panic(err)
	}

	return office
}

func MustBuildUser(attrs Attributes) *models.User {
	user := &models.User{
		ID:    ulid.MustNew(),
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

	ou := &models.OfficeUser{
		ID:       ulid.MustNew(),
		UserID:   user.ID,
		OfficeID: office.ID,
	}

	if err := attrs.overWrite(reflect.ValueOf(user)); err != nil {
		panic(err)
	}

	return user, office, ou
}
