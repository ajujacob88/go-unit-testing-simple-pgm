package math

import "testing"

type MultiplyData struct {
	x, y   int
	result int
}

func TestAdd(t *testing.T) {
	result := Add(1, 3)
	if result != 4 {
		t.Errorf("Add(1,3) FAILED. Expected %d, got %d\n", 4, result)
	} else {
		t.Logf("Add(1,3) PASSED, Expected %d, got %d\n", 4, result)
	}
}

func TestDivide(t *testing.T) {
	result := Divide(5, 0)

	if result != 0 {
		t.Errorf("DIvide(5,0) Failed, Expected %f, got %f", 0.0, result)
	} else {
		t.Logf("DIvide(5,0) PASSED, Expected %f, got %f", 0.0, result)
	}

}

func TestMultiply(t *testing.T) {
	testData := []MultiplyData{
		{1, 2, 2},
		{3, 5, 15},
		{7, 4, 28},
		{10, 0, 0},
		{-8, 5, -40},
	}
	for _, datum := range testData {
		result := Multiply(datum.x, datum.y)

		if result != datum.result {
			t.Errorf("Multiply(%d,%d) Failed. Expected %d, Got %d", datum.x, datum.y, datum.result, result)
		} else {
			t.Logf("Multiply(%d,%d) Passed. Expected %d, Got %d", datum.x, datum.y, datum.result, result)
		}
	}
}
