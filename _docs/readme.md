## Notes

1. [Download OPA](https://www.openpolicyagent.org/docs/latest/#1-download-opa) and run it as server using `./opa run -s`. By default, it listens for HTTP connections on port `8181`.
1. [REST API](https://www.openpolicyagent.org/docs/latest/rest-api/) is used for managing the policies.
1. For data fetching purposes (and other needs/use cases), [extending OPA](https://www.openpolicyagent.org/docs/latest/extensions/) is the path. This implies:
   - Use it in embedded mode: have a Golang based service that embeds OPA logic.
   - Extend its capabilities through custom built-in functions and plugins.
