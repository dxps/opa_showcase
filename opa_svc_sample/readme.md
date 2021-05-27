## Using OPA as a standalone/sidecar service

OPA can run in server mode and following the sidecar pattern it can be used for:

- uploading `policies`
- feeding in `data`
- querying for getting authorization decisions

```
        Upload                  Query for AuthZ
  ------------------           ------------------

      .--------.
      | policy |
      '--------'
           |
           |        .---------.
           |        |   OPA   |        .--------.
           .------->| Service |<-------| access |
           |        '---------'        '--------'
           |
           |
      .--------.
      |  data  |
      '--------'
```
