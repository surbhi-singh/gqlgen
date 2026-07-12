package ast

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValueObject(t *testing.T) {
	t.Run("undefined variable", func(t *testing.T) {
		obj := &Value{Kind: ObjectValue, Children: ChildValueList{
			{Name: "field", Value: &Value{Kind: Variable, Raw: "Var"}},
		}}
		val, err := obj.Value(map[string]any{})
		require.NoError(t, err)
		// Treated as absent so the field's own default can be applied later.
		require.NotContains(t, val.(map[string]any), "field")
	})

	t.Run("explicit null", func(t *testing.T) {
		obj := &Value{Kind: ObjectValue, Children: ChildValueList{
			{Name: "field", Value: &Value{Kind: NullValue}},
		}}
		val, err := obj.Value(map[string]any{})
		require.NoError(t, err)
		m := val.(map[string]any)
		require.Contains(t, m, "field")
		require.Nil(t, m["field"])
	})
}
