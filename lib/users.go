package jc

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/TheJumpCloud/jcapi"
)

// ListUsers will return all JumpCLoud users
func ListUsers(attributes []string, query, outputFormat string) {
	// Grab all system users (with their tags if this is a Tags org):
	userList, err := apiClientV1.GetSystemUsers(!isGroups)
	if err != nil {
		log.Fatalf("Could not read system users, err='%s'\n", err)
	}

	outputData(outputFormat, query, userList)

}

// UserAttributeMatches will return attributes for a selected user
func UserAttributeMatches(userName, attributeName, attributeValue string) (bool, error) {
	user, err := getUserByName(userName)
	if err != nil {
		return false, err
	}
	for _, attribute := range user.Attributes {
		if attribute.Name == attributeName {
			return strings.Contains(strings.ToLower(attribute.Value), strings.ToLower(attributeValue)), nil
		}
	}

	return false, fmt.Errorf("attribute not found: %s", attributeName)

}

// UserAttributes will return attributes for a selected user
func UserAttributes(userName string) error {
	user, err := getUserByName(userName)
	if err != nil {
		return err
	}

	var header, data []string
	for _, attribute := range user.Attributes {
		header = append(header, attribute.Name)
		data = append(data, attribute.Value)
	}

	outputData("table", "[].attributes[]", []jcapi.JCUser{user})
	return nil
}

// GetUserAttribute will return the attribute value for user
func GetUserAttribute(userName, key string) (attr string, err error) {
	user, err := getUserByName(userName)
	if err != nil {
		return attr, err
	}

	for _, attribute := range user.Attributes {
		if attribute.Name == key {
			attr = attribute.Value
			continue
		}
	}
	if attr == "" {
		return attr, fmt.Errorf("attribute not found: %s", key)
	}
	return
}

// DeleteUserAttribute will delete a user attribute
func DeleteUserAttribute(userName, key string) (err error) {
	user, err := getUserByName(userName)
	if err != nil {
		return
	}

	for i, attribute := range user.Attributes {
		if attribute.Name == key {

			copy(user.Attributes[i:], user.Attributes[i+1:])                  // Shift a[i+1:] left one index.
			user.Attributes[len(user.Attributes)-1] = jcapi.JCUserAttribute{} // Erase last element (write zero value).
			user.Attributes = user.Attributes[:len(user.Attributes)-1]        // Truncate slice.

			continue
		}
	}
	_, err = apiClientV1.AddUpdateUser(jcapi.Update, user)
	return err
}

// SetUserAttribute will set attribute value for a user
func SetUserAttribute(userName, key, value string) (err error) {
	user, err := getUserByName(userName)
	if err != nil {
		return err
	}

	var attrExists bool
	for _, attribute := range user.Attributes {
		if attribute.Name == key {
			attribute.Value = value
			attrExists = true
			continue
		}
	}

	if !attrExists {
		user.Attributes = append(user.Attributes, jcapi.JCUserAttribute{
			Name:  key,
			Value: value,
		})
	}

	_, err = apiClientV1.AddUpdateUser(jcapi.Update, user)
	return err

}

// SetUserProperties will set the property value for a user
func SetUserProperties(userName string, properties []string) (err error) {
	user, err := getUserByName(userName)
	if err != nil {
		return err
	}

	s := reflect.ValueOf(&user).Elem()
	t := reflect.TypeOf(user)

	for _, p := range properties {
		ss := strings.SplitN(p, "=", 2)
		propertyName, propertyValue := ss[0], ss[1]

		// Iterate over all available fields and read the tag value
		for i := 0; i < t.NumField(); i++ {

			// Get the field, returns https://golang.org/pkg/reflect/#StructField
			field := t.Field(i)

			// Get the field tag value
			tag := field.Tag.Get("json")

			if tag == propertyName {
				f := s.FieldByName(field.Name)
				switch field.Type.Kind() {
				case reflect.Bool:
					b1, _ := strconv.ParseBool(propertyValue)
					f.SetBool(b1)
				case reflect.String:
					f.SetString(propertyValue)
				}
			}
		}
	}

	_, err = apiClientV1.AddUpdateUser(jcapi.Update, user)
	return err

}

func getUserByName(userName string) (jcapi.JCUser, error) {
	userList, err := apiClientV1.GetSystemUsers(!isGroups)
	if err != nil {
		return jcapi.JCUser{}, err
	}
	for _, user := range userList {
		if user.UserName == userName {
			return user, nil
		}
	}
	return jcapi.JCUser{}, fmt.Errorf("unable to find jumpcloud user with name %s", userName)
}

func userListToString(userList []jcapi.JCUser, attributes []string) (userListString [][]string) {
	for _, user := range userList {

		values := []string{
			user.Id,
			user.UserName,
			user.FirstName,
			user.LastName,
			user.Email,
			user.Uid,
			user.Gid,
			fmt.Sprintf("%t", user.Activated),
			fmt.Sprintf("%t", user.PasswordExpired),
		}

		for _, attributeName := range attributes {
			values = append(values, userAttribute(user, attributeName))
		}

		userListString = append(userListString, values)
	}

	return userListString
}

func userAttribute(user jcapi.JCUser, attributeName string) string {
	for _, attribute := range user.Attributes {
		if attribute.Name == attributeName {
			return attribute.Value
		}
	}
	return ""
}
