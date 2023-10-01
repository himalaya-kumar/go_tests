package nteger

import "testing"

func TestAdder(t *testing.T) {
	type value struct {
		x int
		y int
	}

	type testsStruct struct {
		name string
		val  value
		want int
	}
	tests := []testsStruct{
		{"2 + 2", value{x: 2, y: 2}, 4},
		{"5 + 1", value{x: 5, y: 1}, 6},
		{"5 + 4", value{x: 5, y: 4}, 9},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Add(test.val.x, test.val.y)
			if got != test.want {
				t.Errorf("expected '%d' but got '%d'", test.want, got)
			}
		})
	}

}
