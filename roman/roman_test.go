package roman

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldReturnRomanRepresentationOf1(t *testing.T) {
	assert.Equal(t, AsRoman(1), "I")
}

func TestShouldReturnRomanRepresentationOf2(t *testing.T) {
	assert.Equal(t, AsRoman(2), "II")
}

func TestShouldReturnRomanRepresentationOf3(t *testing.T) {
	assert.Equal(t, AsRoman(3), "III")
}

func TestShouldReturnRomanRepresentationOf4(t *testing.T) {
	assert.Equal(t, AsRoman(4), "IV")
}

func TestShouldReturnRomanRepresentationOf5(t *testing.T) {
	assert.Equal(t, AsRoman(5), "V")
}

func TestShouldReturnRomanRepresentationOf6(t *testing.T) {
	assert.Equal(t, AsRoman(6), "VI")
}

func TestShouldReturnRomanRepresentationOf9(t *testing.T) {
	assert.Equal(t, AsRoman(9), "IX")
}

func TestShouldReturnRomanRepresentationOf10(t *testing.T) {
	assert.Equal(t, AsRoman(10), "X")
}

func TestShouldReturnRomanRepresentationOf11(t *testing.T) {
	assert.Equal(t, AsRoman(11), "XI")
}

func TestShouldReturnRomanRepresentationOf14(t *testing.T) {
	assert.Equal(t, AsRoman(14), "XIV")
}

func TestShouldReturnRomanRepresentationOf20(t *testing.T) {
	assert.Equal(t, AsRoman(20), "XX")
}

func TestShouldReturnRomanRepresentationOf40(t *testing.T) {
	assert.Equal(t, AsRoman(40), "XL")
}

func TestShouldReturnRomanRepresentationOf50(t *testing.T) {
	assert.Equal(t, AsRoman(50), "L")
}

func TestShouldReturnRomanRepresentationOf60(t *testing.T) {
	assert.Equal(t, AsRoman(60), "LX")
}

func TestShouldReturnRomanRepresentationOf90(t *testing.T) {
	assert.Equal(t, AsRoman(90), "XC")
}

func TestShouldReturnRomanRepresentationOf100(t *testing.T) {
	assert.Equal(t, AsRoman(100), "C")
}

func TestShouldReturnRomanRepresentationOf400(t *testing.T) {
	assert.Equal(t, AsRoman(400), "CD")
}

func TestShouldReturnRomanRepresentationOf500(t *testing.T) {
	assert.Equal(t, AsRoman(500), "D")
}

func TestShouldReturnRomanRepresentationOf900(t *testing.T) {
	assert.Equal(t, AsRoman(900), "CM")
}

func TestShouldReturnRomanRepresentationOf1000(t *testing.T) {
	assert.Equal(t, AsRoman(1000), "M")
}
