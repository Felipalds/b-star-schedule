pontos de melhoria: 
manter arquivo de contatos aberto
mudar para json
inserir como linhas


When a struct field is capitalized, it is considered a public field, meaning it can be accessed and modified from outside the package in which the struct is defined. On the other hand, if a struct field is uncapitalized, it is considered private and can only be accessed and modified within the same package.
For JSON unmarshalling to work correctly in Go, the fields you want to map from the JSON data must be exported (i.e., capitalized) because the encoding/json package uses reflection to access and set the struct fields. If the fields are not exported, the json.Unmarshal() function will not be able to access them, and the unmarshalling will fail.