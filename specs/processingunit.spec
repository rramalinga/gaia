# Model
model:
  rest_name: processingunit
  resource_name: processingunits
  entity_name: ProcessingUnit
  package: squall
  description: |-
    A Processing Unit reprents anything that can compute. It can be a Docker
    container, or a simple Unix process. They are created, updated and deleted by
    the system as they come and go. You can only modify its tags.  Processing Units
    use Network Access Policies to define which other Processing Units or External
    Services they can communicate with and File Access Policies to define what File
    Paths they can use.
  aliases:
  - pu
  - pus
  indexes:
  - - :shard
    - zone
    - zHash
  - - namespace
  - - namespace
    - name
  - - namespace
    - archived
  - - namespace
    - operationalStatus
    - archived
  - - namespace
    - normalizedTags
    - archived
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
  - '@metadatable'
  - '@named'
  - '@zonable'

# Attributes
attributes:
  v1:
  - name: collectInfo
    description: |-
      CollectInfo indicates to the enforcer it needs to collect information for this
      PU.
    type: boolean
    exposed: true
    stored: true

  - name: collectedInfo
    description: CollectedInfo represents the latest info collected by the enforcer
      for this PU.
    type: external
    exposed: true
    subtype: map_of_string_of_strings
    stored: true

  - name: enforcementStatus
    description: EnforcementStatus communicates the state of the enforcer for that
      PU.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Active
    - Failed
    - Inactive
    default_value: Inactive
    filterable: true

  - name: enforcerID
    description: EnforcerID is the ID of the enforcer associated with the processing
      unit.
    type: string
    exposed: true
    stored: true
    filterable: true

  - name: enforcerNamespace
    description: |-
      enforcerNamespace is the namespace of the enforcer associated with the
      processing unit.
    type: string
    exposed: true
    stored: true
    filterable: true

  - name: image
    description: Docker image, or path to executable.
    type: string
    exposed: true
    stored: true
    filterable: true

  - name: lastPokeTime
    description: Last poke is the time when the pu got last poked.
    type: time
    stored: true

  - name: lastSyncTime
    description: LastSyncTime is the time when the policy was last resolved.
    type: time
    exposed: true
    stored: true
    autogenerated: true
    orderable: true

  - name: nativeContextID
    description: NativeContextID is the Docker UUID or service PID.
    type: string
    exposed: true
    stored: true

  - name: networkServices
    description: |-
      NetworkServices is the list of services that this processing unit has declared
      that it will be listening to. This can happen either with an activation command
      or by exposing the ports in a container manifest.
    type: refList
    exposed: true
    subtype: processingunitservice
    stored: true
    orderable: true
    validations:
    - $processingunitservices
    extensions:
      refMode: pointer

  - name: operationalStatus
    description: OperationalStatus of the processing unit.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Initialized
    - Paused
    - Running
    - Stopped
    - Terminated
    default_value: Initialized
    filterable: true

  - name: tracing
    description: Tracing indicates if this PU must be placed in tracing mode.
    type: ref
    exposed: true
    subtype: tracemode
    stored: true
    extensions:
      refMode: pointer

  - name: type
    description: Type of the container ecosystem.
    type: enum
    exposed: true
    stored: true
    required: true
    creation_only: true
    allowed_choices:
    - Docker
    - LinuxService
    - RKT
    - User
    - APIGateway
    example_value: Docker
    filterable: true

# Relations
relations:
- rest_name: service
  get:
    description: Retrieves the services used by a processing unit.

- rest_name: renderedpolicy
  get:
    description: Retrieves the policies for the processing unit.
    parameters:
      entries:
      - name: csr
        description: CSR to sign.
        type: string
        example_value: |-
          --- BEGIN CSR ---
          xxx-xxx-xxx-xxx
          --- END CSR ---

- rest_name: vulnerability
  get:
    description: Retrieves the vulnerabilities affecting the processing unit.

- rest_name: poke
  get:
    description: |-
      Sends a poke empty object. This will send a snaphot of the pu to time series
      database.
    parameters:
      entries:
      - name: enforcementStatus
        description: If set, changes the enforcement status of the processing unit
          alongside with the poke.
        type: enum
        allowed_choices:
        - Failed
        - Inactive
        - Active

      - name: status
        description: If set, changes the status of the processing unit alongside with
          the poke.
        type: enum
        allowed_choices:
        - Paused
        - Running
        - Stopped
        example_value: Running

      - name: ts
        description: time of report. If not set, local server time will be used.
        type: time
