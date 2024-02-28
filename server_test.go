package main

import (
	"testing"
)

func TestRuleOne(t *testing.T) {
	retailer := "Target"
	ans := ruleOne(&retailer)

	if ans != 6 {
		t.Errorf("Expected 6, got %d", ans)
	}
}

func TestRuleTwoFalse(t *testing.T) {
	var total float64 = 35.35
	ans := ruleTwo(&total)

	if ans != 0 {
		t.Errorf("Expected 0, got %d", ans)
	}
}

func TestRuleTwoTrue(t *testing.T) {
	var total float64 = 100.00
	ans := ruleTwo(&total)

	if ans != 50 {
		t.Errorf("Expected 50, got %d", ans)
	}
}

func TestRuleThreeFalse(t *testing.T) {
	var total float64 = 11.11
	ans := ruleThree(&total)

	if ans != 0 {
		t.Errorf("Expected 0, got %d", ans)
	}
}

func TestRuleThreeTrue(t *testing.T) {
	var total float64 = 11.25
	ans := ruleThree(&total)

	if ans != 25 {
		t.Errorf("Expected 25, got %d", ans)
	}
}

func TestRuleFour(t *testing.T) {
	items := make([]Item, 5)
	ans := ruleFour(&items)

	if ans != 10 {
		t.Errorf("Expected 10, got %d", ans)
	}
}

func TestRuleFive(t *testing.T) {
	item1 := new(Item)
	item1.ShortDescription = "Emils Cheese Pizza"
	item1.Price = "12.25"
	item2 := new(Item)
	item2.ShortDescription = "Klarbrunn 12-PK 12 FL OZ"
	item2.Price = "12.00"

	items := make([]Item, 2)
	items[0] = *item1
	items[1] = *item2

	ans := ruleFive(&items)

	if ans != 6 {
		t.Errorf("Expected 6, got %d", ans)
	}
}

func TestRuleSixFalse(t *testing.T) {
	purchaseDay := 12
	ans := ruleSix(&purchaseDay)

	if ans != 0 {
		t.Errorf("Expected 0, got %d", ans)
	}
}

func TestRuleSixTrue(t *testing.T) {
	purchaseDay := 13
	ans := ruleSix(&purchaseDay)

	if ans != 6 {
		t.Errorf("Expected 6, got %d", ans)
	}
}

func TestRuleSevenFalse(t *testing.T) {
	purchaseHour := 13
	ans := ruleSeven(&purchaseHour)

	if ans != 0 {
		t.Errorf("Expected 6, got %d", ans)
	}
}

func TestRuleSevenTrue(t *testing.T) {
	purchaseHour := 15
	ans := ruleSeven(&purchaseHour)

	if ans != 10 {
		t.Errorf("Expected 10, got %d", ans)
	}
}

func TestGetScore(t *testing.T) {
	item := new(Item)
	item.ShortDescription = "Emils Cheese Pizza"
	item.Price = "12.25"

	ans := getScore(item)

	if ans != 3 {
		t.Errorf("Expected 3, got %d", ans)
	}

	item.ShortDescription = "Klarbrunn 12-PK 12 FL OZ"
	item.Price = "12.00"

	ans = getScore(item)

	if ans != 3 {
		t.Errorf("Expected 3, got %d", ans)
	}

}

func TestGetPriceTarget(t *testing.T) {
	sum := 0

	retailer := "Target"
	day := 1
	purchaseHour := 13
	total := 35.35
	items := make([]Item, 5)

	item1 := new(Item)
	item1.ShortDescription = "Emils Cheese Pizza"
	item1.Price = "12.25"
	item2 := new(Item)
	item2.ShortDescription = "Klarbrunn 12-PK 12 FL OZ"
	item2.Price = "12.00"

	items[0] = *item1
	items[1] = *item2

	one := ruleOne(&retailer)
	two := ruleTwo(&total)
	three := ruleThree(&total)
	four := ruleFour(&items)
	five := ruleFive(&items)
	six := ruleSix(&day)
	seven := ruleSeven(&purchaseHour)

	if one != 6 {
		t.Errorf("Expected 6 for rule one, got %d", one)
	}

	if two != 0 {
		t.Errorf("Expected 0 for rule two, got %d", two)
	}

	if three != 0 {
		t.Errorf("Expected 0 for rule three, got %d", three)
	}

	if four != 10 {
		t.Errorf("Expected 10 for rule four, got %d", four)
	}

	if five != 6 {
		t.Errorf("Expected 6 for rule five, got %d", five)
	}

	if six != 6 {
		t.Errorf("Expected 6 for rule six, got %d", six)
	}

	if seven != 0 {
		t.Errorf("Expected 0 for rule seven, got %d", seven)
	}

	sum = one + two + three + four + five + six + seven

	if sum != 28 {
		t.Errorf("Expected 28, got %d", sum)
	}
}
