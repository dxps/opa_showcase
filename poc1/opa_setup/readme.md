## OPA Setup

This directory contains:

- the policies used within this PoC
- scripts to upload and test them (outside of any app/service)

The steps are:

0. Start OPA.
1. Upload the policy.<br/>
   (ex: `./001__upload_policy.sh`)
2. Query for authorization.<br/>
   (ex: `./001__query__subject_has_product.sh`)
