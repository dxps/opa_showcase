## PoC 1: OPA Sidecar in a realistic setup

Here is the setup this PoC aims to showcase how authorization can be delegated to the OPA, deployed as a sidecar (near the service that uses it).

```
                            .-------------.              .--------------.
        .------------------>|   IAM_SVC   |              |  POLICY MGMT |
        | 1. Authenticate   '-------------'              '--------------'
        |                                                        |
        |                                                        |
 .------------.                                                  |
 |   CLIENT   |                                                  | 0.
 '------------'                                                  | Push
        |                                                        | Policies
        |                                                        |
        | 2. Get Tasks     .---------------.                     |
        '----------------->| DASHBOARD_SVC |                     |
                           '---------------'                     |
                                   |                             |
                                   |3. Get AuthZ Decision        |
                                   |                             |
                                   |       .-------.             |
                                   '------>|  OPA  |<------------'
                                           '-------'
```
