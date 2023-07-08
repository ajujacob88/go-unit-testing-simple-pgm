package math

import (
	"testing"
)

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

//ref : https://medium.com/code-chasm/unit-test-in-go-8d17d3910818

func TestFactorial(t *testing.T) {
	type args struct {
		n uint
	}
	tests := []struct { //This line declares a slice called tests. Each element of the slice is a struct with three fields: name (string), args (of type args struct), and want (unsigned integer).
		name string
		args args
		want uint
	}{
		{
			name: "N is 0",
			args: args{n: 0},
			want: 1,
		},
		{
			name: "N is 1",
			args: args{n: 1},
			want: 1,
		},
		{
			name: "N is 3",
			args: args{n: 3},
			want: 6,
		},
		{
			name: "N is 10",
			args: args{n: 10},
			want: 3628800,
		},
		//to check a fail case
		{
			name: "N is 5, just for sample to check fail case",
			args: args{n: 5},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Factorial(tt.args.n); got != tt.want {
				t.Errorf("got Factorial()=%v, Failed, want %v", got, tt.want)
			} else {
				t.Logf("got Factorial()=%v, Passed, want %v", got, tt.want)
			}
		})
	}
}
