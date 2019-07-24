# Model
model:
  rest_name: sshauthority
  resource_name: sshauthorities
  entity_name: SSHAuthority
  package: barret
  group: internal/ssh
  description: Internal api to deliver SSH CA.
  private: true
  delete:
    description: Deletes the SSH CA with the given ID.
  extends:
  - '@zoned'
  - '@migratable'
  - '@named'
  - '@timeable'
  - '@identifiable-stored'

# Attributes
attributes:
  v1:
  - name: alg
    description: Algorithm to use for the CA.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - RSA
    - ECDSA
    default_value: ECDSA

  - name: privateKey
    description: Contains the private key of the CA.
    type: string
    stored: true
    read_only: true
    autogenerated: true

  - name: publicKey
    description: Contains the public key of the CA.
    type: string
    exposed: true
    stored: true
    read_only: true
    autogenerated: true
