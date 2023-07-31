package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Contact struct {
	name      string
	address   string
	phone     string
	isDeleted rune
}

func createContactObject() (*Contact, string) {
	var newContact Contact

	scanner := bufio.NewScanner(os.Stdin)

	// previous varnames to validade before create
	var pName string
	var pAddress string
	var pPhone string

	fmt.Println("Nome (tamanho 30):")
	scanner.Scan()
	pName = scanner.Text()
	for len(pName) == 0 {
		fmt.Println("Nome é obrigatório!")
		fmt.Println("Nome (tamanho 30):")
		scanner.Scan()
		pName = scanner.Text()
	}

	fmt.Println("Endereço (tamanho 50):")
	scanner.Scan()
	pAddress = scanner.Text()

	fmt.Println("Telefone (tamanho 15):")
	scanner.Scan()
	pPhone = scanner.Text()

	// VALIDATORS
	if len(pName) > MAX_NAME {
		pName = pName[:MAX_NAME]
	}
	var keyName string
	keyName = pName

	if len(pAddress) > MAX_ADDRESS {
		pAddress = pAddress[:MAX_ADDRESS]
	}

	if len(pPhone) > MAX_PHONE {
		pPhone = pPhone[:MAX_PHONE]
	}

	newContact.name = pName
	newContact.address = pAddress
	newContact.phone = pPhone
	newContact.insertDolar()
	newContact.isDeleted = '0'

	return &newContact, keyName
}

func (tree *BTree) createContact() {
	Clear()
	fmt.Println("Criando um novo contato.")

	newContact, keyName := createContactObject()
	if tree.Search(keyName) != nil {
		Clear()
		fmt.Println("Name already exists! Create a new one.")
		Menu()
		return
	}

	newIndex := insertContactInFile(*newContact)

	var newIndexSolid Index
	newIndexSolid.key = keyName
	newIndexSolid.position = newIndex.position
	tree.Insert(DataType(newIndexSolid))
}

func (contact *Contact) editInfo(key string, position int, tree *BTree) string {
	newContact, keyName := createContactObject()
	editContactInFile(*newContact, position)
	return keyName
}

func (contact *Contact) delete(key string, pos int, tree *BTree) {
	contact.isDeleted = '1'
	editContactInFile(*contact, pos)
	tree.root.Delete(key)
}

func (contact *Contact) retrieve(pos int, tree *BTree) {
	contact.isDeleted = '0'
	editContactInFile(*contact, pos)

	var index Index
	contact.removeDolar()

	index.key = contact.name
	index.position = pos
	tree.root.Insert(DataType(index))
}

func getAndPrintContact(index *Index) {
	contact := getContactFromFile(index.position)
	contact.removeDolar()
	contact.printContact()
}

func (contact *Contact) printContact() {
	fmt.Println("------------------------------------")
	fmt.Println("NOME:       ", contact.name)
	fmt.Println("ENDEREÇO:   ", contact.address)
	fmt.Println("TELEFONE:   ", contact.phone)
	fmt.Println("------------------------------------")
}

func (contact *Contact) removeDolar() {

	nonCharCount := strings.Count(contact.name, "$")
	contact.name = contact.name[:MAX_NAME-nonCharCount]

	nonCharCount = strings.Count(contact.address, "$")
	contact.address = contact.address[:MAX_ADDRESS-nonCharCount]

	nonCharCount = strings.Count(contact.phone, "$")
	contact.phone = contact.phone[:MAX_PHONE-nonCharCount]
}

func (contact *Contact) insertDolar() {

	if len(contact.name) < MAX_NAME {
		s := make([]byte, MAX_NAME)
		for i := 0; i < len(contact.name); i++ {
			s[i] = contact.name[i]
		}
		for i := len(contact.name); i < MAX_NAME; i++ {
			s[i] = '$'
		}
		contact.name = string(s)
	}

	if len(contact.address) < MAX_ADDRESS {
		s := make([]byte, MAX_ADDRESS)
		for i := 0; i < len(contact.address); i++ {
			s[i] = contact.address[i]
		}
		for i := len(contact.address); i < MAX_ADDRESS; i++ {
			s[i] = '$'
		}
		contact.address = string(s)
	}

	if len(contact.phone) < MAX_PHONE {
		s := make([]byte, MAX_PHONE)
		for i := 0; i < len(contact.phone); i++ {
			s[i] = contact.phone[i]
		}
		for i := len(contact.phone); i < MAX_PHONE; i++ {
			s[i] = '$'
		}
		contact.phone = string(s)
	}
}
