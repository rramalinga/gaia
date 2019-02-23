# Model
model:
  rest_name: issue
  resource_name: issue
  entity_name: Issue
  package: midgard
  group: core/authentication
  description: This API issues a new token according to given data.

# Attributes
attributes:
  v1:
  - name: audience
    description: |-
      If given, the issued token will only be valid from that namespace declared in
      that value.
    type: string
    exposed: true
    validations:
    - $audience

  - name: data
    description: Data contains additional data. The value depends on the issuer type.
    type: string
    exposed: true
    deprecated: true
    orderable: true

  - name: metadata
    description: Metadata contains various additional information. Meaning depends
      on the realm.
    type: external
    exposed: true
    subtype: map[string]interface{}
    example_value:
      vinceAccount: acme
      vinceOTP: 665435
      vincePassword: s3cr3t
    orderable: true

  - name: opaque
    description: Opaque data that will be included in the issued token.
    type: external
    exposed: true
    subtype: map[string]string

  - name: quota
    description: Restricts the number of time the issued token should be used.
    type: integer
    exposed: true

  - name: realm
    description: Realm is the authentication realm.
    type: enum
    exposed: true
    required: true
    allowed_choices:
    - AWSIdentityDocument
    - AWSSecurityToken
    - Certificate
    - Google
    - LDAP
    - Vince
    - GCPIdentityToken
    - AzureIdentityToken
    - OIDC
    example_value: Vince

  - name: token
    description: Token is the token to use for the registration.
    type: string
    exposed: true
    read_only: true
    autogenerated: true

  - name: validity
    description: |-
      Validity configures the max validity time for a token. If it is bigger than the
      configured max validity, it will be capped.
    type: string
    exposed: true
    allowed_chars: ^([0-9]+h[0-9]+m[0-9]+s|[0-9]+m[0-9]+s|[0-9]+m[0-9]+s|[0-9]+h[0-9]+s|[0-9]+h[0-9]+m|[0-9]+s|[0-9]+h|[0-9]+m)$
    allowed_chars_message: must be a valid duration like <n>s or <n>s or <n>h
    default_value: 24h
    orderable: true
