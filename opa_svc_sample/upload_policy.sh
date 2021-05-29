#!/bin/sh

curl -X PUT localhost:8181/v1/policies/products \
     --data-binary @products_policy.rego

