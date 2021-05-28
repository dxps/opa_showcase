#!/bin/sh

curl -X PUT localhost:8181/v1/policies/products_acl -d @products_acl_policy.rego

