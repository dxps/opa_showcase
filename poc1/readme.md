## PoC 1: OPA Sidecar in a realistic setup

Here is the setup this PoC aims to showcase.

```
                            .-------------.              .--------------.
        .------------------>|     IAM     |              |  POLICY MGMT |
        | 1. Authenticate   '-------------'              '--------------'
        |                                                        |
        |                                                        |
 .------------.                                                  |
 |   CLIENT   |                                                  | 0.
 '------------'                                                  | Push
        |                                                        | Policies
        |                                                        |
        | 2. Get Tasks     .-------------.                       |
        '----------------->|  TASKS API  |                       |
                           '-------------'                       |
                                  |                              |
                                  | 3. Get AuthZ Decision        |
                                  |                              |
                                  |        .-------.             |
                                  '------->|  OPA  |<------------'
                                           '-------'
```
