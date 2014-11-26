package colormatcher

import (
	"fmt"
	"testing"
)

func TestMath(t *testing.T) {
	c := "#4ab02b"
	colors := []string{
		"#F7C394",
		"#F5A763",
		"#F28B30",
		"#EF6630",
		"#ED462F",
		"#DE96C4",
		"#CC6AA8",
		"#BD428F",
		"#A240A6",
		"#642E8E",
		"#2A2DA0",
		"#2578B4",
		"#46ABBD",
		"#6BC7BE",
		"#4DA551",
		"#6FBA45",
		"#ADD136",
		"#F7BC32",
		"#F9D532",
		"#FFF134",
	}
	r, v, err := GetClosest(c, colors...)
	fmt.Println("This is match", r, v, err)
	if err != nil {
		t.Fail()
	}
	if r != "#6FBA45" {
		t.Fail()
	}
}

func TestConvert(t *testing.T) {
	v, err := hashtorgb("#DE96C4")
	if err != nil {
		t.Fail()
	}
	if v.R != 222 || v.G != 150 || v.B != 196 {
		t.Fail()
	}
	hsv := rgbtohsv(v)
	// fmt.Println("HSV: ", hsv)
	if int(hsv.H) != 321 || int(hsv.S) != 32 || int(hsv.V) != 87 {
		t.Fail()
	}
}
