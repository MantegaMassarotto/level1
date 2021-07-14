package change

import "errors"

var availableMoneyBills = []float64{200, 100, 50, 20, 10, 5, 2, 1, 0.50, 0.25, 0.10, 0.05, 0.01}

func GiveChange(paymentValue float64, billValue float64) ([]float64, error) {
	if paymentValue < billValue {
		return nil, errors.New("The payment amount is insufficient")
	}

	change := []float64{}

	changeAmount := paymentValue - billValue

	if changeAmount == 0 {
		return change, nil
	}

	changeRest := changeAmount
	for i := 0; i < len(availableMoneyBills); i++ {
		moneyBill := availableMoneyBills[i]
		if changeAmount < moneyBill {
			continue
		} else if changeRest > 0 {
			moneyBillQtd := int(changeRest / moneyBill)
			if moneyBillQtd >= 1 {
				for y := 0; y < moneyBillQtd; y++ {
					change = append(change, moneyBill)
					changeRest = changeRest - moneyBill
				}
			}
		}
	}

	return change, nil
}
