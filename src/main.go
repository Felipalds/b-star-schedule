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

type Index struct {
	key      string
	position int
	size     int
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
		fmt.Println("(6) Exit")

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
				contact := getContactFromFile(find.position, find.size)
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
			fmt.Println("View all contacts")
			tree.root.Print(" ", true)
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
				contact := getContactFromFile(find.position, find.size)
				newKey := contact.editInfo(find.key, find.position, find.size, tree)
				find.key = newKey
			} else {
				fmt.Println("Could not find!!!")
			}
		}
		if choice == 6 {
			tree.bulkWrite()
			os.Exit(0)
		}
	}
}
