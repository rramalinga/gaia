# Model
model:
  rest_name: oauthkey
  resource_name: oauthkeys
  entity_name: OAUTHKey
  package: cactuar
  group: internal/token
  description: |-
    OAUTHInfo provides the information for an OAUTH server to retrieve the secrets
    that can validate a JWT token issued by us.
  get:
    description: Retrieves the OAUTH info.

# Attributes
attributes:
  v1:
  - name: keyString
    description: |-
      KeyString is the JWKS key response for an OAUTH verifier. It provides the OAUTH
      compatible signing keys.
    type: string
    exposed: true
    read_only: true
    autogenerated: true