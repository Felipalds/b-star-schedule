package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func (tree *BTree) loadIndexes() {
	indexFile, _err := os.Open("../data/index.data")
	checkErr(_err)
	fileScanner := bufio.NewScanner(indexFile)
	var pKey string
	var pPosition int

	for {
		var newIndex Index
		fileScanner.Scan()

		if fileScanner.Text() == "" {
			break
		}

		pKey = fileScanner.Text()

		fileScanner.Scan()
		pPosition, _ = strconv.Atoi(fileScanner.Text())

		newIndex.key = pKey
		newIndex.position = pPosition

		tree.Insert(DataType(newIndex))
	}
}

func getContactFromFile(pos int) *Contact {
	data, err := os.Open("../data/contacts.data")
	checkErr(err)
	byteSlice := make([]byte, LENGTH)
	data.ReadAt(byteSlice, int64(pos))
	var contact Contact

	charName := []rune{}
	charAddress := []rune{}
	charPhone := []rune{}

	i := 0
	for i = 0; i < LENGTH; i++ {
		if byteSlice[i] == '|' {
			break
		}

		charName = append(charName, rune(byteSlice[i]))
	}
	i++
	j := i
	for j = i; j < LENGTH; j++ {
		if byteSlice[j] == '|' {
			break
		}

		charAddress = append(charAddress, rune(byteSlice[j]))

	}
	j++ // |
	k := j
	for k = j; k < LENGTH; k++ {
		if byteSlice[k] == '|' {
			break
		}
		charPhone = append(charPhone, rune(byteSlice[k]))
	}
	k++
	contact.isDeleted = rune(byteSlice[k])
	contact.name = string(charName)
	contact.address = string(charAddress)
	contact.phone = string(charPhone)

	return &contact
}

func editContactInFile(contact Contact, position int) *Index {
	f, _error := os.OpenFile("../data/contacts.data", os.O_RDWR, 0666)
	f.Seek(int64(position), 0)
	checkErr(_error)
	defer f.Close()
	var index Index
	f.WriteString(contact.name)
	f.WriteString("|")
	f.WriteString(contact.address)
	f.WriteString("|")
	f.WriteString(contact.phone)
	f.WriteString("|")
	f.Write([]byte(string(contact.isDeleted)))

	Clear()

	index.position = lastInserted
	index.key = contact.name

	fmt.Printf("Contact updated at %d position.\n", index.position)

	return &index
}

func insertContactInFile(contact Contact) *Index {
	f, _error := os.OpenFile("../data/contacts.data", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
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

	index.position = lastInserted
	index.key = contact.name

	lastInserted += totalBytes
	fmt.Printf("Contact created at %d position.\n", index.position)

	return &index
}

func insertContactInSecondaryFile(contact Contact, secondaryIndex *int) *Index {
	f, _error := os.OpenFile("../data/contacts-2.data", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	checkErr(_error)
	defer f.Close()
	var index Index
	f.WriteString(contact.name)
	f.WriteString("|")
	f.WriteString(contact.address)
	f.WriteString("|")
	f.WriteString(contact.phone)
	f.WriteString("|")
	f.Write([]byte(string(contact.isDeleted)))
	Clear()

	contact.removeDolar()
	index.position = *secondaryIndex
	*secondaryIndex += int(LENGTH)
	index.key = contact.name

	fmt.Printf("Contact created at %d position.\n", index.position)

	return &index
}

func insertIndexInSecondFile(index *Index) {

	f, _error := os.OpenFile("../data/index-2.data", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	checkErr(_error)
	defer f.Close()

	// Write index data in file, separated by |
	f.WriteString(index.key + "\n")
	f.WriteString(fmt.Sprint(index.position) + "\n")
}

func insertIndexInFile(index *Index) {

	f, _error := os.OpenFile("../data/index.data", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	checkErr(_error)
	defer f.Close()

	// Write index data in file, separated by |
	f.WriteString(index.key + "\n")
	f.WriteString(fmt.Sprint(index.position) + "\n")
}

func (tree *BTree) bulkWrite() {
	f, _error := os.Create("../data/index.data")
	checkErr(_error)
	f.Close()
	tree.root.VisitInOrder()
}

func retrieveFromTrash(tree *BTree) {
	pos := 0
	for {
		contact := getContactFromFile(pos)
		if contact.isDeleted == '1' {
			contact.retrieve(pos, tree)
		}

		pos += LENGTH

		if pos >= lastInserted {
			return
		}
	}
}

func deleteAndReindex(tree *BTree) *BTree {
	tree.bulkWrite()
	newTree := Init()
	newFileIndex := 0
	pos := 0
	for {
		contact := getContactFromFile(pos)
		if contact.isDeleted == '0' {
			index := insertContactInSecondaryFile(*contact, &newFileIndex)
			insertIndexInSecondFile(index)
		}

		pos += LENGTH

		if pos >= lastInserted {
			break
		}
	}

	os.Remove("../data/contacts.data")
	os.Remove("../data/index.data")
	os.Rename("../data/index-2.data", "../data/index.data")
	os.Rename("../data/contacts-2.data", "../data/contacts.data")

	newTree.loadIndexes()
	return newTree
}
