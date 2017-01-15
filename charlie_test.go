package charlie

import (
	"fmt"
	"testing"
)

func TestWinning(t *testing.T) {

	score := Winning(0)
	score = 10
	score = "charlie"

	if _, ok := score.(string); !ok {
		printType(score)
		t.Fatal("Expected: string")
	}

	score = 50.20
	if _, ok := score.(float64); !ok {
		printType(score)
		t.Fatal("Expected: float64")
	}

}

func printType(anything interface{}) {
	switch val := anything.(type) {
	case string:
		fmt.Println("string", val)
	case int, int32, int64:
		fmt.Println("int", val)
	case float32, float64:
		fmt.Println("float", val)
	default:
		fmt.Println("unknown", val)
	}
}
