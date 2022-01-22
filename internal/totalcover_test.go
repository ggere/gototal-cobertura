package totalcover_test

import (
	"errors"
	"strings"
	"testing"
	"testing/iotest"

	"github.com/stretchr/testify/assert"

	totalcover "github.com/ggere/gototal-cobertura/internal"
)

func TestTotalCoverage_GoodCase(t *testing.T) {
	r := strings.NewReader(`
	<?xml version="1.0"?>
	<coverage line-rate="1.00" branch-rate="0.14609273761902078" lines-covered="11411" lines-valid="51928" branches-covered="2593" branches-valid="17749" complexity="2.054109364767518" version="2.0.3" timestamp="1403301904999">
	</coverage>
	`)
	p, err := totalcover.TotalCover(r)
	assert.Nil(t, err)
	assert.Equal(t, float64(100), p)
}

func TestTotalCoverage_ReaderError(t *testing.T) {
	r := iotest.ErrReader(errors.New("error"))
	p, err := totalcover.TotalCover(r)
	assert.NotNil(t, err)
	assert.Equal(t, float64(0), p)
}

func TestTotalCoverage_UnmarshalError(t *testing.T) {
	r := strings.NewReader(`
	<?xml version="1.0"?>
	<ASDASDSADSADSASASA
	`)
	p, err := totalcover.TotalCover(r)
	assert.NotNil(t, err)
	assert.Equal(t, float64(0), p)
}

func TestTotalCoverage_ParseFloatError(t *testing.T) {
	r := strings.NewReader(`
	<?xml version="1.0"?>
	<coverage>
	</coverage>
	`)
	p, err := totalcover.TotalCover(r)
	assert.NotNil(t, err)
	assert.Equal(t, float64(0), p)
}
