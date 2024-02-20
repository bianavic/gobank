# [WIP] JSON API in Golang

### tools
- go1.22.0
- goland ide
- linux
- gorilla/mux v1.8.1
- lib/pq v1.10.9
- golang-jwt/jwt/v5
- golang.org/x/crypto v0.19.0
- github.com/stretchr/testify
- postgres

### routes

| API ROUTE		            | DESCRIPTION                            | STATUS | PROTECTED |
|:-----------------------|:---------------------------------------|:-------|:----------|
| [GET] /account	        | Fetch all accounts                     | 200    | no        |
| [GET] /account/{id}    | Get an account by its ID               | 200    | yes       |
| [DELETE] /account/{id} | Delete an account by its ID            | 200    | yes       |
| [POST] /account        | Create a new account                   | 201    | no        |
| [POST] /login          | Login into your account                | 200    | no        |
| [POST] /transfer       | Transfer an amount to an accountNumber | 200    | no        |


## ACCOUNTS

#### get all accounts
```bash
curl --location 'http://localhost:3000/account'
```
###### 200 OK
``` json
[
    {
        "ID": 9,
        "firstName": "User Test1",
        "lastName": "UT1",
        "number": 963577,
        "balance": 0,
        "createdAt": "0001-01-01T00:00:00Z",
        "email": "user1@email.com"
    },
    {
        "ID": 10,
        "firstName": "User Test1",
        "lastName": "UT1",
        "number": 963577,
        "balance": 0,
        "createdAt": "0001-01-01T00:00:00Z",
        "email": "user1@email.com"
    },
    {
        "ID": 11,
        "firstName": "User Test2",
        "lastName": "UT2",
        "number": 782674,
        "balance": 0,
        "createdAt": "0001-01-01T00:00:00Z",
        "email": "user2@email.com"
    }
]
```

#### get an account by its ID
```bash
curl --location 'http://localhost:3000/account/10' \
--header 'x-jwt-token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50TnVtYmVyIjo5NjM1NzcsImVtYWlsIjoidXNlcjFAZW1haWwuY29tIiwiZXhwaXJlc0F0IjoxNTE2MjM5MDIyLCJpZCI6MH0.CLsKx-4dmyvSZwNIZVdPrUYbnNwpmk0RV0eP2IcSt2w'
```
###### 200 OK
``` json
{
    "ID": 10,
    "firstName": "User Test1",
    "lastName": "UT1",
    "number": 963577,
    "balance": 0,
    "createdAt": "0001-01-01T00:00:00Z",
    "email": "user1@email.com"
}
```

#### delete an account by its ID
```bash
curl --location --request DELETE 'http://localhost:3000/account/10' \
--header 'x-jwt-token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50TnVtYmVyIjo5NjM1NzcsImVtYWlsIjoidXNlcjFAZW1haWwuY29tIiwiZXhwaXJlc0F0IjoxNTE2MjM5MDIyLCJpZCI6MTB9.fXgASXwuP5QvfoTD2fiaP8E1JpNtVbHTsjTKX0K7oes'
```
###### 200 OK
``` json
{
    "deleted": 10
}
```

#### create a new account
```bash
curl --location 'http://localhost:3000/account' \
--header 'Content-Type: application/json' \
--data-raw '{
    "firstName": "User Test1",
    "lastName": "UT1",
     "email": "user1@email.com",
    "password": "test1"
}'
```
###### 200 OK
``` json
{
    "ID": 0,
    "firstName": "User Test1",
    "lastName": "UT1",
    "email": "user1@email.com",
    "number": 963577,
    "balance": 0,
    "createdAt": "2024-02-19T23:34:55.094602198Z"
}
```

## LOGIN

#### login into your account
```bash
curl --location 'http://localhost:3000/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "user1@email.com",
    "password": "test1"
}'
```
###### 200 OK
``` json
{
    "email": "user1@email.com",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50TnVtYmVyIjo5NjM1NzcsImVtYWlsIjoidXNlcjFAZW1haWwuY29tIiwiZXhwaXJlc0F0IjoxNTE2MjM5MDIyLCJpZCI6OX0.bba0mT7yltn-9AGL6s2H9fA73loeLMBGCC7aOc1vcM8"
}
```

## TRANSFER

#### transfer an amount to a specific account
```bash
curl --location 'http://localhost:3000/transfer' \
--header 'Content-Type: application/json' \
--data '{
    "toAccount": 963577,
    "amount": 100000
}'
```
###### 200 OK
``` json
{
    "toAccount": 963577,
    "amount": 100000
}
```