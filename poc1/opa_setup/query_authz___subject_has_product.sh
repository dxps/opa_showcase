#!/bin/sh

curl -X POST localhost:8181/v1/data/products/dashboard/policy1/subject_has_product \
     -d@query_authz___subject_has_product___input.json
