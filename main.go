package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MAX_NAME = 30
const MAX_ADDRESS = 50
const MAX_PHONE = 15

var lastInserted int

type Contact struct {
	name      string `json:"name"`
	address   string `json:"address"`
	phone     string `json:"phone"`
	isDeleted uint8
}

type Index struct {
	key      string
	position int
	size     int
}

func (tree *BTree) createContact() {
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

	newIndex := insertContactInFile(newContact)
	var newIndexSolid Index
	newIndexSolid.key = newIndex.key
	newIndexSolid.position = newIndex.position
	newIndexSolid.size = newIndex.size
	tree.Insert(DataType(newIndexSolid))
	insertIndexInFile(newIndex)
}

func insertContactInFile(contact Contact) *Index {
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

	return &index
}

func insertIndexInFile(index *Index) {
	f, _error := os.OpenFile("./data/index.data", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	checkErr(_error)
	defer f.Close()

	// Write index data in file, separated by |
	f.WriteString(index.key + "\n")
	f.WriteString(fmt.Sprint(index.position) + "\n")
	f.WriteString(fmt.Sprint(index.size) + "\n")
}

func getContactsFromFile(pos int, length int) {
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

func (tree *BTree) loadIndexes() {
	indexFile, _err := os.Open("./data/index.data")
	checkErr(_err)
	fileScanner := bufio.NewScanner(indexFile)
	var pKey string
	var pPosition int
	var pSize int

	for {
		var newIndex Index
		fileScanner.Scan()

		if fileScanner.Text() == "" {
			break
		}

		pKey = fileScanner.Text()

		fileScanner.Scan()
		pPosition, _ = strconv.Atoi(fileScanner.Text())

		fileScanner.Scan()
		pSize, _ = strconv.Atoi(fileScanner.Text())

		newIndex.key = pKey
		newIndex.position = pPosition
		newIndex.size = pSize

		tree.Insert(DataType(newIndex))
	}

}

func main() {
	// Create data
	// Insert data in a file
	// Create file index

	fileInfo, err := os.Stat("./data/contacts.data")
	checkErr(err)
	lastInserted = int(fileInfo.Size())

	tree := Init()
	tree.loadIndexes()

	for {
		var choice int
		fmt.Println("==============================")
		fmt.Println("Go Lang Schedule - with B Tree")
		fmt.Println("==============================")
		fmt.Println("(1) Create a new contact")
		fmt.Println("(2) View a contact")
		fmt.Println("(3) View all contacts")

		fmt.Scanf("%d", &choice)

		if choice == 1 {
			tree.createContact()
		}
		if choice == 2 {
			fmt.Println("View contact")
			var pos int
			var length int
			fmt.Scanf("%d %d", &pos, &length)
			getContactsFromFile(pos, length)
		}
		if choice == 3 {
			fmt.Println("View all contacts")
			tree.root.Print(" ", true)
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
