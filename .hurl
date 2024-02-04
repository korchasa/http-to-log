# GET with query
GET http://127.0.0.1:8080/get?aaa=bbb

# JSON body
POST http://127.0.0.1:8080/json_body
x-app: MY_APP
{
  "type": "FOO",
  "value": 356789,
  "ordered": true,
  "index": 10
}

# Form
PUT http://127.0.0.1:8080/form
x-app: MY_APP
[FormParams]
default: false
email: john.doe@rookie.org
number: 33611223344

# Basic Auth
DELETE http://127.0.0.1:8080/basic_auth
[BasicAuth]
bob: secret

# Multipart Form
POST http://127.0.0.1:8080/multipart
[MultipartFormData]
field1: value1
field2: file,go.mod;
field3: file,main.go; application/zip
