#!/bin/sh

curl -X PUT localhost:8181/v1/policies/products_enablement \
     --data-binary @001__products_enablement.rego
