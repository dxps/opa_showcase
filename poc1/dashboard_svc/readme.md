## Dashboard Service

TBD

### Design

- Each app/service has an id, and an associated set of policies.
- The entry points where authorization needs to happen, also called *Policy Enforcement Point*s, are using the `AuthzFacade` component to ask for the authorization decision.
