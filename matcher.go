package colormatcher

import (
	"errors"
	// "log"
	"math"
	"strconv"
	"strings"
)

type RGB struct {
	R float64
	G float64
	B float64
}

type HSV struct {
	H float64
	S float64
	V float64
}

type candidate struct {
	c     HSV
	hash  string
	delta float64
}

func hashtorgb(hash string) (RGB, error) {
	rgb := RGB{}
	hash = strings.TrimLeft(strings.TrimSpace(hash), "#")
	if len(hash) != 6 {
		return rgb, errors.New("Invalid format, should be #ab01ed")
	}

	v := make([]float64, 3)

	for i := 0; i < 3; i++ {
		z, err := strconv.ParseInt("0x"+hash[i*2:(i+1)*2], 0, 32)
		if err != nil {
			return rgb, err
		}
		v[i] = math.Max(0, math.Min(255, float64(z)))
	}
	rgb.R = v[0]
	rgb.G = v[1]
	rgb.B = v[2]

	return rgb, nil
}

func rgbtohsv(rgb RGB) HSV {
	hsv := HSV{}

	R := rgb.R / 255
	G := rgb.G / 255
	B := rgb.B / 255

	max := math.Max(R, math.Max(G, B))
	d := max - math.Min(R, math.Min(G, B))

	if d > 0 {
		if max == R {
			hsv.H = 60 * math.Mod((G-B)/d, 6)
		} else if max == G {
			hsv.H = 60 * ((B-R)/d + 2)
		} else if max == B {
			hsv.H = 60 * ((R-G)/d + 4)
		}
	}
	if hsv.H < 0 {
		hsv.H += 360
	}

	if max > 0 {
		hsv.S = d / max * 100
	}

	hsv.V = max * 100

	return hsv
}

func delta(c1, c2 HSV) float64 {
	d := 0.0
	hd := math.Abs(c1.H - c2.H)
	d += math.Min(hd, 360-hd) * 0.475
	// log.Println("Difference ", d, "parts", hd, math.Abs(c1.S-c2.S), math.Abs(c1.V-c2.V))
	d += math.Abs(c1.S-c2.S) * 0.2875
	d += math.Abs(c1.V-c2.V) * 0.2375
	return d
}

func GetClosest(current string, defined ...string) (string, float64, error) {
	c_rgb, err := hashtorgb(current)
	if err != nil {
		return "", 0, err
	}

	if len(defined) == 0 {
		return "", 0, errors.New("Provide at least one predefined color")
	}

	c_hsv := rgbtohsv(c_rgb)
	candidates := make([]candidate, len(defined))
	for i, d := range defined {
		rgb, err := hashtorgb(d)
		if err != nil {
			return "", 0, err
		}
		candidates[i].c = rgbtohsv(rgb)
		candidates[i].hash = d
		candidates[i].delta = delta(candidates[i].c, c_hsv)
	}

	min := 1000.0
	pos := 0
	for i := range candidates {
		// log.Println("Candidate", candidates[i])
		if candidates[i].delta < min {
			pos = i
			min = candidates[i].delta
		}
	}
	return candidates[pos].hash, candidates[pos].delta, nil
}
