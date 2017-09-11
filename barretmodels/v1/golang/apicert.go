package barretmodels

import "fmt"
import "github.com/aporeto-inc/elemental"

import "sync"

import "time"

// APICertIdentity represents the Identity of the object
var APICertIdentity = elemental.Identity{
	Name:     "apicert",
	Category: "apicerts",
}

// APICertsList represents a list of APICerts
type APICertsList []*APICert

// ContentIdentity returns the identity of the objects in the list.
func (o APICertsList) ContentIdentity() elemental.Identity {

	return APICertIdentity
}

// List converts the object to an elemental.IdentifiablesList.
func (o APICertsList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o APICertsList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o APICertsList) Version() int {

	return 1
}

// APICert represents the model of a apicert
type APICert struct {
	// CSR contains the Certificate Signing Request as a PEM encoded string.
	CSR string `json:"CSR" bson:"-"`

	// ID contains the identifier of the certificate.
	ID string `json:"ID" bson:"-"`

	// Certificate contains the certificate data in PEM format.
	Certificate string `json:"certificate" bson:"-"`

	// ExpirationDate contains the requested expiration date.
	ExpirationDate time.Time `json:"expirationDate" bson:"-"`

	// KeyUsage contains the generated key usage,
	KeyUsage string `json:"keyUsage" bson:"-"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewAPICert returns a new *APICert
func NewAPICert() *APICert {

	return &APICert{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *APICert) Identity() elemental.Identity {

	return APICertIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *APICert) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *APICert) SetIdentifier(ID string) {

	o.ID = ID
}

// Version returns the hardcoded version of the model
func (o *APICert) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *APICert) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *APICert) Doc() string {
	return `This API allows to retrieve an client certifcate for api authentication.`
}

func (o *APICert) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *APICert) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("CSR", o.CSR); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("CSR", o.CSR); err != nil {
		errors = append(errors, err)
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
func (*APICert) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := APICertAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return APICertLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*APICert) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return APICertAttributesMap
}

// APICertAttributesMap represents the map of attribute for APICert.
var APICertAttributesMap = map[string]elemental.AttributeSpecification{
	"CSR": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `CSR contains the Certificate Signing Request as a PEM encoded string.`,
		Exposed:        true,
		Format:         "free",
		Name:           "CSR",
		Required:       true,
		Type:           "string",
	},
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ID contains the identifier of the certificate.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		ReadOnly:       true,
		Type:           "string",
	},
	"Certificate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Certificate contains the certificate data in PEM format.`,
		Exposed:        true,
		Format:         "free",
		Name:           "certificate",
		ReadOnly:       true,
		Type:           "string",
	},
	"ExpirationDate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `ExpirationDate contains the requested expiration date.`,
		Exposed:        true,
		Name:           "expirationDate",
		SubType:        "string",
		Type:           "time",
	},
	"KeyUsage": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `KeyUsage contains the generated key usage,`,
		Exposed:        true,
		Format:         "free",
		Name:           "keyUsage",
		ReadOnly:       true,
		Type:           "string",
	},
}

// APICertLowerCaseAttributesMap represents the map of attribute for APICert.
var APICertLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"csr": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `CSR contains the Certificate Signing Request as a PEM encoded string.`,
		Exposed:        true,
		Format:         "free",
		Name:           "CSR",
		Required:       true,
		Type:           "string",
	},
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ID contains the identifier of the certificate.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		ReadOnly:       true,
		Type:           "string",
	},
	"certificate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Certificate contains the certificate data in PEM format.`,
		Exposed:        true,
		Format:         "free",
		Name:           "certificate",
		ReadOnly:       true,
		Type:           "string",
	},
	"expirationdate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `ExpirationDate contains the requested expiration date.`,
		Exposed:        true,
		Name:           "expirationDate",
		SubType:        "string",
		Type:           "time",
	},
	"keyusage": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `KeyUsage contains the generated key usage,`,
		Exposed:        true,
		Format:         "free",
		Name:           "keyUsage",
		ReadOnly:       true,
		Type:           "string",
	},
}
