# liqi-challenge


## Running and Using Locally
    In project root
        - $ docker-compose up -d 
            or 
        - $ go run cmd/server/main.go
        -
        
## Pratical tests
    Get public and private key
        - make a "Get" request to "http://localhost:8000/api/keys

    Get address
        - make a "Post" request to "http://localhost:8000/api/address 
        - use the public key from the get keys response
        - use a json like this:
            
            {"public_key": "string_public_key"}
        
## Run tests
    In project root
        $ go test ./...

    the file "sqs_service_test.go" contains the test:
        - TestSendMsgWithSQSSenderSucess()
    wich send the the value to SQS AWS service:
        - {"to":"0xAbC1234567", "value":"0x1bc16d674ec80000"}

## Comments
    Ao longo nas estapas finais do desenvolvimento tive um problema com o serviço Lambda da AWS o que me empediu de seguir desenvolvendo para o mesmo, e tambem tive com alguns dúvidas no que diz respeito ao segundo desafio, segue prints na pasta "AWSProblem" na raiz do projeto.

    P.S.: O serviço SQS continua "up" para a execução de alguns testes
