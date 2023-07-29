package main

import (
	"bufio"
	"fmt"
	"os"
)

const MAX_NAME = 30
const MAX_ADDRESS = 50
const MAX_PHONE = 15

var lastInserted int

type Contact struct {
	name      string
	address   string
	phone     string
	isDeleted uint8
}

type Index struct {
	key      string
	position int
	size     int
}

func createContact() {
	Clear()
	fmt.Println("Creating a new contact!")
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

	if len(pAddress) > MAX_ADDRESS {
		pAddress = pAddress[:MAX_ADDRESS]
	}

	if len(pPhone) > MAX_PHONE {
		pPhone = pPhone[:MAX_PHONE]
	}
	newContact.name = pName
	newContact.address = pAddress
	newContact.phone = pPhone
	newContact.isDeleted = 2

	insertContactInFile(newContact)

}

func insertContactInFile(contact Contact) {
	f, _error := os.OpenFile("./data/contacts.data", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	checkErr(_error)
	defer f.Close()
	var index Index
	nameBytes, _ := f.WriteString(contact.name)
	pipe1, _ := f.WriteString("|")
	addressBytes, _ := f.WriteString(contact.address)
	pipe2, _ := f.WriteString("|")
	phoneBytes, _ := f.WriteString(contact.phone)
	pipe3, _ := f.WriteString("|")
	isDeletedBytes, _ := f.Write([]byte(string(contact.isDeleted)))
	totalBytes := nameBytes + addressBytes + phoneBytes + isDeletedBytes + pipe1 + pipe2 + pipe3
	Clear()
	index.size = totalBytes
	index.position = lastInserted
	index.key = contact.name
	lastInserted += totalBytes
	fmt.Printf("Contact created with %d bytes at %d position.\n", index.size, index.position)
}

func getFromFile(pos int, length int) {
	data, err := os.Open("./data/contacts.data")
	checkErr(err)
	byteSlice := make([]byte, length)
	data.ReadAt(byteSlice, int64(pos))
	fmt.Println(byteSlice)

	var contact Contact

	charName := []rune{}
	charAddress := []rune{}
	charPhone := []rune{}

	i := 0
	for i = 0; i < length; i++ {
		if byteSlice[i] == '|' {
			break
		}

		charName = append(charName, rune(byteSlice[i]))
	}
	i++
	j := i
	for j = i; j < length; j++ {
		if byteSlice[j] == '|' {
			break
		}

		charAddress = append(charAddress, rune(byteSlice[j]))

	}
	j++
	k := j
	for k = j; k < length; k++ {
		if byteSlice[k] == '|' {
			break
		}
		charPhone = append(charPhone, rune(byteSlice[k]))

	}
	k++
	contact.isDeleted = byteSlice[k]
	contact.name = string(charName)
	contact.address = string(charAddress)
	contact.phone = string(charPhone)

	fmt.Println(contact.name)
	fmt.Println(contact.address)
	fmt.Println(contact.phone)
}

func main() {
	// Create data
	// Insert data in a file
	// Create file index

	fileInfo, err := os.Stat("./data/contacts.data")
	checkErr(err)
	lastInserted = int(fileInfo.Size())
	for {
		var choice int
		fmt.Println("==============================")
		fmt.Println("Go Lang Schedule - with B Tree")
		fmt.Println("==============================")
		fmt.Println("(1) Create a new contact")
		fmt.Println("(2) View a contact")

		fmt.Scanf("%d", &choice)

		if choice == 1 {
			createContact()
		}
		if choice == 2 {
			fmt.Println("View contact")
			var pos int
			var length int
			fmt.Scanf("%d %d", &pos, &length)
			getFromFile(pos, length)
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
