package observer

import "testing"

func TestObserver(t *testing.T) {
	item := NewItem(1)
	customer1 := &Customer{ID: 1}
	customer2 := &Customer{ID: 2}
	customer3 := &Customer{ID: 3}

	item.Register(customer1)
	item.Register(customer2)
	item.Register(customer3)
	if len(item.Observers) != 3 {
		t.Error("Expected 3 observers, got ", len(item.Observers))
	}

	item.Notify()
	if !item.Available {
		t.Error("Expected item to be available")
	}

	item.Unregister(customer2)

	if len(item.Observers) != 2 {
		t.Error("Expected 2 observers, got ", len(item.Observers))
	}
}
