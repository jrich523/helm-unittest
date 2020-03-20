package validators

import (
	"reflect"
	"strconv"

	"github.com/lrills/helm-unittest/unittest/valueutils"
)

// HasSizeValidator validate value of Path is empty
type HasSizeValidator struct {
	Path  string
	Count int
}

func (v HasSizeValidator) failInfo(actual int, not bool) []string {
	var notAnnotation string
	if not {
		notAnnotation = " NOT EQUAL"
	}

	failFormat := `
Path:%s
Expected size` + notAnnotation + `:%s, actual:%s`
	return splitInfof(failFormat, v.Path, strconv.Itoa(v.Count), strconv.Itoa(actual))
}

// Validate implement Validatable
func (v HasSizeValidator) Validate(context *ValidateContext) (bool, []string) {
	manifest, err := context.getManifest()
	if err != nil {
		return false, splitInfof(errorFormat, err.Error())
	}

	actual, err := valueutils.GetValueOfSetPath(manifest, v.Path)

	if err != nil {
		return false, splitInfof(errorFormat, err.Error())
	}

	actualValue := reflect.ValueOf(actual)

	var size int
	switch actualValue.Kind() {
	case reflect.Invalid:
		size = 0
	case reflect.Array, reflect.Map, reflect.Slice:
		size = actualValue.Len()
	default:
		size = 0
	}
	if (size == v.Count) == !context.Negative {
		return true, []string{}
	}
	return false, v.failInfo(size, context.Negative)
}
