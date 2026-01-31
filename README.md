# AulaArrayFatiaMapa

Projeto didático em Go que demonstra o uso de Arrays, Slices (fatias) e Maps.

## Descrição

Este repositório contém exemplos simples organizados em pacotes para ilustrar conceitos básicos de coleções em Go:

- `array` — exemplos com arrays
- `fatia` — exemplos com slices (fatias)
- `mapa` — exemplos com mapas (map)
- `pacoteprincipal` — executável com um menu que chama os exemplos

## Requisitos

- Go 1.18+ instalado

## Estrutura do repositório

- array/
- fatia/
- mapa/
- pacoteprincipal/
- go.mod

## Como executar

Abra um terminal na raiz do projeto e execute:

```powershell
go run ./pacoteprincipal
```

O programa exibirá um menu interativo; escolha a opção desejada (1, 2, 3 ou 5 para sair).

Alternativamente, para compilar um binário:

```powershell
go build -o bin/menu ./pacoteprincipal
./bin/menu
```

## O que cada pacote faz

- `array`: função `Array()` demonstra manipulação de arrays.
- `fatia`: função `Fatia()` demonstra slices (criação, append, slicing).
- `mapa`: função `Mapa()` demonstra uso de mapas (inserção, leitura, iteração).
- `pacoteprincipal/menu.go`: contém o `main()` e o menu interativo que chama os exemplos.

## Contribuição

Melhorias e correções são bem-vindas. Abra uma issue ou pull request com a sugestão.

## Licença

Sem licença definida — use conforme necessário ou adicione uma licença apropriada.
