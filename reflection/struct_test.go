package reflection

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStructAppendByField(t *testing.T) {
	type testStruct struct {
		S []string
		I []int
		i []int
		F float64
	}

	v := testStruct{
		S: []string{"data1", "data2"},
		I: []int{1, 2},
		i: []int{1, 2},
		F: 1.0,
	}

	t.Run("append_string", func(t *testing.T) {
		input := v
		err := StructSliceAppend(&input, "data3", "S")
		require.Nil(t, err)
		output := v
		output.S = append(output.S, "data3")
		require.Equal(t, output, input)
	})

	t.Run("append_int", func(t *testing.T) {
		input := v
		err := StructSliceAppend(&input, 3, "I")
		require.Nil(t, err)
		output := v
		output.I = append(output.I, 3)
		require.Equal(t, output, input)
	})

	t.Run("append_not_slice", func(t *testing.T) {
		input := v
		err := StructSliceAppend(&input, 3, "F")
		require.NotNil(t, err)
	})

	t.Run("append_not_settable", func(t *testing.T) {
		input := v
		err := StructSliceAppend(input, 3, "I")
		require.NotNil(t, err)
	})

	t.Run("append_not_settable_field", func(t *testing.T) {
		input := v
		err := StructSliceAppend(&input, 3, "i")
		require.NotNil(t, err)
	})

	t.Run("append_missed_field", func(t *testing.T) {
		input := v
		err := StructSliceAppend(&input, 3, "J")
		require.NotNil(t, err)
	})

}

func TestStructSet(t *testing.T) {
	type testStruct struct {
		I int
		i int
		S string
	}

	t.Run("set", func(t *testing.T) {
		var input testStruct
		err := StructSet(&input, 1, "I")
		require.Nil(t, err)
		err = StructSet(&input, "test", "S")
		require.Nil(t, err)
		require.Equal(t, 0, input.i)
		require.Equal(t, 1, input.I)
		require.Equal(t, "test", input.S)
	})

	t.Run("not_settable_input", func(t *testing.T) {
		var input testStruct
		err := StructSet(input, 1, "I")
		require.NotNil(t, err)
	})

	t.Run("not_settable_value", func(t *testing.T) {
		var input testStruct
		err := StructSet(input, 1, "i")
		require.NotNil(t, err)
	})

	t.Run("type", func(t *testing.T) {
		var input testStruct
		err := StructSet(&input, "wrong type", "I")
		require.NotNil(t, err)
	})
}

func TestSetStructValues(t *testing.T) {
	type testStruct struct {
		S1 string
		S2 string
		S3 string
	}

	t.Run("set normal data length", func(t *testing.T) {
		var input testStruct
		SetStructValues(&input, []string{"test", "", "3"})
		require.Equal(t, "test", input.S1)
		require.Equal(t, "", input.S2)
		require.Equal(t, "3", input.S3)
	})

	t.Run("set less data length", func(t *testing.T) {
		var input testStruct
		SetStructValues(&input, []string{"test", "test2"})
		require.Equal(t, "test", input.S1)
		require.Equal(t, "test2", input.S2)
		require.Equal(t, "", input.S3)
	})
}

func TestGetStructNumField(t *testing.T) {
	type testStruct struct {
		S1 string
		S2 string
		S3 string
	}

	t.Run("get normal structure numfield", func(t *testing.T) {
		var input testStruct
		require.Equal(t, GetStructNumField(&input), 3)
	})
}
