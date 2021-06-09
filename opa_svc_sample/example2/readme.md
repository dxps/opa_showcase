### Example 2: Evaluation is just based on `input`

In this example we have just one `Policy` (with a `rule` in it) and one `Query`.<br/>

The logic in the rule uses just the provided `input` and no any other `data` elements.

<br/>

#### Usage

You have all done in three simple steps:

1. Start the OPA in server mode (same as in a sidecar approach).

   - Use `./run_opa.sh` script or run `opa run --server`.

1. Upload the `Policy`.

   - Use `./upload_policy.sh` script or run:<br/>
     `curl -X PUT localhost:8181/v1/policies/products --data-binary @products_policy_ex2.rego`

1. `Query` for authorization decisions.

   - Use `./query_authz.sh` script or run:<br/>
     `curl -X POST localhost:8181/v1/data/products/policy/ex2/subject_has_product --data-binary @query_input.json`<br/>

   - The response must be `{"result":true}`.

   - By changing the _current_ (`context.product`) product to a non-existent value (such as `product_3`) in `query_input.json` file, querying again using `./query_authz.sh` should return a negative answer, that is `{"result":false}`
