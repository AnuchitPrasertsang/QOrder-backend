package db

import mgo "gopkg.in/mgo.v2"

func Orders() *mgo.Collection {
	return Collection("orders")
}

func Menus() *mgo.Collection {
	return Collection("menus")
}
