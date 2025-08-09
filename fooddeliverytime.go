package piscine

func FoodDeliveryTime(order string) int {
	type food struct {
		preptime int
	}

	menu := map[string]food{
		"burger":  {preptime: 15},
		"chips":   {preptime: 10},
		"nuggets": {preptime: 12},
	}

	if item, i := menu[order]; i {
		return item.preptime
	}

	return 404
}
