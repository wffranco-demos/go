package factory

import "testing"

func TestNewLaptop(t *testing.T) {
	laptop, err := GetComputerFactory("laptop", 10)
	if err != nil {
		t.Error("Error: ", err)
	} else {
		PrintDetails(laptop)
	}
	if laptop.GetName() != "Laptop" {
		t.Error("Expected Laptop, got ", laptop.GetName())
	}
	if laptop.GetStock() != 10 {
		t.Error("Expected 10, got ", laptop.GetStock())
	}
}

func TestNewDesktop(t *testing.T) {
	desktop, err := GetComputerFactory("desktop", 20)
	if err != nil {
		t.Error("Error: ", err)
	} else {
		PrintDetails(desktop)
	}
	if desktop.GetName() != "Desktop" {
		t.Error("Expected Desktop, got ", desktop.GetName())
	}
	if desktop.GetStock() != 20 {
		t.Error("Expected 20, got ", desktop.GetStock())
	}
}
