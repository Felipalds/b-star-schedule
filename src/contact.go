package main

import (
	"bufio"
	"fmt"
	"os"
)

type Contact struct {
	name      string
	address   string
	phone     string
	isDeleted uint8
}

func createContactObject() (*Contact, string) {
	var newContact Contact

	scanner := bufio.NewScanner(os.Stdin)

	// previous varnames to validade before create
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

	// VALIDATORS
	if len(pName) > MAX_NAME {
		pName = pName[:MAX_NAME]
	}
	var keyName string
	keyName = pName

	if len(pName) < MAX_NAME {
		s := make([]byte, MAX_NAME)
		for i := 0; i < len(pName); i++ {
			s[i] = pName[i]
		}
		for i := len(pName); i < MAX_NAME; i++ {
			s[i] = '$'
		}
		pName = string(s)
	}

	if len(pAddress) > MAX_ADDRESS {
		pAddress = pAddress[:MAX_ADDRESS]
	}
	if len(pAddress) < MAX_ADDRESS {
		s := make([]byte, MAX_ADDRESS)
		for i := 0; i < len(pAddress); i++ {
			s[i] = pAddress[i]
		}
		for i := len(pAddress); i < MAX_ADDRESS; i++ {
			s[i] = '$'
		}
		pAddress = string(s)
	}

	if len(pPhone) > MAX_PHONE {
		pPhone = pPhone[:MAX_PHONE]
	}
	if len(pPhone) < MAX_PHONE {
		s := make([]byte, MAX_PHONE)
		for i := 0; i < len(pPhone); i++ {
			s[i] = pPhone[i]
		}
		for i := len(pPhone); i < MAX_PHONE; i++ {
			s[i] = '$'
		}
		pPhone = string(s)
	}

	newContact.name = pName
	newContact.address = pAddress
	newContact.phone = pPhone
	newContact.isDeleted = 1

	return &newContact, keyName
}

func (tree *BTree) createContact() {
	Clear()
	fmt.Println("Creating a new contact!")

	newContact, keyName := createContactObject()
	newIndex := insertContactInFile(*newContact)

	var newIndexSolid Index
	newIndexSolid.key = keyName
	newIndexSolid.position = newIndex.position
	newIndexSolid.size = newIndex.size
	tree.Insert(DataType(newIndexSolid))
}

func (contact *Contact) editInfo(key string, position int, size int, tree *BTree) string {
	newContact, keyName := createContactObject()
	editContactInFile(*newContact, position, size)
	return keyName
}

func (contact *Contact) printContact() {
	fmt.Println("NAME: ", contact.name)
	fmt.Println("ADDRESS: ", contact.address)
	fmt.Println("PHONE: ", contact.phone)
}

func (contact *Contact) removeDolar() {
	s := make([]byte, MAX_NAME)
	for i := 0; i < MAX_NAME; i++ {
		if contact.name[i] != '$' {
			s[i] = contact.name[i]
		}
	}
	contact.name = string(s)

	s = make([]byte, MAX_ADDRESS)
	for i := 0; i < MAX_ADDRESS; i++ {
		if contact.address[i] != '$' {
			s[i] = contact.address[i]
		}
	}
	contact.address = string(s)

	s = make([]byte, MAX_PHONE)
	for i := 0; i < MAX_PHONE; i++ {
		if contact.phone[i] != '$' {
			s[i] = contact.phone[i]
		}
	}
	contact.phone = string(s)
}
