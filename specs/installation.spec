# Model
model:
  rest_name: installation
  resource_name: installations
  entity_name: Installation
  package: highwind
  description: Installation represents an installation for a given account.
  get: true
  update: true
  delete: true

# Attributes
attributes:
  v1:
  - name: ID
    description: ID is the identifier of the object.
    type: string
    exposed: true
    stored: true
    filterable: true
    format: free
    identifier: true
    orderable: true
    primary_key: true

  - name: accountName
    description: AccountName that should be installed.
    type: string
    exposed: true
    stored: true
    filterable: true
    format: free
    orderable: true