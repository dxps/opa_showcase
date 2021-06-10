#!/bin/sh

curl -X POST localhost:8181/v1/data/rbac/subject_has_support_role \
     -d@002__query__subject_has_support_role__input.json
