GET http://localhost:3000/
HTTP 200

GET http://localhost:3000/aaa
[Options]
delay: 1s
HTTP 404

POST http://localhost:3000/_/graphql
```graphql
query {
  listInvocations {
    id
    status
    method
    url
  }
}
```
HTTP 200
[Asserts]
jsonpath "$.data.listInvocations" count == 2
jsonpath "$.data.listInvocations[0].url" == "https://example.com/aaa"
jsonpath "$.data.listInvocations[0].status" == 404
jsonpath "$.data.listInvocations[1].url" == "https://example.com/"
jsonpath "$.data.listInvocations[1].status" == 200
