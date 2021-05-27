## Using OPA as a standalone/sidecar service

OPA can run in server mode and following the sidecar pattern it can be used for:

- uploading 1-N x `Policy` (one or multiple policies)
- feeding in `Data`
  - representing facts about external world (attributes of users, request/action, or target)
- doing a `Query Input` for getting authorization decisions

```
  policies & data mgmt         authorization decisions
 ----------------------       -------------------------

      .--------.
      | Policy |----------.
      '--------'          |
                          v
                     .---------.          .---------.
                     |   OPA   |<---------|  Query  |
                     '---------'          '---------'
                          ^
      .--------.          |
      |  Data  |----------'
      '--------'
```

### Usage

Follow these steps:

1. Start OPA in server mode.

- Use `opa run --server` or the provided `run_opa.sh` script.
- It will listen on port `8181` for HTTP requests.

1. Upload a `Policy`.

- Use `curl -X PUT localhost:8181/v1/data/products-acl -d @products-acl_data.json`<br>
  or the provided `upload_data.sh` script.
