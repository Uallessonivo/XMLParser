# Projeto de Conversão de Dados de XML com Mapeamento a partir de CSV

Este projeto é um programa em Go (Golang) que lê um arquivo XML, realiza operações de mapeamento usando um arquivo CSV chamado "base.csv" e, em seguida, gera um novo arquivo XML com os dados atualizados. O programa também faz a leitura do número da nota fiscal (nNF) no arquivo XML e utiliza esse número para nomear o novo arquivo XML gerado.

## Requisitos

- [Go (Golang)](https://golang.org/dl/): Certifique-se de ter o Go instalado no seu sistema.

## Funcionalidade

O programa lê o arquivo "base.csv" e cria um mapa associando códigos (chave) a valores. Em seguida, ele lê o arquivo XML e procura por elementos `<cProd>`. Se um código `<cProd>` corresponder a uma chave no mapa, o programa atualiza o valor desse elemento com o valor correspondente do mapa. Além disso, o programa renomeia o arquivo XML resultante com o número da nota fiscal (nNF) encontrado no arquivo de entrada.

## Como Usar

1. Certifique-se de que o arquivo "base.csv" esteja no mesmo diretório em que você executará o programa.

2. Compile e execute o programa:

   ```bash
   go run main.go
   ```

3. O programa solicitará que você digite o nome do arquivo XML que deseja processar. Certifique-se de que o arquivo XML de entrada esteja no mesmo diretório.

4. Se o nome do arquivo XML não tiver a extensão ".xml", o programa a adicionará automaticamente.

5. Após a execução bem-sucedida, o programa criará um novo arquivo XML com os dados atualizados e o nomeará com o número da nota fiscal (nNF) encontrado no arquivo XML de entrada.

## Estrutura do Projeto

- `main.go`: O código-fonte principal do programa.
- `base.csv`: O arquivo CSV que contém os dados de mapeamento (códigos e valores).

## Exemplo de Saída

Após a execução bem-sucedida do programa, você verá o conteúdo do arquivo XML resultante no console, e um novo arquivo XML será criado com o nome baseado no número da nota fiscal (nNF).

## Observações

- Certifique-se de que o arquivo "base.csv" está localizado no mesmo diretório em que o programa é executado.
- Os dados do arquivo XML são lidos com a biblioteca xmlquery, que é usada para processar o XML.
- O programa usa o número da nota fiscal (nNF) encontrado no arquivo XML para nomear o novo arquivo XML gerado.