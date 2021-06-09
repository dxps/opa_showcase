#!/bin/sh

curl -X POST localhost:8181/v1/data/products/dashboard/policy1/subject_is_support \
     -d@query_authz___subject_is_support___input.json
