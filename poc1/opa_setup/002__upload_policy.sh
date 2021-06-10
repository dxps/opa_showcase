#!/bin/sh

curl -X PUT localhost:8181/v1/policies/rbac \
     --data-binary @002__rbac.rego
