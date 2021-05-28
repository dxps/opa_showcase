#!/bin/sh

curl -X PUT localhost:8181/v1/policies/products -d @products_policy.rego

