## Dashboard Service

TBD

### Usage

The following steps can be used in this case:

1. Authenticate the subject and get the JWT.<br/>
   - See all the details in the sibling `iam_svc`.
   - Ex: `curl -i -d '{ "email":"john@doe.com", "password":"pass1234" }' localhost:3001/v1/authenticate`
   - To get the `id` of the subject you can either:
     - Decode the received `access_token` at [jwt.io](https://jwt.io) and get the value of `sub` claim;
     - Get the subject attributes from the same `iam_svc`:<br/>
       - Assuming that this id is `9f83a904-152a-4a5a-8c4e-6cf78dfbb399`
       - Use `curl localhost:3001/v1/subjects/9f83a904-152a-4a5a-8c4e-6cf78dfbb399/attributes`<br/>
2. Get the subject's portfolio (of products)
   - Get the JWT<br/>
     Ex: `export JWT="eyJhbGciOiJFUzI1NiJ9.eyJpc3MiOiJpYW0uc2VydmljZSIsInN1YiI6IjlmODNhOTA0LTE1MmEtNGE1YS04YzRlLTZjZjc4ZGZiYjM5OSIsImF1ZCI6WyJhbnlvbmUiXSwiZXhwIjoxNjIzMzQ3MzE2LjgyOTAxNTcsIm5iZiI6MTYyMzM0MzcxNi44MjkwMTU3LCJpYXQiOjE2MjMzNDM3MTYuODI5MDE1N30.gZuK3XWrg-dDRvJt_4hIN1o4LI_7m-UEIyCMdwWdXfnNGvZ0JmBrBOnA4___1WlK7MUydWsyC44Y3ikORhSeIQ"`
   - Use it while calling the protected resource.<br/>
     Ex: `curl -i -H "Authorization: Bearer $JWT" http://localhost:3002/v1/subjects/9f83a904-152a-4a5a-8c4e-6cf78dfbb399/portfolio`
     - An _successful_ (HTTP status code `200`) and empty response will be retrieved, if all good.
     - A _forbidden_ (HTTP status code `403`), otherwise.

<br/>

### Design Choices

- Each app/service has its (unique within the deployment) name, and an associated set of policies.
- The entry points where authorization needs to happen, also called *Policy Enforcement Point*s, are using the `AuthzFacade` component to ask for the authorization decision.
