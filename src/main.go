/*CRIADO POR: LUIZ FELIPE E PEDRO HENRIQUE ZOZ*/

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

	check1, err := os.Open("../data/contacts.data")
	if err != nil {
		os.Create("../data/contacts.data")
	}
	check1.Close()

	check2, err := os.Open("../data/index.data")
	if err != nil {
		os.Create("../data/index.data")
	}
	check2.Close()

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
		fmt.Println("(1) Criar um contato")
		fmt.Println("(2) Buscar um contato")
		fmt.Println("(3) Ver a árvore de índices em memória")
		fmt.Println("(4) View todos os contatos")
		fmt.Println("(5) Editar um contato")
		fmt.Println("(6) Remover um contato")
		fmt.Println("(7) Recuperar contatos da lixeira")
		fmt.Println("(8) Limpar a lixeira")
		fmt.Println("(0) Salvar e sair")

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
				contact.removeDolar()
				contact.printContact()
			} else {
				fmt.Println("Contato não encontrado.")
			}
		}
		if choice == 3 {
			fmt.Println("Buscando todos os contatos")
			tree.root.Print(" ", true)
		}
		if choice == 4 {
			Clear()
			fmt.Println("Ver todos os contatos")
			if len(tree.root.keys) == 0 {
				fmt.Println("Não há contatos para serem exibidos")
			} else {
				tree.root.PrintContacts()
			}
			Menu()
		}
		if choice == 5 {
			Clear()
			fmt.Println("Qual contato você gostaria de editar? [Digite o nome]")
			scanner := bufio.NewScanner(os.Stdin)
			var name string
			scanner.Scan()
			name = scanner.Text()
			find := tree.Search(name)
			var newIndex Index
			if find != nil {
				add := find.position
				contact := getContactFromFile(find.position)
				contact.removeDolar()
				newKey := contact.editInfo(find.key, find.position, tree)
				tree.root.Delete(find.key)
				newIndex.key = newKey
				newIndex.position = add
				tree.Insert(DataType(newIndex))
			} else {
				fmt.Println("Não foi possível encontrar.")
			}
		}
		if choice == 6 {
			Clear()
			fmt.Println("Qual contato você gostaria de remover? [Digite o nome]")
			scanner := bufio.NewScanner(os.Stdin)
			var name string
			scanner.Scan()
			name = scanner.Text()
			find := tree.Search(name)
			if find != nil {
				contact := getContactFromFile(find.position)
				contact.delete(find.key, find.position, tree)
			} else {
				fmt.Println("Contato não encontrado e não apagado.")
			}
		}
		if choice == 7 {
			Clear()
			retrieveFromTrash(tree)
		}
		if choice == 8 {
			Clear()
			tree = deleteAndReindex(tree)
			fileInfo, err := os.Stat("../data/contacts.data")
			checkErr(err)
			lastInserted = int(fileInfo.Size())
		}
		if choice == 0 {
			tree.bulkWrite()
			os.Exit(0)
		}
	}
}
