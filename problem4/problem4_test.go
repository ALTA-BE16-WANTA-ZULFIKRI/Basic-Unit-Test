package main

import "testing"

// Test hitung luas segitaga dengan input valid
func TestHitungLuasSegitiga(t *testing.T) {
    // your code here
	var alas float64 = 20 
	var tinggi float64 = 25 
	expected := 250.0 

	result := PrintLuas(alas, tinggi)

	if result != expected {
		t.Errorf("Error: hasil %f, seharusnya %f", result, expected)
	}
}
// Test untuk memastikan bahwa fungsi PrintLuas mengembalikan nilai 0 ketika alas atau tinggi bernilai 0:
func TestPrintLuas_Zero(t *testing.T){
	var alas float64 = 0 
	var tinggi float64 = 25 
	var expected float64 = 0 

	result := PrintLuas(alas, tinggi)
	if result!= expected {
        t.Errorf("PrintLuas(%f, %f) = %f; expected %f", alas, tinggi, result, expected)
    }

	alas = 20 
	tinggi = 0 
	result = PrintLuas(alas, tinggi)
	if result != expected {
		t.Errorf("PrintLuas(%f, %f) = %f; expected %f", alas, tinggi, result, expected)
	}
}

// Test untuk memastikan bahwa fungsi PrintLuas mengembalikan nilai negatif ketika alas atau tinggi bernilai negatif:

func TestPrintLuas_Negatif(t *testing.T){
    var alas float64 = -20 
    var tinggi float64 = 25 
    var expected float64 = -250.0 

    result := PrintLuas(alas, tinggi)
    if result!= expected {
        t.Errorf("PrintLuas(%f, %f) = %f; expected %f", alas, tinggi, result, expected)
    }

    alas = 20 
    tinggi = -25 
    result = PrintLuas(alas, tinggi)
    if result!= expected {
        t.Errorf("PrintLuas(%f, %f) = %f; expected %f", alas, tinggi, result, expected)
    }
}

// Test untuk memastikan bahwa fungsi PrintLuas mengembalikan nilai yang benar ketika alas atau tinggi bernilai desimal:
func TestPrintLuas_Decimal(t *testing.T) {
	var alas float64 = 20.5 
	var tinggi float64 = 25.25
	var expected float64 = 258.53125

	result := PrintLuas(alas, tinggi)
	if result != expected {
		t.Errorf("PrintLuas(%f,%f) = %f; expected %f", alas, tinggi, result, expected)
	}
}

// test untuk memastikan bahwa fungsi PrintLuas mengembalikan nilai yang sama setiap kali dipanggil dengan input yang sama:
func TestPrintLuas(t *testing.T) {
	var alas float64 = 45.45
	var tinggi float64 = 35.35 
	var expected float64 = 258.53125
	
	result := PrintLuas(alas, tinggi)
	if result!= expected {
        t.Errorf("PrintLuas(%f,%f) = %f; expected %f", alas, tinggi, result, expected)
    }

}