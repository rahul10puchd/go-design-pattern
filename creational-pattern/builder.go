package main

import (
	"errors"
	"log"
)

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

//Director
type ManufacturingDirector struct {
	builder BuildProcess
}

func (f *ManufacturingDirector) Construct() {
	log.Println("hi")
	f.builder.SetSeats().SetStructure().SetWheels()
}

func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}

//Product
type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

//A Builder of type car
type CarBuilder struct {
	v VehicleProduct
	e []error
}

func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	c.e = append(c.e, errors.New("fucked"))
	return c
}

func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = 5
	return c
}

func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Car"
	return c
}

func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}
func main() {
	manufacturingComplex := ManufacturingDirector{}

	carBuilder := &CarBuilder{}
	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()
	car := carBuilder.GetVehicle()
	// car, err := carBuilder.GetVehicle()
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	if car.Wheels != 4 {
		log.Println(car.Wheels)
	}

	if car.Structure != "Car" {
		log.Println(car.Structure)
	}

	if car.Seats != 5 {
		log.Println(car.Seats)
	}
}

//A Builder of type motorbike
// type BikeBuilder struct {
// 	v VehicleProduct
// }

// func (b *BikeBuilder) SetWheels() BuildProcess {
// 	b.v.Wheels = 2
// 	return b
// }

// func (b *BikeBuilder) SetSeats() BuildProcess {
// 	b.v.Seats = 2
// 	return b
// }

// func (b *BikeBuilder) SetStructure() BuildProcess {
// 	b.v.Structure = "Motorbike"
// 	return b
// }

// func (b *BikeBuilder) GetVehicle() VehicleProduct {
// 	return b.v
// }

// //A Builder of type motorbike
// type BusBuilder struct {
// 	v VehicleProduct
// }

// func (b *BusBuilder) SetWheels() BuildProcess {
// 	b.v.Wheels = 8
// 	return b
// }

// func (b *BusBuilder) SetSeats() BuildProcess {
// 	b.v.Seats = 30
// 	return b
// }

// func (b *BusBuilder) SetStructure() BuildProcess {
// 	b.v.Structure = "Bus"
// 	return b
// }

// func (b *BusBuilder) GetVehicle() VehicleProduct {
// 	return b.v
// }
