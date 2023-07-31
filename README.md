# B Tree Schedule
## Luiz Felipe e Pedro Henrique Zoz
## Como rodar o programa?
1. Entre na pasta ./src
2. Rode o comando ```go build```
3. Rode o executável criado ```./src```

## Especificações técnicas:
1. Ao rodar o programa, ele verificará se existem os arquivos de index e de contatos
- Caso não exista, ele criará.

2. Os índices do arquivo serão carregados para uma árvore B. Esta árvore é uma árvore de indexação e busca e possui os dados de chave e posição.
- A chave é o nome do contato
- A posição é a posição em bytes do contato no arquivo de contatos.

3. Ao criar um novo contato, validamos se o nome já não existe.
- Nome é obrigatório e deve possuir até 30 caracteres.
- Endereço é opcional (pode ser vazio) e deve possuir até 50 caracteres.
- Telefone é opcional (pode ser vazio) e deve possuir até 15 caracteres.
- Após criar o contato, criamos seu índice na árvore B e o inserimos no arquivo de dados.

4. Ao exibir os contatos, buscamos todos os contatos indexados na árvore, em um percurso Em Ordem. Então buscamos sua posição no arquivo, recuperamos os dados e escrevemos na tela.

5. Ao editar um contato, pedimos para digitar o nome. Então, damos uma busca na árvore e, ao achar, buscamos o contato no arquivo a partir do endereço. Pedimos ao usuário para escrever os dados novamente e então, atualizamos a árvore de indices com a nova chave.

6. Ao buscar um contato, buscamos pela árvore B.

7. Ao remover um contato, buscamos pela árvore B e então, marcamos o bit de exclusão como 1 no arquivo. Removemos então seu índice da árvore.

8. Ao recuperar contatos da lixeira, buscamos todos os contatos que tenham o bit marcado, flipamos o bit e inserimos o contato na árvore novamente.

9. Ao apagar toda a lixeira, criamos um novo arquivo de indice e um novo arquivo de dados. Movemos todos os dados não marcados como apagados para o novo arquivo e criamos a árvore novamente.

