package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	unitsMap := make(map[string]int)
	unitsMap["quarter_of_a_dozen"] = 3
	unitsMap["half_of_a_dozen"] = 6
	unitsMap["dozen"] = 12
	unitsMap["small_gross"] = 120
	unitsMap["gross"] = 144
	unitsMap["great_gross"] = 1728
	return unitsMap
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	valueOfUnit, exists := units[unit]
	if exists == false {
		return false
	}

	bill[item] += valueOfUnit
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	valueOfUnit, exists := units[unit]
	if exists == false {
		return false
	}

	billQuantity, existsInBill := bill[item]
	if existsInBill == false {
		return false
	} else if billQuantity < valueOfUnit {
		return false
	} else if billQuantity == valueOfUnit {
		delete(bill, item)
	} else {
		bill[item] -= valueOfUnit
	}
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	quantity, exists := bill[item]
	return quantity, exists
}
