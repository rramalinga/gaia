# Model
model:
  rest_name: httpresourcespec
  resource_name: httpresourcespecs
  entity_name: HTTPResourceSpec
  package: squall
  description: |-
    HTTPResourceSpec descibes an HTTP resource exposed by a service. These APIs
    can be associated with one or more services.
  indexes:
  - - :shard
    - zone
    - zhash
  - - namespace
  - - namespace
    - name
  - - namespace
    - archived
  - - namespace
    - normalizedtags
  get:
    description: Retrieves the object with the given ID.
    global_parameters:
    - $archivable
  update:
    description: Updates the object with the given ID.
  delete:
    description: Deletes the object with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@archivable'
  - '@base'
  - '@described'
  - '@identifiable-pk-stored'
  - '@named'
  - '@propagated'
  - '@metadatable'
  - '@zonable'

# Attributes
attributes:
  v1:
  - name: endpoints
    description: EndPoints is a list of API endpoints that are exposed for the service.
    type: refList
    exposed: true
    subtype: endpoint
    stored: true
    extensions:
      refMode: pointer