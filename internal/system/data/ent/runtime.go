// Code generated by ent, DO NOT EDIT.

package ent

import (
	"vine-template-rpc/internal/system/data/ent/permission"
	"vine-template-rpc/internal/system/data/ent/resource"
	"vine-template-rpc/internal/system/data/ent/role"
	"vine-template-rpc/internal/system/data/ent/schema"
	"vine-template-rpc/internal/system/data/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	permissionFields := schema.Permission{}.Fields()
	_ = permissionFields
	// permissionDescName is the schema descriptor for name field.
	permissionDescName := permissionFields[0].Descriptor()
	// permission.NameValidator is a validator for the "name" field. It is called by the builders before save.
	permission.NameValidator = permissionDescName.Validators[0].(func(string) error)
	// permissionDescCode is the schema descriptor for code field.
	permissionDescCode := permissionFields[1].Descriptor()
	// permission.CodeValidator is a validator for the "code" field. It is called by the builders before save.
	permission.CodeValidator = permissionDescCode.Validators[0].(func(string) error)
	// permissionDescStatus is the schema descriptor for status field.
	permissionDescStatus := permissionFields[3].Descriptor()
	// permission.DefaultStatus holds the default value on creation for the status field.
	permission.DefaultStatus = permissionDescStatus.Default.(int)
	resourceFields := schema.Resource{}.Fields()
	_ = resourceFields
	// resourceDescName is the schema descriptor for name field.
	resourceDescName := resourceFields[0].Descriptor()
	// resource.NameValidator is a validator for the "name" field. It is called by the builders before save.
	resource.NameValidator = resourceDescName.Validators[0].(func(string) error)
	// resourceDescCode is the schema descriptor for code field.
	resourceDescCode := resourceFields[1].Descriptor()
	// resource.CodeValidator is a validator for the "code" field. It is called by the builders before save.
	resource.CodeValidator = resourceDescCode.Validators[0].(func(string) error)
	// resourceDescPath is the schema descriptor for path field.
	resourceDescPath := resourceFields[2].Descriptor()
	// resource.PathValidator is a validator for the "path" field. It is called by the builders before save.
	resource.PathValidator = resourceDescPath.Validators[0].(func(string) error)
	// resourceDescType is the schema descriptor for type field.
	resourceDescType := resourceFields[3].Descriptor()
	// resource.TypeValidator is a validator for the "type" field. It is called by the builders before save.
	resource.TypeValidator = resourceDescType.Validators[0].(func(string) error)
	// resourceDescStatus is the schema descriptor for status field.
	resourceDescStatus := resourceFields[5].Descriptor()
	// resource.DefaultStatus holds the default value on creation for the status field.
	resource.DefaultStatus = resourceDescStatus.Default.(int)
	roleFields := schema.Role{}.Fields()
	_ = roleFields
	// roleDescName is the schema descriptor for name field.
	roleDescName := roleFields[0].Descriptor()
	// role.NameValidator is a validator for the "name" field. It is called by the builders before save.
	role.NameValidator = roleDescName.Validators[0].(func(string) error)
	// roleDescCode is the schema descriptor for code field.
	roleDescCode := roleFields[1].Descriptor()
	// role.CodeValidator is a validator for the "code" field. It is called by the builders before save.
	role.CodeValidator = roleDescCode.Validators[0].(func(string) error)
	// roleDescStatus is the schema descriptor for status field.
	roleDescStatus := roleFields[3].Descriptor()
	// role.DefaultStatus holds the default value on creation for the status field.
	role.DefaultStatus = roleDescStatus.Default.(int)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescStatus is the schema descriptor for status field.
	userDescStatus := userFields[9].Descriptor()
	// user.DefaultStatus holds the default value on creation for the status field.
	user.DefaultStatus = userDescStatus.Default.(int)
}