# Desafio Clean Architecture GO

Para este desafio, você precisará criar o usecase de listagem das orders.
Esta listagem precisa ser feita com:
- Endpoint REST (GET /order)
- Service ListOrders com GRPC
- Query ListOrders GraphQL
Não esqueça de criar as migrações necessárias e o arquivo api.http com a request para criar e listar as orders.

Para a criação do banco de dados, utilize o Docker (Dockerfile / docker-compose.yaml), com isso ao rodar o comando docker compose up tudo deverá subir, preparando o banco de dados.
Inclua um README.md com os passos a serem executados no desafio e a porta em que a aplicação deverá responder em cada serviço.


Passo a passo para rodar a aplicação:

1º) Executar o comando: "docker compose up -d"

2º) Para criar a tabela no banco de dados executar os comandos: "cd cmd/initdb; go run main.go"

3º) Para executar a aplicação executar os comandos: "cd cmd/ordersystem; go run main.go wire_gen.go"
    Os serviços iniciarão em:
        Starting web server on port :8000
        Starting gRPC server on port 50051
        Starting GraphQL server on port 8081

4º) Para testar:
    4.1) web server:
         Na pasta api/create_order.http selecione o "Send Request"
         Na pasta api/list_order.http selecione o "Send Request"
         [image: ListOrderHttpRequest.png]
    4.2) gRPC:
        Executar os comandos abaixo no terminal:
        -para startar o evans => evans -r repl
        -para setar o pacote => package pb
        -para setar o serviço => service Order
        -para chamar o método => call ListOrder
        [image: ListOrderGRCP.png]
    4.3) Graph:
        abrir o navegador em localhost:8081/graphql
        Colar e executar os comandos:
        mutation createOrder{
            createOrder(input: {id: "2", Price: 100, Tax: 10})
            {
                id,
                Price,
                Tax,
                FinalPrice
            }
            }

            query listOrder{
            listOrder{
                id,
                Price,
                Tax,
                FinalPrice
            }
            }

        [image: ListOrderGraphQL.png]
