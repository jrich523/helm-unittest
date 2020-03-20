package validators_test

import (
	"testing"

	"github.com/lrills/helm-unittest/unittest/common"
	. "github.com/lrills/helm-unittest/unittest/validators"
	"github.com/stretchr/testify/assert"
)

var docWithNoElements = ``

var documentWith2Keys = `
a:
  x: 1
  y: 2
  z: 3
`

func TestHasSizeValidatorWhenEmpty(t *testing.T) {
	manifest := makeManifest(docWithNoElements)
	validator := HasSizeValidator{"a", 0}

	pass, msg := validator.Validate(&ValidateContext{
		Docs: []common.K8sManifest{manifest},
	})

	assert.True(t, pass)
	assert.Equal(t, []string{}, msg)
}

func TestHasSizeValidatorWhenCorrectSize(t *testing.T) {
	manifest := makeManifest(documentWith2Keys)
	validator := HasSizeValidator{"a", 3}

	pass, msg := validator.Validate(&ValidateContext{
		Docs: []common.K8sManifest{manifest},
	})

	assert.True(t, pass)
	assert.Equal(t, []string{}, msg)
}

func TestHasSizeValidatorWhenCorrectSizeAndNegative(t *testing.T) {
	manifest := makeManifest(documentWith2Keys)
	validator := HasSizeValidator{"a", 3}

	pass, msg := validator.Validate(&ValidateContext{
		Docs:     []common.K8sManifest{manifest},
		Negative: true,
	})

	assert.False(t, pass)
	assert.Equal(t, []string{
		"Path:	a",
		"Expected size NOT EQUAL:	3, actual:	3",
	}, msg)
}

func TestHasSizeValidatorWhenIncorrectSize(t *testing.T) {
	manifest := makeManifest(documentWith2Keys)
	validator := HasSizeValidator{"a", 1}

	pass, msg := validator.Validate(&ValidateContext{
		Docs: []common.K8sManifest{manifest},
	})

	assert.False(t, pass)
	assert.Equal(t, []string{
		"Path:	a",
		"Expected size:	1, actual:	3",
	}, msg)
}

func TestHasSizeValidatorWhenIncorrectSizeAndNegative(t *testing.T) {
	manifest := makeManifest(documentWith2Keys)
	validator := HasSizeValidator{"a", 2}

	pass, msg := validator.Validate(&ValidateContext{
		Docs:     []common.K8sManifest{manifest},
		Negative: true,
	})

	assert.True(t, pass)
	assert.Equal(t, []string{}, msg)
}
