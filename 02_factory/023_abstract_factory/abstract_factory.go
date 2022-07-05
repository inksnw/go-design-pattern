package main

import (
	"fmt"
)

type Table interface {
	createTable()
}

type TableFactory interface {
	CreateIns() Table
}

type mysql struct{}

func (*mysql) createTable() {
	fmt.Print("mysql table create \n")
}

type mysqlTableFactory struct{}

func (*mysqlTableFactory) CreateIns() Table {
	return &mysql{}
}

type pg struct{}

func (*pg) createTable() {
	fmt.Print("pg table create\n")
}

type PgTableFactory struct{}

func (*PgTableFactory) CreateIns() Table {
	return &pg{}
}

func abstractCreate(factory TableFactory) Table {
	//do something
	return factory.CreateIns()
}

func main() {
	factory1 := &mysqlTableFactory{}
	ins1 := abstractCreate(factory1)
	ins1.createTable()
	//
	factory2 := &PgTableFactory{}
	ins2 := abstractCreate(factory2)
	ins2.createTable()
}
