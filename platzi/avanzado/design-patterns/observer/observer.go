package observer

import "fmt"

type Topic interface {
	Register(observer Observer)
	Notify()
	Unregister(observer Observer)
}

// An Observer is like a subscriber, waiting for any changes in the topic
type Observer interface {
	GetID() int
	Update(int)
}

type Item struct {
	Observers []Observer
	ID        int
	Available bool
}

func NewItem(id int) *Item {
	return &Item{
		ID:        id,
		Observers: make([]Observer, 0),
	}
}

func (i *Item) Register(observer Observer) {
	i.Observers = append(i.Observers, observer)
}

func (i *Item) Notify() {
	i.Available = true
	for _, observer := range i.Observers {
		observer.Update(i.ID)
	}
}

func (i *Item) Unregister(observer Observer) {
	for index, o := range i.Observers {
		if o.GetID() == observer.GetID() {
			i.Observers = append(i.Observers[:index], i.Observers[index+1:]...)
			break
		}
	}
}

type Customer struct {
	ID int
}

func (c *Customer) GetID() int {
	return c.ID
}

func (c *Customer) Update(itemID int) {
	fmt.Printf("Sending email to customer %d for item %d\n", c.ID, itemID)
}
