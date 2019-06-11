package gaia

import (
	"fmt"
	"time"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// PrivilegeEscalationPolicyIdentity represents the Identity of the object.
var PrivilegeEscalationPolicyIdentity = elemental.Identity{
	Name:     "privilegeescalationpolicy",
	Category: "privilegeescalationpolicies",
	Package:  "squall",
	Private:  false,
}

// PrivilegeEscalationPoliciesList represents a list of PrivilegeEscalationPolicies
type PrivilegeEscalationPoliciesList []*PrivilegeEscalationPolicy

// Identity returns the identity of the objects in the list.
func (o PrivilegeEscalationPoliciesList) Identity() elemental.Identity {

	return PrivilegeEscalationPolicyIdentity
}

// Copy returns a pointer to a copy the PrivilegeEscalationPoliciesList.
func (o PrivilegeEscalationPoliciesList) Copy() elemental.Identifiables {

	copy := append(PrivilegeEscalationPoliciesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the PrivilegeEscalationPoliciesList.
func (o PrivilegeEscalationPoliciesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(PrivilegeEscalationPoliciesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*PrivilegeEscalationPolicy))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o PrivilegeEscalationPoliciesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o PrivilegeEscalationPoliciesList) DefaultOrder() []string {

	return []string{
		"namespace",
		"name",
	}
}

// ToSparse returns the PrivilegeEscalationPoliciesList converted to SparsePrivilegeEscalationPoliciesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o PrivilegeEscalationPoliciesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparsePrivilegeEscalationPoliciesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparsePrivilegeEscalationPolicy)
	}

	return out
}

// Version returns the version of the content.
func (o PrivilegeEscalationPoliciesList) Version() int {

	return 1
}

// PrivilegeEscalationPolicy represents the model of a privilegeescalationpolicy
type PrivilegeEscalationPolicy struct {
	// ID is the identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// ActiveDuration defines for how long the policy will be active according to the
	// activeSchedule.
	ActiveDuration string `json:"activeDuration" msgpack:"activeDuration" bson:"activeduration" mapstructure:"activeDuration,omitempty"`

	// ActiveSchedule defines when the policy should be active using the cron notation.
	// The policy will be active for the given activeDuration.
	ActiveSchedule string `json:"activeSchedule" msgpack:"activeSchedule" bson:"activeschedule" mapstructure:"activeSchedule,omitempty"`

	// Annotation stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" msgpack:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// AssociatedTags are the list of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" msgpack:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey string `json:"-" msgpack:"-" bson:"createidempotencykey" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" msgpack:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Description is the description of the object.
	Description string `json:"description" msgpack:"description" bson:"description" mapstructure:"description,omitempty"`

	// Disabled defines if the propert is disabled.
	Disabled bool `json:"disabled" msgpack:"disabled" bson:"disabled" mapstructure:"disabled,omitempty"`

	// Fallback indicates that this is fallback policy. It will only be
	// applied if no other policies have been resolved. If the policy is also
	// propagated it will become a fallback for children namespaces.
	Fallback bool `json:"fallback" msgpack:"fallback" bson:"fallback" mapstructure:"fallback,omitempty"`

	// Metadata contains tags that can only be set during creation. They must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata []string `json:"metadata" msgpack:"metadata" bson:"metadata" mapstructure:"metadata,omitempty"`

	// Name is the name of the entity.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// NormalizedTags contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" msgpack:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Object contains the tag expression matching the enforcers the subject is allowed
	// to connect to.
	Object [][]string `json:"object" msgpack:"object" bson:"-" mapstructure:"object,omitempty"`

	// Propagate will propagate the policy to all of its children.
	Propagate bool `json:"propagate" msgpack:"propagate" bson:"propagate" mapstructure:"propagate,omitempty"`

	// Protected defines if the object is protected.
	Protected bool `json:"protected" msgpack:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// Subject contains the tag expression the tags need to match for the policy to
	// apply.
	Subject [][]string `json:"subject" msgpack:"subject" bson:"-" mapstructure:"subject,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey string `json:"-" msgpack:"-" bson:"updateidempotencykey" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" msgpack:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// geographical zone. This is used for sharding and
	// georedundancy.
	Zone int `json:"-" msgpack:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewPrivilegeEscalationPolicy returns a new *PrivilegeEscalationPolicy
func NewPrivilegeEscalationPolicy() *PrivilegeEscalationPolicy {

	return &PrivilegeEscalationPolicy{
		ModelVersion:   1,
		Annotations:    map[string][]string{},
		AssociatedTags: []string{},
		Metadata:       []string{},
		NormalizedTags: []string{},
		Object:         [][]string{},
		Subject:        [][]string{},
	}
}

// Identity returns the Identity of the object.
func (o *PrivilegeEscalationPolicy) Identity() elemental.Identity {

	return PrivilegeEscalationPolicyIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *PrivilegeEscalationPolicy) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *PrivilegeEscalationPolicy) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *PrivilegeEscalationPolicy) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *PrivilegeEscalationPolicy) DefaultOrder() []string {

	return []string{
		"namespace",
		"name",
	}
}

// Doc returns the documentation for the object
func (o *PrivilegeEscalationPolicy) Doc() string {

	return `The privilege escalation policy controls which PU/users can escalate privilege
on a linux system.`
}

func (o *PrivilegeEscalationPolicy) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetActiveDuration returns the ActiveDuration of the receiver.
func (o *PrivilegeEscalationPolicy) GetActiveDuration() string {

	return o.ActiveDuration
}

// SetActiveDuration sets the property ActiveDuration of the receiver using the given value.
func (o *PrivilegeEscalationPolicy) SetActiveDuration(activeDuration string) {

	o.ActiveDuration = activeDuration
}

// GetActiveSchedule returns the ActiveSchedule of the receiver.
func (o *PrivilegeEscalationPolicy) GetActiveSchedule() string {

	return o.ActiveSchedule
}

// SetActiveSchedule sets the property ActiveSchedule of the receiver using the given value.
func (o *PrivilegeEscalationPolicy) SetActiveSchedule(activeSchedule string) {

	o.ActiveSchedule = activeSchedule
}

// GetAnnotations returns the Annotations of the receiver.
func (o *PrivilegeEscalationPolicy) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *PrivilegeEscalationPolicy) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *PrivilegeEscalationPolicy) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *PrivilegeEscalationPolicy) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *PrivilegeEscalationPolicy) GetCreateIdempotencyKey() string {

	return o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the given value.
func (o *PrivilegeEscalationPolicy) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *PrivilegeEscalationPolicy) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *PrivilegeEscalationPolicy) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetDescription returns the Description of the receiver.
func (o *PrivilegeEscalationPolicy) GetDescription() string {

	return o.Description
}

// SetDescription sets the property Description of the receiver using the given value.
func (o *PrivilegeEscalationPolicy) SetDescription(description string) {

	o.Description = description
}

// GetDisabled returns the Disabled of the receiver.
func (o *PrivilegeEscalationPolicy) GetDisabled() bool {

	return o.Disabled
}

// SetDisabled sets the property Disabled of the receiver using the given value.
func (o *PrivilegeEscalationPolicy) SetDisabled(disabled bool) {

	o.Disabled = disabled
}

// GetFallback returns the Fallback of the receiver.
func (o *PrivilegeEscalationPolicy) GetFallback() bool {

	return o.Fallback
}

// SetFallback sets the property Fallback of the receiver using the given value.
func (o *PrivilegeEscalationPolicy) SetFallback(fallback bool) {

	o.Fallback = fallback
}

// GetMetadata returns the Metadata of the receiver.
func (o *PrivilegeEscalationPolicy) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the given value.
func (o *PrivilegeEscalationPolicy) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// GetName returns the Name of the receiver.
func (o *PrivilegeEscalationPolicy) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *PrivilegeEscalationPolicy) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *PrivilegeEscalationPolicy) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *PrivilegeEscalationPolicy) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *PrivilegeEscalationPolicy) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *PrivilegeEscalationPolicy) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetPropagate returns the Propagate of the receiver.
func (o *PrivilegeEscalationPolicy) GetPropagate() bool {

	return o.Propagate
}

// SetPropagate sets the property Propagate of the receiver using the given value.
func (o *PrivilegeEscalationPolicy) SetPropagate(propagate bool) {

	o.Propagate = propagate
}

// GetProtected returns the Protected of the receiver.
func (o *PrivilegeEscalationPolicy) GetProtected() bool {

	return o.Protected
}

// SetProtected sets the property Protected of the receiver using the given value.
func (o *PrivilegeEscalationPolicy) SetProtected(protected bool) {

	o.Protected = protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *PrivilegeEscalationPolicy) GetUpdateIdempotencyKey() string {

	return o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the given value.
func (o *PrivilegeEscalationPolicy) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *PrivilegeEscalationPolicy) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *PrivilegeEscalationPolicy) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *PrivilegeEscalationPolicy) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *PrivilegeEscalationPolicy) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *PrivilegeEscalationPolicy) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *PrivilegeEscalationPolicy) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *PrivilegeEscalationPolicy) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparsePrivilegeEscalationPolicy{
			ID:                   &o.ID,
			ActiveDuration:       &o.ActiveDuration,
			ActiveSchedule:       &o.ActiveSchedule,
			Annotations:          &o.Annotations,
			AssociatedTags:       &o.AssociatedTags,
			CreateIdempotencyKey: &o.CreateIdempotencyKey,
			CreateTime:           &o.CreateTime,
			Description:          &o.Description,
			Disabled:             &o.Disabled,
			Fallback:             &o.Fallback,
			Metadata:             &o.Metadata,
			Name:                 &o.Name,
			Namespace:            &o.Namespace,
			NormalizedTags:       &o.NormalizedTags,
			Object:               &o.Object,
			Propagate:            &o.Propagate,
			Protected:            &o.Protected,
			Subject:              &o.Subject,
			UpdateIdempotencyKey: &o.UpdateIdempotencyKey,
			UpdateTime:           &o.UpdateTime,
			ZHash:                &o.ZHash,
			Zone:                 &o.Zone,
		}
	}

	sp := &SparsePrivilegeEscalationPolicy{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "activeDuration":
			sp.ActiveDuration = &(o.ActiveDuration)
		case "activeSchedule":
			sp.ActiveSchedule = &(o.ActiveSchedule)
		case "annotations":
			sp.Annotations = &(o.Annotations)
		case "associatedTags":
			sp.AssociatedTags = &(o.AssociatedTags)
		case "createIdempotencyKey":
			sp.CreateIdempotencyKey = &(o.CreateIdempotencyKey)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "description":
			sp.Description = &(o.Description)
		case "disabled":
			sp.Disabled = &(o.Disabled)
		case "fallback":
			sp.Fallback = &(o.Fallback)
		case "metadata":
			sp.Metadata = &(o.Metadata)
		case "name":
			sp.Name = &(o.Name)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "normalizedTags":
			sp.NormalizedTags = &(o.NormalizedTags)
		case "object":
			sp.Object = &(o.Object)
		case "propagate":
			sp.Propagate = &(o.Propagate)
		case "protected":
			sp.Protected = &(o.Protected)
		case "subject":
			sp.Subject = &(o.Subject)
		case "updateIdempotencyKey":
			sp.UpdateIdempotencyKey = &(o.UpdateIdempotencyKey)
		case "updateTime":
			sp.UpdateTime = &(o.UpdateTime)
		case "zHash":
			sp.ZHash = &(o.ZHash)
		case "zone":
			sp.Zone = &(o.Zone)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparsePrivilegeEscalationPolicy to the object.
func (o *PrivilegeEscalationPolicy) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparsePrivilegeEscalationPolicy)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.ActiveDuration != nil {
		o.ActiveDuration = *so.ActiveDuration
	}
	if so.ActiveSchedule != nil {
		o.ActiveSchedule = *so.ActiveSchedule
	}
	if so.Annotations != nil {
		o.Annotations = *so.Annotations
	}
	if so.AssociatedTags != nil {
		o.AssociatedTags = *so.AssociatedTags
	}
	if so.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = *so.CreateIdempotencyKey
	}
	if so.CreateTime != nil {
		o.CreateTime = *so.CreateTime
	}
	if so.Description != nil {
		o.Description = *so.Description
	}
	if so.Disabled != nil {
		o.Disabled = *so.Disabled
	}
	if so.Fallback != nil {
		o.Fallback = *so.Fallback
	}
	if so.Metadata != nil {
		o.Metadata = *so.Metadata
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.NormalizedTags != nil {
		o.NormalizedTags = *so.NormalizedTags
	}
	if so.Object != nil {
		o.Object = *so.Object
	}
	if so.Propagate != nil {
		o.Propagate = *so.Propagate
	}
	if so.Protected != nil {
		o.Protected = *so.Protected
	}
	if so.Subject != nil {
		o.Subject = *so.Subject
	}
	if so.UpdateIdempotencyKey != nil {
		o.UpdateIdempotencyKey = *so.UpdateIdempotencyKey
	}
	if so.UpdateTime != nil {
		o.UpdateTime = *so.UpdateTime
	}
	if so.ZHash != nil {
		o.ZHash = *so.ZHash
	}
	if so.Zone != nil {
		o.Zone = *so.Zone
	}
}

// DeepCopy returns a deep copy if the PrivilegeEscalationPolicy.
func (o *PrivilegeEscalationPolicy) DeepCopy() *PrivilegeEscalationPolicy {

	if o == nil {
		return nil
	}

	out := &PrivilegeEscalationPolicy{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *PrivilegeEscalationPolicy.
func (o *PrivilegeEscalationPolicy) DeepCopyInto(out *PrivilegeEscalationPolicy) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy PrivilegeEscalationPolicy: %s", err))
	}

	*out = *target.(*PrivilegeEscalationPolicy)
}

// Validate valides the current information stored into the structure.
func (o *PrivilegeEscalationPolicy) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidatePattern("activeDuration", o.ActiveDuration, `^[0-9]+[smh]$`, `must be a valid duration like <n>s or <n>s or <n>h`, false); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateTagsWithoutReservedPrefixes("associatedTags", o.AssociatedTags); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("description", o.Description, 1024, false); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateMetadata("metadata", o.Metadata); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("name", o.Name, 256, false); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateTagsExpression("object", o.Object); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateTagsExpression("subject", o.Subject); err != nil {
		errors = errors.Append(err)
	}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*PrivilegeEscalationPolicy) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := PrivilegeEscalationPolicyAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return PrivilegeEscalationPolicyLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*PrivilegeEscalationPolicy) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return PrivilegeEscalationPolicyAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *PrivilegeEscalationPolicy) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "activeDuration":
		return o.ActiveDuration
	case "activeSchedule":
		return o.ActiveSchedule
	case "annotations":
		return o.Annotations
	case "associatedTags":
		return o.AssociatedTags
	case "createIdempotencyKey":
		return o.CreateIdempotencyKey
	case "createTime":
		return o.CreateTime
	case "description":
		return o.Description
	case "disabled":
		return o.Disabled
	case "fallback":
		return o.Fallback
	case "metadata":
		return o.Metadata
	case "name":
		return o.Name
	case "namespace":
		return o.Namespace
	case "normalizedTags":
		return o.NormalizedTags
	case "object":
		return o.Object
	case "propagate":
		return o.Propagate
	case "protected":
		return o.Protected
	case "subject":
		return o.Subject
	case "updateIdempotencyKey":
		return o.UpdateIdempotencyKey
	case "updateTime":
		return o.UpdateTime
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// PrivilegeEscalationPolicyAttributesMap represents the map of attribute for PrivilegeEscalationPolicy.
var PrivilegeEscalationPolicyAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `ID is the identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "string",
	},
	"ActiveDuration": elemental.AttributeSpecification{
		AllowedChars:   `^[0-9]+[smh]$`,
		AllowedChoices: []string{},
		ConvertedName:  "ActiveDuration",
		Description: `ActiveDuration defines for how long the policy will be active according to the
activeSchedule.`,
		Exposed: true,
		Getter:  true,
		Name:    "activeDuration",
		Setter:  true,
		Stored:  true,
		Type:    "string",
	},
	"ActiveSchedule": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ActiveSchedule",
		Description: `ActiveSchedule defines when the policy should be active using the cron notation.
The policy will be active for the given activeDuration.`,
		Exposed: true,
		Getter:  true,
		Name:    "activeSchedule",
		Setter:  true,
		Stored:  true,
		Type:    "string",
	},
	"Annotations": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Annotations",
		Description:    `Annotation stores additional information about an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "annotations",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string][]string",
		Type:           "external",
	},
	"AssociatedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedTags",
		Description:    `AssociatedTags are the list of tags attached to an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"CreateIdempotencyKey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateIdempotencyKey",
		Description:    `internal idempotency key for a create operation.`,
		Getter:         true,
		Name:           "createIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"CreateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `Creation date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"Description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `Description is the description of the object.`,
		Exposed:        true,
		Getter:         true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"Disabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Disabled",
		Description:    `Disabled defines if the propert is disabled.`,
		Exposed:        true,
		Getter:         true,
		Name:           "disabled",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"Fallback": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Fallback",
		Description: `Fallback indicates that this is fallback policy. It will only be
applied if no other policies have been resolved. If the policy is also
propagated it will become a fallback for children namespaces.`,
		Exposed:   true,
		Getter:    true,
		Name:      "fallback",
		Orderable: true,
		Setter:    true,
		Stored:    true,
		Type:      "boolean",
	},
	"Metadata": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Metadata",
		CreationOnly:   true,
		Description: `Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.`,
		Exposed:    true,
		Filterable: true,
		Getter:     true,
		Name:       "metadata",
		Setter:     true,
		Stored:     true,
		SubType:    "string",
		Type:       "list",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		DefaultOrder:   true,
		Description:    `Name is the name of the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		MaxLength:      256,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"Namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Namespace",
		DefaultOrder:   true,
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"NormalizedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "NormalizedTags",
		Description:    `NormalizedTags contains the list of normalized tags of the entities.`,
		Exposed:        true,
		Getter:         true,
		Name:           "normalizedTags",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Transient:      true,
		Type:           "list",
	},
	"Object": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Object",
		Description: `Object contains the tag expression matching the enforcers the subject is allowed
to connect to.`,
		Exposed:   true,
		Name:      "object",
		Orderable: true,
		SubType:   "[][]string",
		Type:      "external",
	},
	"Propagate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Propagate",
		Description:    `Propagate will propagate the policy to all of its children.`,
		Exposed:        true,
		Getter:         true,
		Name:           "propagate",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"Protected": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Protected",
		Description:    `Protected defines if the object is protected.`,
		Exposed:        true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"Subject": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Subject",
		Description: `Subject contains the tag expression the tags need to match for the policy to
apply.`,
		Exposed:   true,
		Name:      "subject",
		Orderable: true,
		SubType:   "[][]string",
		Type:      "external",
	},
	"UpdateIdempotencyKey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateIdempotencyKey",
		Description:    `internal idempotency key for a update operation.`,
		Getter:         true,
		Name:           "updateIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"UpdateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"ZHash": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ZHash",
		Description: `geographical hash of the data. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zHash",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
	"Zone": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Zone",
		Description: `geographical zone. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zone",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
}

// PrivilegeEscalationPolicyLowerCaseAttributesMap represents the map of attribute for PrivilegeEscalationPolicy.
var PrivilegeEscalationPolicyLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `ID is the identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "string",
	},
	"activeduration": elemental.AttributeSpecification{
		AllowedChars:   `^[0-9]+[smh]$`,
		AllowedChoices: []string{},
		ConvertedName:  "ActiveDuration",
		Description: `ActiveDuration defines for how long the policy will be active according to the
activeSchedule.`,
		Exposed: true,
		Getter:  true,
		Name:    "activeDuration",
		Setter:  true,
		Stored:  true,
		Type:    "string",
	},
	"activeschedule": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ActiveSchedule",
		Description: `ActiveSchedule defines when the policy should be active using the cron notation.
The policy will be active for the given activeDuration.`,
		Exposed: true,
		Getter:  true,
		Name:    "activeSchedule",
		Setter:  true,
		Stored:  true,
		Type:    "string",
	},
	"annotations": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Annotations",
		Description:    `Annotation stores additional information about an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "annotations",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string][]string",
		Type:           "external",
	},
	"associatedtags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedTags",
		Description:    `AssociatedTags are the list of tags attached to an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"createidempotencykey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateIdempotencyKey",
		Description:    `internal idempotency key for a create operation.`,
		Getter:         true,
		Name:           "createIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"createtime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `Creation date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `Description is the description of the object.`,
		Exposed:        true,
		Getter:         true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"disabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Disabled",
		Description:    `Disabled defines if the propert is disabled.`,
		Exposed:        true,
		Getter:         true,
		Name:           "disabled",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"fallback": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Fallback",
		Description: `Fallback indicates that this is fallback policy. It will only be
applied if no other policies have been resolved. If the policy is also
propagated it will become a fallback for children namespaces.`,
		Exposed:   true,
		Getter:    true,
		Name:      "fallback",
		Orderable: true,
		Setter:    true,
		Stored:    true,
		Type:      "boolean",
	},
	"metadata": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Metadata",
		CreationOnly:   true,
		Description: `Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.`,
		Exposed:    true,
		Filterable: true,
		Getter:     true,
		Name:       "metadata",
		Setter:     true,
		Stored:     true,
		SubType:    "string",
		Type:       "list",
	},
	"name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		DefaultOrder:   true,
		Description:    `Name is the name of the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		MaxLength:      256,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Namespace",
		DefaultOrder:   true,
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"normalizedtags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "NormalizedTags",
		Description:    `NormalizedTags contains the list of normalized tags of the entities.`,
		Exposed:        true,
		Getter:         true,
		Name:           "normalizedTags",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Transient:      true,
		Type:           "list",
	},
	"object": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Object",
		Description: `Object contains the tag expression matching the enforcers the subject is allowed
to connect to.`,
		Exposed:   true,
		Name:      "object",
		Orderable: true,
		SubType:   "[][]string",
		Type:      "external",
	},
	"propagate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Propagate",
		Description:    `Propagate will propagate the policy to all of its children.`,
		Exposed:        true,
		Getter:         true,
		Name:           "propagate",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"protected": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Protected",
		Description:    `Protected defines if the object is protected.`,
		Exposed:        true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"subject": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Subject",
		Description: `Subject contains the tag expression the tags need to match for the policy to
apply.`,
		Exposed:   true,
		Name:      "subject",
		Orderable: true,
		SubType:   "[][]string",
		Type:      "external",
	},
	"updateidempotencykey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateIdempotencyKey",
		Description:    `internal idempotency key for a update operation.`,
		Getter:         true,
		Name:           "updateIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"updatetime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"zhash": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ZHash",
		Description: `geographical hash of the data. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zHash",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
	"zone": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Zone",
		Description: `geographical zone. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zone",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
}

// SparsePrivilegeEscalationPoliciesList represents a list of SparsePrivilegeEscalationPolicies
type SparsePrivilegeEscalationPoliciesList []*SparsePrivilegeEscalationPolicy

// Identity returns the identity of the objects in the list.
func (o SparsePrivilegeEscalationPoliciesList) Identity() elemental.Identity {

	return PrivilegeEscalationPolicyIdentity
}

// Copy returns a pointer to a copy the SparsePrivilegeEscalationPoliciesList.
func (o SparsePrivilegeEscalationPoliciesList) Copy() elemental.Identifiables {

	copy := append(SparsePrivilegeEscalationPoliciesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparsePrivilegeEscalationPoliciesList.
func (o SparsePrivilegeEscalationPoliciesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparsePrivilegeEscalationPoliciesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparsePrivilegeEscalationPolicy))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparsePrivilegeEscalationPoliciesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparsePrivilegeEscalationPoliciesList) DefaultOrder() []string {

	return []string{
		"namespace",
		"name",
	}
}

// ToPlain returns the SparsePrivilegeEscalationPoliciesList converted to PrivilegeEscalationPoliciesList.
func (o SparsePrivilegeEscalationPoliciesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparsePrivilegeEscalationPoliciesList) Version() int {

	return 1
}

// SparsePrivilegeEscalationPolicy represents the sparse version of a privilegeescalationpolicy.
type SparsePrivilegeEscalationPolicy struct {
	// ID is the identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// ActiveDuration defines for how long the policy will be active according to the
	// activeSchedule.
	ActiveDuration *string `json:"activeDuration,omitempty" msgpack:"activeDuration,omitempty" bson:"activeduration,omitempty" mapstructure:"activeDuration,omitempty"`

	// ActiveSchedule defines when the policy should be active using the cron notation.
	// The policy will be active for the given activeDuration.
	ActiveSchedule *string `json:"activeSchedule,omitempty" msgpack:"activeSchedule,omitempty" bson:"activeschedule,omitempty" mapstructure:"activeSchedule,omitempty"`

	// Annotation stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" msgpack:"annotations,omitempty" bson:"annotations,omitempty" mapstructure:"annotations,omitempty"`

	// AssociatedTags are the list of tags attached to an entity.
	AssociatedTags *[]string `json:"associatedTags,omitempty" msgpack:"associatedTags,omitempty" bson:"associatedtags,omitempty" mapstructure:"associatedTags,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey *string `json:"-" msgpack:"-" bson:"createidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" msgpack:"createTime,omitempty" bson:"createtime,omitempty" mapstructure:"createTime,omitempty"`

	// Description is the description of the object.
	Description *string `json:"description,omitempty" msgpack:"description,omitempty" bson:"description,omitempty" mapstructure:"description,omitempty"`

	// Disabled defines if the propert is disabled.
	Disabled *bool `json:"disabled,omitempty" msgpack:"disabled,omitempty" bson:"disabled,omitempty" mapstructure:"disabled,omitempty"`

	// Fallback indicates that this is fallback policy. It will only be
	// applied if no other policies have been resolved. If the policy is also
	// propagated it will become a fallback for children namespaces.
	Fallback *bool `json:"fallback,omitempty" msgpack:"fallback,omitempty" bson:"fallback,omitempty" mapstructure:"fallback,omitempty"`

	// Metadata contains tags that can only be set during creation. They must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata *[]string `json:"metadata,omitempty" msgpack:"metadata,omitempty" bson:"metadata,omitempty" mapstructure:"metadata,omitempty"`

	// Name is the name of the entity.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// NormalizedTags contains the list of normalized tags of the entities.
	NormalizedTags *[]string `json:"normalizedTags,omitempty" msgpack:"normalizedTags,omitempty" bson:"normalizedtags,omitempty" mapstructure:"normalizedTags,omitempty"`

	// Object contains the tag expression matching the enforcers the subject is allowed
	// to connect to.
	Object *[][]string `json:"object,omitempty" msgpack:"object,omitempty" bson:"-" mapstructure:"object,omitempty"`

	// Propagate will propagate the policy to all of its children.
	Propagate *bool `json:"propagate,omitempty" msgpack:"propagate,omitempty" bson:"propagate,omitempty" mapstructure:"propagate,omitempty"`

	// Protected defines if the object is protected.
	Protected *bool `json:"protected,omitempty" msgpack:"protected,omitempty" bson:"protected,omitempty" mapstructure:"protected,omitempty"`

	// Subject contains the tag expression the tags need to match for the policy to
	// apply.
	Subject *[][]string `json:"subject,omitempty" msgpack:"subject,omitempty" bson:"-" mapstructure:"subject,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey *string `json:"-" msgpack:"-" bson:"updateidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime *time.Time `json:"updateTime,omitempty" msgpack:"updateTime,omitempty" bson:"updatetime,omitempty" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// geographical zone. This is used for sharding and
	// georedundancy.
	Zone *int `json:"-" msgpack:"-" bson:"zone,omitempty" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparsePrivilegeEscalationPolicy returns a new  SparsePrivilegeEscalationPolicy.
func NewSparsePrivilegeEscalationPolicy() *SparsePrivilegeEscalationPolicy {
	return &SparsePrivilegeEscalationPolicy{}
}

// Identity returns the Identity of the sparse object.
func (o *SparsePrivilegeEscalationPolicy) Identity() elemental.Identity {

	return PrivilegeEscalationPolicyIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparsePrivilegeEscalationPolicy) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparsePrivilegeEscalationPolicy) SetIdentifier(id string) {

	o.ID = &id
}

// Version returns the hardcoded version of the model.
func (o *SparsePrivilegeEscalationPolicy) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparsePrivilegeEscalationPolicy) ToPlain() elemental.PlainIdentifiable {

	out := NewPrivilegeEscalationPolicy()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.ActiveDuration != nil {
		out.ActiveDuration = *o.ActiveDuration
	}
	if o.ActiveSchedule != nil {
		out.ActiveSchedule = *o.ActiveSchedule
	}
	if o.Annotations != nil {
		out.Annotations = *o.Annotations
	}
	if o.AssociatedTags != nil {
		out.AssociatedTags = *o.AssociatedTags
	}
	if o.CreateIdempotencyKey != nil {
		out.CreateIdempotencyKey = *o.CreateIdempotencyKey
	}
	if o.CreateTime != nil {
		out.CreateTime = *o.CreateTime
	}
	if o.Description != nil {
		out.Description = *o.Description
	}
	if o.Disabled != nil {
		out.Disabled = *o.Disabled
	}
	if o.Fallback != nil {
		out.Fallback = *o.Fallback
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.NormalizedTags != nil {
		out.NormalizedTags = *o.NormalizedTags
	}
	if o.Object != nil {
		out.Object = *o.Object
	}
	if o.Propagate != nil {
		out.Propagate = *o.Propagate
	}
	if o.Protected != nil {
		out.Protected = *o.Protected
	}
	if o.Subject != nil {
		out.Subject = *o.Subject
	}
	if o.UpdateIdempotencyKey != nil {
		out.UpdateIdempotencyKey = *o.UpdateIdempotencyKey
	}
	if o.UpdateTime != nil {
		out.UpdateTime = *o.UpdateTime
	}
	if o.ZHash != nil {
		out.ZHash = *o.ZHash
	}
	if o.Zone != nil {
		out.Zone = *o.Zone
	}

	return out
}

// GetActiveDuration returns the ActiveDuration of the receiver.
func (o *SparsePrivilegeEscalationPolicy) GetActiveDuration() string {

	return *o.ActiveDuration
}

// SetActiveDuration sets the property ActiveDuration of the receiver using the address of the given value.
func (o *SparsePrivilegeEscalationPolicy) SetActiveDuration(activeDuration string) {

	o.ActiveDuration = &activeDuration
}

// GetActiveSchedule returns the ActiveSchedule of the receiver.
func (o *SparsePrivilegeEscalationPolicy) GetActiveSchedule() string {

	return *o.ActiveSchedule
}

// SetActiveSchedule sets the property ActiveSchedule of the receiver using the address of the given value.
func (o *SparsePrivilegeEscalationPolicy) SetActiveSchedule(activeSchedule string) {

	o.ActiveSchedule = &activeSchedule
}

// GetAnnotations returns the Annotations of the receiver.
func (o *SparsePrivilegeEscalationPolicy) GetAnnotations() map[string][]string {

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparsePrivilegeEscalationPolicy) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparsePrivilegeEscalationPolicy) GetAssociatedTags() []string {

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparsePrivilegeEscalationPolicy) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *SparsePrivilegeEscalationPolicy) GetCreateIdempotencyKey() string {

	return *o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the address of the given value.
func (o *SparsePrivilegeEscalationPolicy) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = &createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparsePrivilegeEscalationPolicy) GetCreateTime() time.Time {

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparsePrivilegeEscalationPolicy) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetDescription returns the Description of the receiver.
func (o *SparsePrivilegeEscalationPolicy) GetDescription() string {

	return *o.Description
}

// SetDescription sets the property Description of the receiver using the address of the given value.
func (o *SparsePrivilegeEscalationPolicy) SetDescription(description string) {

	o.Description = &description
}

// GetDisabled returns the Disabled of the receiver.
func (o *SparsePrivilegeEscalationPolicy) GetDisabled() bool {

	return *o.Disabled
}

// SetDisabled sets the property Disabled of the receiver using the address of the given value.
func (o *SparsePrivilegeEscalationPolicy) SetDisabled(disabled bool) {

	o.Disabled = &disabled
}

// GetFallback returns the Fallback of the receiver.
func (o *SparsePrivilegeEscalationPolicy) GetFallback() bool {

	return *o.Fallback
}

// SetFallback sets the property Fallback of the receiver using the address of the given value.
func (o *SparsePrivilegeEscalationPolicy) SetFallback(fallback bool) {

	o.Fallback = &fallback
}

// GetMetadata returns the Metadata of the receiver.
func (o *SparsePrivilegeEscalationPolicy) GetMetadata() []string {

	return *o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the address of the given value.
func (o *SparsePrivilegeEscalationPolicy) SetMetadata(metadata []string) {

	o.Metadata = &metadata
}

// GetName returns the Name of the receiver.
func (o *SparsePrivilegeEscalationPolicy) GetName() string {

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparsePrivilegeEscalationPolicy) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparsePrivilegeEscalationPolicy) GetNamespace() string {

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparsePrivilegeEscalationPolicy) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparsePrivilegeEscalationPolicy) GetNormalizedTags() []string {

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparsePrivilegeEscalationPolicy) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetPropagate returns the Propagate of the receiver.
func (o *SparsePrivilegeEscalationPolicy) GetPropagate() bool {

	return *o.Propagate
}

// SetPropagate sets the property Propagate of the receiver using the address of the given value.
func (o *SparsePrivilegeEscalationPolicy) SetPropagate(propagate bool) {

	o.Propagate = &propagate
}

// GetProtected returns the Protected of the receiver.
func (o *SparsePrivilegeEscalationPolicy) GetProtected() bool {

	return *o.Protected
}

// SetProtected sets the property Protected of the receiver using the address of the given value.
func (o *SparsePrivilegeEscalationPolicy) SetProtected(protected bool) {

	o.Protected = &protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *SparsePrivilegeEscalationPolicy) GetUpdateIdempotencyKey() string {

	return *o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the address of the given value.
func (o *SparsePrivilegeEscalationPolicy) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = &updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparsePrivilegeEscalationPolicy) GetUpdateTime() time.Time {

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparsePrivilegeEscalationPolicy) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *SparsePrivilegeEscalationPolicy) GetZHash() int {

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparsePrivilegeEscalationPolicy) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparsePrivilegeEscalationPolicy) GetZone() int {

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparsePrivilegeEscalationPolicy) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparsePrivilegeEscalationPolicy.
func (o *SparsePrivilegeEscalationPolicy) DeepCopy() *SparsePrivilegeEscalationPolicy {

	if o == nil {
		return nil
	}

	out := &SparsePrivilegeEscalationPolicy{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparsePrivilegeEscalationPolicy.
func (o *SparsePrivilegeEscalationPolicy) DeepCopyInto(out *SparsePrivilegeEscalationPolicy) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparsePrivilegeEscalationPolicy: %s", err))
	}

	*out = *target.(*SparsePrivilegeEscalationPolicy)
}