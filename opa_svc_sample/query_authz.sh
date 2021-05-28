#!/bin/sh

curl -X POST -H "Content-Type: application/json" \
     localhost:8181/v1/data/products/policy/allow \
     -d@query_input.json


