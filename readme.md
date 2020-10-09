# Aplicação com Golang + gRPC

Desenvolvimento de aplicação utilizando Golang & gRPC .<br>
Exemplo de client e server.

## Tecnologias Utilizadas 🚀
* **[Visual Studio Code](https://code.visualstudio.com/)**
* **[Golang - Visual Studio Code (Extensão)](https://code.visualstudio.com/docs/languages/go)**
* **[Golang](https://golang.org/)**
* **[gRPC](https://grpc.io/docs/languages/go/basics/)**

## Instalando o Protoc 
1. Para que possamos gerar os arquivos com a extensão `.pb.go`, precisamos instalar o protoc. (Exemplo de instação no Windows 10 com chocolatey)
```
> choco install protoc
```
## Como Executar? 🔥
1. Para executar o `server.go`, devemos utilizar o seguinte comando: 
```
> go run server.go
```
2. Para executar o `client.go`, entrar na pasta: `/clients` e executar:
```
> go run client.go
```

Gerar arquivo pb
```
> protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:pb
```