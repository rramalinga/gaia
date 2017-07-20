package squallmodels

import "fmt"
import "github.com/aporeto-inc/elemental"

import "sync"

import "time"

// ActivityIdentity represents the Identity of the object
var ActivityIdentity = elemental.Identity{
	Name:     "activity",
	Category: "activities",
}

// ActivitiesList represents a list of Activities
type ActivitiesList []*Activity

// ContentIdentity returns the identity of the objects in the list.
func (o ActivitiesList) ContentIdentity() elemental.Identity {

	return ActivityIdentity
}

// List converts the object to an elemental.IdentifiablesList.
func (o ActivitiesList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o ActivitiesList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o ActivitiesList) Version() int {

	return 1.0
}

// Activity represents the model of a activity
type Activity struct {
	// ID is the identifier of the object.
	ID string `json:"ID" bson:"_id"`

	// Claims of the user who performed the operation.
	Claims interface{} `json:"claims" bson:"claims"`

	// Data of the notification.
	Data interface{} `json:"data" bson:"data"`

	// Date of the notification.
	Date time.Time `json:"date" bson:"date"`

	// Error contains the eventual error.
	Error interface{} `json:"error" bson:"error"`

	// Message of the notification.
	Message string `json:"message" bson:"message"`

	// Namespace of the notification.
	Namespace string `json:"namespace" bson:"namespace"`

	// Operation describe what kind of operation the notification represents.
	Operation string `json:"operation" bson:"operation"`

	// OriginalData contains the eventual original data of the object that has been modified.
	OriginalData interface{} `json:"originalData" bson:"originaldata"`

	// Source contains meta information abbout the source.
	Source string `json:"source" bson:"source"`

	// TargetIdentity is the Identity of the related object.
	TargetIdentity string `json:"targetIdentity" bson:"targetidentity"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewActivity returns a new *Activity
func NewActivity() *Activity {

	return &Activity{
		ModelVersion: 1.0,
	}
}

// Identity returns the Identity of the object.
func (o *Activity) Identity() elemental.Identity {

	return ActivityIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Activity) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Activity) SetIdentifier(ID string) {

	o.ID = ID
}

// Version returns the hardcoded version of the model
func (o *Activity) Version() int {

	return 1.0
}

// DefaultOrder returns the list of default ordering fields.
func (o *Activity) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Activity) Doc() string {
	return `Contains all the activity log that happened in a namespace. All sucessful or failed actions will be available, and eventual errors as well as the claims of the user who triggered the actiions. This log is capped and only keeps the last 50k entries by default. `
}

func (o *Activity) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *Activity) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*Activity) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := ActivityAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return ActivityLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Activity) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ActivityAttributesMap
}

// ActivityAttributesMap represents the map of attribute for Activity.
var ActivityAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ID is the identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	"Claims": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Claims of the user who performed the operation.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "claims",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		SubType:        "raw_data",
		Type:           "external",
	},
	"Data": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Data of the notification.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "data",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		SubType:        "raw_data",
		Type:           "external",
	},
	"Date": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Date of the notification.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "date",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
	},
	"Error": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Error contains the eventual error.`,
		Exposed:        true,
		Name:           "error",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		SubType:        "raw_data",
		Type:           "external",
	},
	"Message": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Message of the notification.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "message",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Namespace of the notification.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Operation": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Operation describe what kind of operation the notification represents.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "operation",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"OriginalData": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `OriginalData contains the eventual original data of the object that has been modified.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "originalData",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		SubType:        "raw_data",
		Type:           "external",
	},
	"Source": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Source contains meta information abbout the source.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "source",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"TargetIdentity": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `TargetIdentity is the Identity of the related object.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "targetIdentity",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
}

// ActivityLowerCaseAttributesMap represents the map of attribute for Activity.
var ActivityLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ID is the identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	"claims": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Claims of the user who performed the operation.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "claims",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		SubType:        "raw_data",
		Type:           "external",
	},
	"data": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Data of the notification.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "data",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		SubType:        "raw_data",
		Type:           "external",
	},
	"date": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Date of the notification.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "date",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
	},
	"error": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Error contains the eventual error.`,
		Exposed:        true,
		Name:           "error",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		SubType:        "raw_data",
		Type:           "external",
	},
	"message": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Message of the notification.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "message",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Namespace of the notification.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"operation": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Operation describe what kind of operation the notification represents.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "operation",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"originaldata": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `OriginalData contains the eventual original data of the object that has been modified.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "originalData",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		SubType:        "raw_data",
		Type:           "external",
	},
	"source": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Source contains meta information abbout the source.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "source",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"targetidentity": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `TargetIdentity is the Identity of the related object.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "targetIdentity",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
}
