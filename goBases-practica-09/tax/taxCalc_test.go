package tax

import (
	"fmt"
	"testing"
)

func TestTaxCalculator(t *testing.T) {
	// arrange
	salaryMin50 := 25000.0
	salaryMax50 := 65000.0
	salaryMax150 := 200000.0

	salaryMax50Tax := salaryMax50 * (0.83)
	salaryMax150Tax := salaryMax150 * (0.73)

	// act
	taxResultMin50 := taxCalculator(salaryMin50)
	taxResultMax50 := taxCalculator(salaryMax50)
	taxResultMax150 := taxCalculator(salaryMax150)

	// assert
	if taxResultMin50 != salaryMin50 {
		t.Errorf("Expected %f, got %f", salaryMin50, taxResultMin50)
	} else {
		fmt.Println("OK!")
	}

	if taxResultMax50 != taxResultMax50 {
		t.Errorf("Expected %f, got %f", salaryMax50Tax, taxResultMax50)
	} else {
		fmt.Println("OK!")
	}

	if taxResultMax150 != salaryMax150Tax {
		t.Errorf("Expected %f, got %f", salaryMax150Tax, taxResultMax150)
	} else {
		fmt.Println("OK!")
	}

}
