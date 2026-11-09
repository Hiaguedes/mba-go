# Sobre modulos privados no golang

```
go env | grep PRIVATE
```

Pra ver o que esta na variavel de ambiente golang GOPRIVATE

nele voce coloca os caminhos dos repos que sao privados

pra setar faca

```
export GOPRIVATE=<path>
```

voce vai fazer a autenticacao no repo e depois rodar o mod tidy

```
go mod tidy
```

## Sobre a autenticacao

Pode ser por ssh ou cli

Pode editar o arquivo netrc do linux (pra autenticacao via http)

e colocar

```
machine github.com
login <user>
password <github token gerado no pat na pagina do github>
```

pra autenticacao pelo ssh voce mexe no gitconfig dentro .git do seu projeto ou no geral

## Detalhes interessantes que e interessante saber

```
ls ~/go/pkg/mod
```

voce consegue ver o cache de todos os pacotes ja baixados na maquina

Se o github cair existe um proxy chamado [golang proxy](https://proxy.golang.org/) porem nem todos os pacotes podem estar ai

pra isso voce pode ter as dependencias com voce em caso de qualquer infortunio, pra isso tem o 

```
go mod vendor
```

com isso crio uma especie de node_modules de todos os pacotes que estou utilizando
