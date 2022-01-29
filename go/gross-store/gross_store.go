package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

func getUnits(units map[string]int, unit string) (int, bool) {
	u, ok := units[unit]
	return u, ok
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	u, ok := getUnits(units, unit)
	if ok {
		bill[item] += u
	}

	return ok
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	q, ok := bill[item]
	return q, ok
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	b, ok := GetItem(bill, item)
	if !ok {
		return false
	}

	u, ok := getUnits(units, unit)
	if !ok {
		return false
	}

	if b < u {
		return false
	}

	if b == u {
		delete(bill, item)
		return true
	}

	bill[item] = b - u
	return true
}
