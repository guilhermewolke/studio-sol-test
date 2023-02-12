# studio-sol-test
Teste prático do processo seletivo do Studio Sol

## Descrição
O teste prático para desenvolvedor Backend do Studio Sol consiste em 2 exercícios:

### API REST
Desenvolver uma API capaz de receber uma requisição via POST, com uma senha, e com um conjunto de regras de validação também enviadas no payload, realizar a validação da senha baseada
nos critérios deste conjunto de regras, e devolver um JSON com o resultado, no padrão REST

### API GraphQL
Desenvolver uma API capaz de receber uma requisição via POST, com uma senha, e com um conjunto de regras de validação também enviadas no payload, realizar a validação da senha baseada
nos critérios deste conjunto de regras, e devolver um JSON com o resultado, no padrão GraphQL

## Modo de usar
### API Rest
- Navegar até o diretório __rest__ e executar o arquivo __main.go__:

    `(studiosol)$ cd rest`

    `(rest)$ go run main.go`

- Executar via Postman, uma requisição POST para o endereço http://localhost:8080/verify, com o seguinte formato:
    ```
    {
        "password": "TesteSenhaForte!123&",
        "rules": [
            {
                "rule": "minSize",
                "value": 8
            },
            {
                "rule": "minSpecialChars",
                "value": 2
            },
            {
                "rule": "noRepeted",
                "value": 0
            },
            {
                "rule": "minDigit",
                "value": 4
            }
        ]
    }
    ```

### GraphQL
- Navegar até o diretório __graphql__ e executar o arquivo __main.go__:

    `(studiosol)$ cd graphql`
    
    `(graphql)$ go run main.go`

- Acessar http://localhost:8080/graphql
- Executar a seguinte query no "Playground":
    ```
        query{
            verify(input: {password:"TesteSenhaForte!123&", rules: [
                {rule: "minSize",value: 8},
                {rule: "minSpecialChars",value: 2},
                {rule: "noRepeted",value: 0},
                    {rule: "minDigit",value: 3}
            ]}) {
                verify
                noMatch
            }
        }
    ```
