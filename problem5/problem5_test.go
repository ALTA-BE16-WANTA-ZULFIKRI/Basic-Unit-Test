package main

import (
		"testing"
		"math"
)

//TestPrintLuasPermukaan: melakukan pengujian dengan radius dan tinggi yang bernilai positif, 
//sehingga hasil perhitungan yang diharapkan dapat diperoleh. 
func TestPrintLuasPermukaan(t *testing.T) {
	var radius float64 = 4
	var tinggi float64 = 20 
	var expected float64 = 361.679773
	var tolerance float64 = 1e-6

	result := PrintLuasPermukaan(radius, tinggi)

	if math.Abs(result-expected) > tolerance {
		t.Errorf("Expected %f, but got %f", expected, result)
	}
}


//TestPrintLuasPermukaanZero: melakukan pengujian dengan radius dan tinggi yang bernilai 0,
//sehingga hasil perhitungan yang diharapkan juga harus 0.
func TestPrintLuasPermukaanZero(t *testing.T) {
	var radius float64 = 0
	var tinggi float64 = 0 
	var expected float64 = 0
	var tolerance float64 = 1e-6

	result := PrintLuasPermukaan(radius, tinggi)

	if math.Abs(result-expected) > tolerance {
		t.Errorf("Expected %f, but got %f", expected, result)
	}
}
// TestPrintLuasPermukaanNegative: melakukan pengujian dengan radius dan tinggi yang berniali negatif,
// sehingga hasil perhitungan yang diharapkan adalah 0 
func TestPrintLuasPermukaanNegative(t *testing.T) {
	var radius float64 = -2 
	var tinggi float64 = -10
	var expected float64 = 0
	var tolerance float64 = 1e-6

	result := PrintLuasPermukaan(radius, tinggi)

	if math.Abs(result-expected) > tolerance {
        t.Errorf("Expected %f, but got %f", expected, result)
    }
}	

//TestPrintLuasPermukaanLargeRadius: melakukan pengujian dengan radius yang sangat besar, sehingga 
// dapat diuji apakah program dapat menangani kasus tersebut.
func TestPrintLuasPermukaanLargeRadius(t *testing.T) {
	var radius float64 = 1000
	var tinggi float64 = 10
	var expected float64 = 6283185.307179586
	var tolerance float64 = 1e-6

	result := PrintLuasPermukaan(radius, tinggi)

	if math.Abs(result-expected) > tolerance {
        t.Errorf("Expected %f, but got %f", expected, result)
    }
}
// Test PrintLuasPermukaanLargeHeight: melakukan pengujian dengan tinggi yang sangat besar,
// sehingga dapat diuji apakah program dapat menangani kasus tersebut. 
func TestPrintLuasPermukaanNegativeLargeHeight(t *testing.T) {
	var radius float64 = 5
    var tinggi float64 = 10000
    var expected float64 = 125663.70614359172
    var tolerance float64 = 1e-6

    result := PrintLuasPermukaan(radius, tinggi)

    if math.Abs(result-expected) > tolerance {
        t.Errorf("Expected %f, but got %f", expected, result)
    }
}