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

