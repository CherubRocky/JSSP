package cs
import "testing"

func TestHatch(t *testing.T) {
	cs := CuckooSearch{}
	egg := cs.Hatch()
	if egg == nil {
		t.Fatal("No sirve")
	}
}
