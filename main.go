package main

import (
  "log"
)

type Person struct {
  Id int
  Name string
  Website string
}

var (
	TableName = "people"
	RegionName = "eu-central-1"
	Id = 1234567
)

func init() {
  dynamo = connectDynamo()
}

func main() {
  err := CreateTable()
  if err != nil {
    log.Println(err)
  }

  // make new person
  person := Person {
    Id: Id,
    Name: "Mathis",
    Website: "https://mathisvaneetvelde.com",
  }

  PrintAndGet(Id)

  err = PutItem(person)
  if err != nil {
    log.Println(err)
  }

  PrintAndGet(Id)

  // change person name
  person.Name = "John Mayer"
  person.Website = "https://johnmayer.com/"

  err = UpdateItem(person)
  if err != nil {
    log.Println(err)
  }

  PrintAndGet(Id)

  err = DeleteItem(person)
  if err != nil {
    log.Println(err)
  }

  PrintAndGet(Id)
}