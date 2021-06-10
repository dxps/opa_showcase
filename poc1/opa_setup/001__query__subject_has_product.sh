#!/bin/sh

curl -X POST localhost:8181/v1/data/products_enablement/subject_has_product \
     -d@001__query__subject_has_product__input.json
