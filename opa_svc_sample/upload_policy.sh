#!/bin/sh

curl -X PUT localhost:8181/v1/policies/products/acl \
     --data-binary @products_acl_policy.rego

