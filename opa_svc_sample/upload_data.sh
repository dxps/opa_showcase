#!/bin/sh

curl -X PUT localhost:8181/v1/data/products/acl \
    --data-binary @products_acl.json
