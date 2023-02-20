package parser_test

import (
	"github.com/stretchr/testify/require"
	"github.com/wailsapp/wails/v3/internal/parser"
	"testing"
)

func TestNewApplicationContext(t *testing.T) {
	appContext := parser.NewApplicationContext("../../examples/binding")
	require.NotNil(t, appContext)

	require.Len(t, appContext.BoundCandidates, 2)
}
