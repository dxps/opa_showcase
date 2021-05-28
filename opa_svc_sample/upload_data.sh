#!/bin/sh

curl -X PUT localhost:8181/v1/data/products_acl -d @products_acl.json
