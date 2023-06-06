package models

type Role struct {
	Customer []Customer
}

type Cargo struct {
	Role                   Role
	DeliveryHistory        DeliveryHistory
	DeliveryEspecification DeliveryEspecification
}

func TestSomething() {
	// cargo := Cargo{}
	// vamo := Zap{}

}
