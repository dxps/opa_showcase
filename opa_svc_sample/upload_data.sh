#!/bin/sh

curl -X PUT localhost:8181/v1/data/products/acl -d @products_acl__data.json

