package cyrillicfilter

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCyrillicFilter(t *testing.T) {
	type testStruct struct {
		id           int
		field        string
		pointerField *string
		insideStruct *struct {
			id           int
			field        string
			pointerField *string
		}
	}

	s1 := "Goodbye -> пока!"
	sample := testStruct{
		id:           0,
		field:        "Hello -> привет!",
		pointerField: &s1,
		insideStruct: &struct {
			id           int
			field        string
			pointerField *string
		}{id: 0, field: "Это структура внутри структуры -> This is the struct inside struct!", pointerField: &s1},
	}
	if err := CyrillicFilter(&sample); err != nil {
		_ = fmt.Errorf("error: %v", err)
	}

	s2 := "Goodbye -> !"
	expected := testStruct{
		id:           0,
		field:        "Hello -> !",
		pointerField: &s2,
		insideStruct: &struct {
			id           int
			field        string
			pointerField *string
		}{id: 0, field: "    -> This is the struct inside struct!", pointerField: &s2},
	}

	assert.Equal(t, expected, sample)
	assert.Equal(t, NilInterfaceError, CyrillicFilter(nil))
	assert.Equal(t, NotStructError, CyrillicFilter(&s1))
}
