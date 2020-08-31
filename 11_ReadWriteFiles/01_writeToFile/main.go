package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

func main()  {
	_, error := os.Create("testResults-v1.csv")
	if error != nil {
		return
	}
	c1 := Customer{custId: "1", name:"Ram", age:30, gender:"male", mEmpolyee:true}
	c2 := Customer{custId:"2", name:"Cathy", age:29, gender:"female", mEmpolyee:false}
	customers := []Customer{c1, c2}
	writeToPNFile(customers)
	
}

type Customer struct {
	custId string
	name string
	age int
	gender string
	mEmpolyee bool
}

func writeToPNFile(records []Customer) {
	file, error := os.OpenFile("testResults-v1.csv", os.O_APPEND|os.O_WRONLY, 0600)
	if error != nil {
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	for _, v := range records {
		things := []string{v.custId,
			v.name,
			strconv.Itoa(v.age),
			v.gender,
			strconv.FormatBool(v.mEmpolyee),
		}
		writer.Write(things)
	}
	defer writer.Flush()
}