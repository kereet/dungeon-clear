package parser

import (
	"testing"
)

func TestParseLine(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		wantID       int
		wantPlayerID int
		wantExtra    string
		wantErr      bool
	}{
		{
			name:         "register event",
			input:        "[14:00:00] 1 1",
			wantID:       1,
			wantPlayerID: 1,
			wantExtra:    "",
			wantErr:      false,
		},
		{
			name:         "enter dungeon event",
			input:        "[14:00:00] 1 2",
			wantID:       2,
			wantPlayerID: 1,
			wantExtra:    "",
			wantErr:      false,
		},
		{
			name:         "kill the monster",
			input:        "[14:00:00] 1 3",
			wantID:       3,
			wantPlayerID: 1,
			wantExtra:    "",
			wantErr:      false,
		},
		{
			name:         "go to the next floor",
			input:        "[14:00:00] 1 4",
			wantID:       4,
			wantPlayerID: 1,
			wantExtra:    "",
			wantErr:      false,
		},
		{
			name:         "player can't continue",
			input:        "[14:00:00] 1 9 test test test",
			wantID:       9,
			wantPlayerID: 1,
			wantExtra:    "test test test",
			wantErr:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			event, err := ParseLine(tt.input)

			if (err != nil) != tt.wantErr {
				t.Errorf("ParseLine() error = %v, want %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if event.ID != tt.wantID {
					t.Errorf("ID = %d, want %d", event.ID, tt.wantID)
				}
				if event.PlayerID != tt.wantPlayerID {
					t.Errorf("PlayerID = %d, want %d", event.PlayerID, tt.wantID)
				}
				if event.Extra != tt.wantExtra {
					t.Errorf("Extra = %s, want %s", event.Extra, tt.wantExtra)
				}
			}
		})
	}
}
