package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
    gross_units := map[string]int{}
    gross_units["quarter_of_a_dozen"] = 3
    gross_units["half_of_a_dozen"] = 6
    gross_units["dozen"] = 12 
    gross_units["small_gross"] = 120
    gross_units["gross"] =  144
    gross_units["great_gross"] = 1728
    
  
    return gross_units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
    bill := make(map[string]int)
	return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	value,exits := units[unit]

    if !exits{
        return false
    }
   
     bill[item]  += value

    return true 
   
    
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
		rpieces, rexists := units[unit]
    	bpieces, bexists := bill[item]
    	if !rexists || !bexists || rpieces > bpieces {
    		return false
    	} else if rpieces == bpieces {
    		delete(bill, item)
    	} else {
    		bill[item] -= rpieces
    	}
    	
    	return true
    
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	pieces, exists := bill[item]
	return pieces, exists
}
