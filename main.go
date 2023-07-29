package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const MAX_NAME = 30
const MAX_ADDRESS = 50
const MAX_PHONE = 15

var lastInserted int

/* The data must be capitalize so I can export this */
type Contact struct {
	Name      string `json:"name"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	IsDeleted bool   `json:"isDeleted"`
}

type Contacts struct {
	contacts []Contact
}

type Index struct {
	Key      string `json:"key"`
	Position int    `json:"position"`
}

type Indexes struct {
	indexes []Index
}

/* I read the data from the JSON file. */
func (contacts *Contacts) loadContacts() {
	data, err := ioutil.ReadFile("./data/contacts.json")
	checkErr(err)
	err = json.Unmarshal(data, &contacts.contacts)
	checkErr(err)
}

func (contacts *Contacts) viewAll() {
	fmt.Println("[")
	for i := 0; i < len(contacts.contacts); i++ {
		fmt.Println("\t{")
		fmt.Printf("\t\tNAME: %s\n", contacts.contacts[i].Name)
		fmt.Printf("\t\tADDRESS: %s\n", contacts.contacts[i].Address)
		fmt.Printf("\t\tPHONE: %s\n", contacts.contacts[i].Phone)
		fmt.Println("\t}")
	}
	fmt.Println("]")
}

func (contacts *Contacts) createNewContact(tree *BTree, indexes *Indexes) {
	Clear()
	fmt.Println("Creating a new contact!")
	var newContact Contact

	scanner := bufio.NewScanner(os.Stdin)

	// Previous varnames to validade before create
	var pName string
	var pAddress string
	var pPhone string

	fmt.Println("Name (len 30):")
	scanner.Scan()
	pName = scanner.Text()

	fmt.Println("Address (len 50):")
	scanner.Scan()
	pAddress = scanner.Text()

	fmt.Println("Phone (len 15):")
	scanner.Scan()
	pPhone = scanner.Text()

	// Data validators
	if len(pName) > MAX_NAME {
		pName = pName[:MAX_NAME]
	}

	if len(pAddress) > MAX_ADDRESS {
		pAddress = pAddress[:MAX_ADDRESS]
	}

	if len(pPhone) > MAX_PHONE {
		pPhone = pPhone[:MAX_PHONE]
	}

	newContact.Name = pName
	newContact.Address = pAddress
	newContact.Phone = pPhone
	newContact.IsDeleted = false
	contacts.contacts = append(contacts.contacts, newContact)
	newPosition := len(contacts.contacts)
	var newIndex Index
	newIndex.Key = newContact.Name
	newIndex.Position = newPosition

	tree.Insert(DataType(newIndex))
	indexes.indexes = append(indexes.indexes, newIndex)
}

func (contacts *Contacts) bulkWrite() {
	updatedContacts, err := json.MarshalIndent(contacts.contacts, "", "	")
	checkErr(err)
	err = ioutil.WriteFile("./data/contacts.json", updatedContacts, 0644)
	checkErr(err)
}

func (indexes *Indexes) bulkWrite() {
	updatedIndexes, err := json.MarshalIndent(indexes.indexes, "", "	")
	checkErr(err)
	err = ioutil.WriteFile("./data/index.json", updatedIndexes, 0644)
	checkErr(err)
}

func (indexes *Indexes) loadIndex(tree *BTree) {
	data, err := ioutil.ReadFile("./data/index.json")
	checkErr(err)
	err = json.Unmarshal(data, &indexes.indexes)

	for i := 0; i < len(indexes.indexes); i++ {
		tree.Insert(DataType(indexes.indexes[i]))
	}
}

func main() {
	// Create data
	// Insert data in a file
	// Create file index

	tree := Init()
	var contacts Contacts
	contacts.loadContacts()

	var indexes Indexes
	indexes.loadIndex(tree)

	for {
		var choice int
		fmt.Println("==============================")
		fmt.Println("Go Lang Schedule - with B Tree")
		fmt.Println("==============================")
		fmt.Println("(1) Create a new contact")
		fmt.Println("(2) Search a contact")
		fmt.Println("(3) View all contacts")
		fmt.Println("(4) View tree of indexes")

		fmt.Println("(0) Exit")

		fmt.Scanf("%d", &choice)

		if choice == 1 {
			contacts.createNewContact(tree, &indexes)
		}
		if choice == 2 {
			Clear()
			var name string
			fmt.Println("Searching for which contact? Type the name:")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			name = scanner.Text()
			a := tree.Search(name)
			if a != nil {
				fmt.Println(contacts.contacts[a.Position])
			} else {
				fmt.Println("Contact does not exists")
			}
		}
		if choice == 3 {
			fmt.Println("View all contacts")
			contacts.viewAll()
		}
		if choice == 4 {
			tree.root.Print("", true)
		}
		if choice == 0 {
			fmt.Println("Saving and exiting...")
			contacts.bulkWrite()
			indexes.bulkWrite()
			os.Exit(0)
		}
	}

}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func Clear() {
	fmt.Print("\033[H\033[2J") // escape codes para limpar a tela (Unix)
}
