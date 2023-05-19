package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestDefaultGameType returns the default dispute game type.
func TestDefaultGameType(t *testing.T) {
	require.Equal(t, AttestationDisputeGameType, DefaultGameType())
}

// TestGameType_Valid tests the Valid function with valid inputs.
func TestGameType_Valid(t *testing.T) {
	require.True(t, AttestationDisputeGameType.Valid())
	require.True(t, FaultDisputeGameType.Valid())
	require.True(t, ValidityDisputeGameType.Valid())
}

// TestGameType_Invalid tests the Valid function with an invalid input.
func TestGameType_Invalid(t *testing.T) {
	require.False(t, GameType(ValidityDisputeGameType+1).Valid())
}

// FuzzGameType_Invalid checks that invalid game types are correctly
// returned as invalid by the validation [Valid] function.
func FuzzGameType_Invalid(f *testing.F) {
	maxCount := len(DisputeGameTypes)
	f.Fuzz(func(t *testing.T, number uint8) {
		if number >= uint8(maxCount) {
			require.False(t, GameType(number).Valid())
		} else {
			require.True(t, GameType(number).Valid())
		}
	})
}

// TestGameType_Default tests the default value of the DisputeGameType.
func TestGameType_Default(t *testing.T) {
	d := NewDisputeGameType()
	require.Equal(t, DefaultGameType(), d.selected)
	require.Equal(t, DefaultGameType(), d.Type())
}

// TestGameType_String tests the Set and String function on the DisputeGameType.
func TestGameType_String(t *testing.T) {
	var table []struct {
		name     string
		gameType GameType
	}
	for i, gameType := range DisputeGameTypes {
		table = append(table, struct {
			name     string
			gameType GameType
		}{name: gameType, gameType: GameType(i)})
	}
	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			d := NewDisputeGameType()
			require.Equal(t, tt.name, tt.gameType.String())
			require.NoError(t, d.Set(tt.name))
			require.Equal(t, tt.name, d.String())
			require.Equal(t, tt.gameType, d.selected)
		})
	}
}

// TestGameType_Type tests the Type function on the DisputeGameType.
func TestGameType_Type(t *testing.T) {
	var table []struct {
		name     string
		gameType GameType
	}
	for i, gameType := range DisputeGameTypes {
		table = append(table, struct {
			name     string
			gameType GameType
		}{name: gameType, gameType: GameType(i)})
	}
	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			d := NewDisputeGameType()
			require.Equal(t, tt.name, tt.gameType.String())
			require.NoError(t, d.Set(tt.name))
			require.Equal(t, tt.gameType, d.Type())
			require.Equal(t, tt.gameType, d.selected)
		})
	}
}
