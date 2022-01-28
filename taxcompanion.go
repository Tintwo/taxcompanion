package taxcompanion

// Excl tax (HT), Incl tax (TTC), VAT (TVA)
// TODO: secure operation with corrected round everywhere

type TaxCompanionValue struct {
	value float64
	dec   uint8
}

func (tc *TaxCompanionValue) InclFromExcl(excl float64) float64 {
	// use E2 (or E3) to escape wrong float precision
	return float64(int64(excl*tc.getExp()*(1+tc.value))) / tc.getExp()
}

func (tc *TaxCompanionValue) ExclFromIncl(incl float64) float64 {
	// use E2 (or E3) to escape wrong float precision
	return float64(int64(incl*tc.getExp()/(1+tc.value))) / tc.getExp()
}

func (tc *TaxCompanionValue) TaxFromExcl(excl float64) float64 {
	return excl * tc.value
}

func (tc *TaxCompanionValue) TaxFromIncl(incl float64) float64 {
	return tc.ExclFromIncl(incl) * tc.value
}

func (tc *TaxCompanionValue) ExclFromTax(tax float64) float64 {
	return tax * (1 / tc.value)
}

func (tc *TaxCompanionValue) InclFromTax(tax float64) float64 {
	// use E2 (or E3) to escape wrong float precision
	return float64(int64(tax*(((1/tc.value)+1)*tc.getExp()))) / tc.getExp()
}

func (tc TaxCompanionValue) getExp() float64 {
	switch tc.dec {
	case 2:
		return 100.
	case 3:
		return 1000.
	default:
		return 1.
	}
}
