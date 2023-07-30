package main

import (
	"bufio"
	"fmt"
	"os"
)

const MAX_NAME = 30
const MAX_ADDRESS = 50
const MAX_PHONE = 15
const LENGTH = 99

var lastInserted int

type Index struct {
	key      string
	position int
}

func main() {

	fileInfo, err := os.Stat("../data/contacts.data")
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
		fmt.Println("(2) Search a contact")
		fmt.Println("(3) View index tree")
		fmt.Println("(4) View contacts")
		fmt.Println("(5) Edit a contact")
		fmt.Println("(6) Remove a contact")
		fmt.Println("(7) Retrieve trash")
		fmt.Println("(8) Empty trash")
		fmt.Println("(0) Exit")

		fmt.Scanf("%d", &choice)

		if choice == 1 {
			tree.createContact()
		}
		if choice == 2 {
			Clear()
			fmt.Println("View contact")
			scanner := bufio.NewScanner(os.Stdin)
			var name string
			scanner.Scan()
			name = scanner.Text()
			find := tree.Search(name)
			if find != nil {
				contact := getContactFromFile(find.position)
				contact.printContact()
			} else {
				fmt.Println("Could not find!!!")
			}
		}
		if choice == 3 {
			fmt.Println("View all contacts")
			tree.root.Print(" ", true)
		}
		if choice == 4 {
			Clear()
			fmt.Println("View all contacts")
			if len(tree.root.keys) == 0 {
				fmt.Println("There are no contacts inside.")
			} else {
				tree.root.PrintContacts()
			}
			Menu()
		}
		if choice == 5 {
			Clear()
			fmt.Println("Which contact would you like to edit? [type the name]")
			scanner := bufio.NewScanner(os.Stdin)
			var name string
			scanner.Scan()
			name = scanner.Text()
			find := tree.Search(name)
			if find != nil {
				contact := getContactFromFile(find.position)
				contact.removeDolar()
				newKey := contact.editInfo(find.key, find.position, tree)
				find.key = newKey
			} else {
				fmt.Println("Could not find!")
			}
		}
		if choice == 6 {
			Clear()
			fmt.Println("Which contact would you like to remove? [type the name]")
			scanner := bufio.NewScanner(os.Stdin)
			var name string
			scanner.Scan()
			name = scanner.Text()
			find := tree.Search(name)
			if find != nil {
				contact := getContactFromFile(find.position)
				contact.delete(find.key, find.position, tree)
			} else {
				fmt.Println("Name not found. Contact not deleted.")
			}
		}
		if choice == 7 {
			Clear()
			retrieveFromTrash(tree)
		}
		if choice == 0 {
			tree.bulkWrite()
			os.Exit(0)
		}
	}
}
