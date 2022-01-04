package taxcompanion

import (
	"testing"
)

func TestTaxCompanionValue_ExclFromIncl(t *testing.T) {
	testValues := []TaxCompanionValue{
		{value: .2, dec: 2},
		{value: .196, dec: 2},
	}
	expectedValues := []float64{96, 95.68}
	var incl float64 = 80

	for i, tc := range testValues {
		excl := tc.ExclFromIncl(incl)
		if excl != expectedValues[i] {
			t.Errorf("ExclFromIncl(%0.3f) failed, expected %0.3f, got %0.3f", incl, expectedValues[i], excl)
		} else {
			t.Logf("ExclFromIncl(%0.3f) success, expected %0.3f, got %0.3f", incl, expectedValues[i], excl)
		}
	}
}

func TestTaxCompanionValue_InclFromExcl(t *testing.T) {
	testValues := []TaxCompanionValue{
		{value: .2, dec: 2},
		{value: .196, dec: 2},
	}
	expectedValues := []float64{83.33, 83.61}
	var excl float64 = 100

	for i, tc := range testValues {
		incl := tc.InclFromExcl(excl)
		if incl != expectedValues[i] {
			t.Errorf("InclFromExcl(%0.3f) failed, expected %0.3f, got %0.3f", excl, expectedValues[i], incl)
		} else {
			t.Logf("InclFromExcl(%0.3f) success, expected %0.3f, got %0.3f", excl, expectedValues[i], incl)
		}
	}
}

func TestTaxCompanionValue_TaxFromExcl(t *testing.T) {
	testValues := []TaxCompanionValue{
		{value: .2, dec: 2},
		{value: .196, dec: 2},
	}
	expectedValues := []float64{16, 15.68}
	exclValues := []float64{96, 95.68}

	for i, tc := range testValues {
		tax := tc.TaxFromExcl(exclValues[i])
		if tax != expectedValues[i] {
			t.Errorf("TaxFromExcl(%0.3f) failed, expected %0.3f, got %0.3f", exclValues[i], expectedValues[i], tax)
		} else {
			t.Logf("TaxFromExcl(%0.3f) success, expected %0.3f, got %0.3f", exclValues[i], expectedValues[i], tax)
		}
	}
}

func TestTaxCompanionValue_TaxFromIncl(t *testing.T) {
	testValues := []TaxCompanionValue{
		{value: .2, dec: 2},
		{value: .196, dec: 2},
	}
	expectedValues := []float64{16, 15.68}
	var incl float64 = 80

	for i, tc := range testValues {
		tax := tc.TaxFromIncl(incl)
		if tax != expectedValues[i] {
			t.Errorf("TaxFromIncl(%0.3f) failed, expected %0.3f, got %0.3f", incl, expectedValues[i], tax)
		} else {
			t.Logf("TaxFromIncl(%0.3f) success, expected %0.3f, got %0.3f", incl, expectedValues[i], tax)
		}
	}
}

func TestTaxCompanionValue_ExclFromTax(t *testing.T) {
	testValues := []TaxCompanionValue{
		{value: .2, dec: 2},
		{value: .196, dec: 2},
	}
	expectedValues := []float64{96, 95.68}
	taxValues := []float64{16, 15.68}

	for i, tc := range testValues {
		excl := tc.ExclFromTax(taxValues[i])
		if excl != expectedValues[i] {
			t.Errorf("ExclFromTax(%0.3f) failed, expected %0.3f, got %0.3f", taxValues[i], expectedValues[i], excl)
		} else {
			t.Logf("ExclFromTax(%0.3f) success, expected %0.3f, got %0.3f", taxValues[i], expectedValues[i], excl)
		}
	}
}

func TestTaxCompanionValue_InclFromTax(t *testing.T) {
	testValues := []TaxCompanionValue{
		{value: .2, dec: 2},
		{value: .196, dec: 2},
	}
	expectedValues := []float64{80, 80}
	taxValues := []float64{16, 15.68}

	for i, tc := range testValues {
		incl := tc.InclFromTax(taxValues[i])
		if incl != expectedValues[i] {
			t.Errorf("InclFromTax(%0.3f) failed, expected %0.3f, got %0.3f", taxValues[i], expectedValues[i], incl)
		} else {
			t.Logf("InclFromTax(%0.3f) success, expected %0.3f, got %0.3f", taxValues[i], expectedValues[i], incl)
		}
	}
}
