package main

import "testing"

func Test_fuelDistance(t *testing.T) {

	tests := []struct {
		name            string
		fuelConsumption int
		distance        int
		want            int
	}{

		{"fuelAvailable", 10, 20, 200},
	}
	for _, test := range tests {
		got := fuelDistance(test.fuelConsumption)
		if got != test.want {
			t.Error("Distance", "got:", got, "want:", test.want)
		}

	}
}
