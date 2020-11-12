# Model
model:
  rest_name: eventlog
  resource_name: eventlogs
  entity_name: EventLog
  package: leon
  group: core/monitoring
  description: Allows you to report various events on any object.
  extends:
  - '@identifiable-stored'
  - '@zoned'
  - '@migratable'

# Indexes
indexes:
- - namespace
  - timestamp

# Attributes
attributes:
  v1:
  - name: category
    description: Category of the event log.
    type: string
    exposed: true
    stored: true
    required: true
    creation_only: true
    example_value: enforcerd:policy
    omit_empty: true
    extensions:
      bson_name: a

  - name: content
    description: Content of the event log.
    type: string
    exposed: true
    stored: true
    required: true
    creation_only: true
    example_value: Unable to activate docker container xyz because abc.
    omit_empty: true
    extensions:
      bson_name: b

  - name: date
    description: Creation date of the event log.
    type: time
    exposed: true
    creation_only: true
    autogenerated: true
    deprecated: true
    orderable: true
    transient: true
    omit_empty: true
    extensions:
      bson_name: c

  - name: level
    description: Sets the log level.
    type: enum
    exposed: true
    stored: true
    creation_only: true
    allowed_choices:
    - Debug
    - Info
    - Warning
    - Error
    - Critical
    default_value: Info
    omit_empty: true
    extensions:
      bson_name: d

  - name: namespace
    description: Namespace tag attached to the event log.
    type: string
    exposed: true
    stored: true
    read_only: true
    creation_only: true
    autogenerated: true
    getter: true
    setter: true
    orderable: true
    omit_empty: true
    extensions:
      bson_name: e

  - name: opaque
    description: |-
      Opaque data that can be attached to the event log, for further machine
      processing.
    type: string
    exposed: true
    stored: true
    creation_only: true
    orderable: true
    omit_empty: true
    extensions:
      bson_name: f

  - name: targetID
    description: |-
      ID of the object this event log is attached to. The object must be in the same
      namespace than the event log.
    type: string
    exposed: true
    stored: true
    required: true
    creation_only: true
    example_value: xxx-xxx-xxx-xxx
    omit_empty: true
    extensions:
      bson_name: g

  - name: targetIdentity
    description: Identity of the object this event log is attached to.
    type: string
    exposed: true
    stored: true
    required: true
    creation_only: true
    example_value: processingunit
    omit_empty: true
    extensions:
      bson_name: h

  - name: timestamp
    description: Creation date of the event log.
    type: time
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: i

  - name: title
    description: Title of the event log.
    type: string
    exposed: true
    stored: true
    required: true
    creation_only: true
    example_value: Error while activating processing unit.
    omit_empty: true
    extensions:
      bson_name: j
