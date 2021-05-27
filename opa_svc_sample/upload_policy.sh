#!/bin/sh

curl -X PUT localhost:8181/v1/policies/products -d @products__policy.rego

