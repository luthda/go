package gross

func Units() map[string]int {
	units := make(map[string]int)
	units["quarter_of_a_dozen"] = 3
	units["half_of_a_dozen"] = 6
	units["dozen"] = 12
	units["small_gross"] = 120
	units["gross"] = 144
	units["great_gross"] = 1728

	return units
}

func NewBill() map[string]int {
	return make(map[string]int)
}

func AddItem(bill, units map[string]int, item, unit string) bool {
	value, exists := units[unit]
	if  !exists {
		return false;
	}	
		bill[item] += value

	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	itemCount, existsInBill := bill[item]
	unitValue, existsInUnits := units[unit]

	if !existsInBill || !existsInUnits || itemCount - unitValue < 0{
		return false
	} else if itemCount - unitValue == 0 {
		delete(bill, item)
		return true
	}

	bill[item] -= unitValue

	return true
}

func GetItem(bill map[string]int, item string) (int, bool) {
	value, exists := bill[item]
	
	return value, exists
}
