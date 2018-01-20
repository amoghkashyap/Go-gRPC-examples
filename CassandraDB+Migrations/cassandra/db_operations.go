package cassandra

import (
	"Go-gRPC-examples/CassandraDB+Migrations/cassandra/entities"
	"log"
)

var(
	//database manager instance
	dbManager = GetInstance()

	name string
	age int
	emailId string
)

const(
	//Query statements for biodata
	BioInsertQuery = "INSERT INTO info (emailId, name, age) VALUES (?,?,?)"
	UpdateAgeQuery = "UPDATE info SET age = ? WHERE emailId = ?"
	BioDeleteQuery = "DELETE FROM info WHERE emailId = ?"
	BioDeleteAllQuery = "TRUNCATE info"
	BioFindNameQuery = "SELECT * FROM info WHERE name = ? LIMIT 1 ALLOW FILTERING"
	BioFindAllDetailsQuery = "SELECT * FROM info"
)

/*
 These functions are responsible for handling the database operations
 Function names will describe the operation it performs
*/

// Database operations for inserting a Biodata
func InsertBio(biodata entities.Biodata){
	query := dbManager.Session.Query(BioInsertQuery,biodata.GetEmailID(),biodata.GetName(),biodata.GetAge())
	if err := query.Exec(); err != nil {
		log.Println("Biodata entry successful")
	} else {
		log.Println("Biodata entry Failed error: %v",err)
	}
}

// Database operations for updating age where emailId is provided
func EditAge(biodata entities.Biodata){
	query := dbManager.Session.Query(UpdateAgeQuery,biodata.GetAge(),biodata.GetEmailID())
	if err := query.Exec(); err != nil {
		log.Println("Updating age entry successful")
	} else {
		log.Println("Updating age entry Failed error: %v",err)
	}
}

// Database operations for deleting a biodata where emailId is provided
func DeleteBioWithEmailID(biodata entities.Biodata){
	query := dbManager.Session.Query(BioDeleteQuery,biodata.GetEmailID())
	if err := query.Exec(); err != nil {
		log.Println("Deleting entry successful")
	} else {
		log.Println("Deleting entry Failed error: %v",err)
	}
}

// Database operations for deleting all biodata stored in database
func DeleteAllBiodata(biodata entities.Biodata){
	query := dbManager.Session.Query(BioDeleteAllQuery)
	if err := query.Exec(); err != nil {
		log.Println("Failed to delete all Biodata from database", err)
	} else {
		log.Println("Biodata deletion successful")

	}
}

// Database operations for finding biodata where name is provided
func FindBioWithName(biodata entities.Biodata){
	query := dbManager.Session.Query(BioFindNameQuery,biodata.GetName())
	if err := query.Scan(&emailId,&name,&age); err != nil {
		log.Println("Biodata: %s %s %s", emailId, name, age)
	}
}

// Database operations for finding all biodata entries in database
func FindAllBiodata() {
	query := dbManager.Session.Query(BioFindAllDetailsQuery)
	iterator := query.Iter()
	for iterator.Scan(&emailId,&name,&age){
		log.Println(emailId,name,age)
	}
}


