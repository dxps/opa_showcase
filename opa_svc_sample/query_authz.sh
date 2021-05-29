#!/bin/sh

curl -X POST localhost:8181/v1/data/products/policy/user_has_product \
     -d@query_input.json
