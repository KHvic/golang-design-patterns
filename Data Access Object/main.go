package dao

import "fmt"

/*
	Type: Architectural
	Purpose: Object provides an abstract interface to some type of database or other persistence mechanism.
	Additional: - provides specific data operations without exposing details of the database
				- separate/decouple low level data accessing operations from high level business services
*/

// Data entity
type Customer struct {
	id   int
	name string
}

type CustomerDAO interface {
	GetAllCustomers() []*Customer
	GetCustomer(id int) *Customer
	AddCustomer(id int, name string)
	UpdateCustomer(id int, name string)
}

// Concrete DAO implementation to handle data operations
type FakeDAOImpl struct {
	customers map[int]*Customer
}

func (d *FakeDAOImpl) GetAllCustomers() []*Customer {
	fmt.Printf("SELECT * FROM Customer\n")
	result := make([]*Customer, 0, len(d.customers))
	for _, customer := range d.customers {
		result = append(result, customer)
	}
	return result
}

func (d *FakeDAOImpl) GetCustomer(id int) *Customer {
	fmt.Printf("SELECT * FROM Customer WHERE id = %d LIMIT 1\n", id)
	for _, customer := range d.customers {
		if customer.id == id {
			return customer
		}
	}
	return nil
}

func (d *FakeDAOImpl) AddCustomer(id int, name string) {
	fmt.Printf("INSERT INTO Customer (%d, %s)\n", id, name)
	_, exist := d.customers[id]
	if !exist {
		d.customers[id] = &Customer{id, name}
	}
}

func (d *FakeDAOImpl) UpdateCustomer(id int, name string) {
	fmt.Printf("UPDATE Customer SET name = %s WHERE id = %d\n", name, id)
	customer, ok := d.customers[id]
	if ok {
		customer.name = name
	}
}

func main() {
	// use dao without needing to specify low level SQL query, just abstract high level operation
	var dao CustomerDAO = &FakeDAOImpl{customers: make(map[int]*Customer)}
	dao.AddCustomer(1, "Bob")
	dao.AddCustomer(2, "Mary")
	customers := dao.GetAllCustomers()
	for _, c := range customers {
		fmt.Println(c.name)
	}

	dao.UpdateCustomer(2, "Jane")
	res := dao.GetCustomer(2)
	fmt.Println(res.name)
}
