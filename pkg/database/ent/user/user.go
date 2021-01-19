// Code generated by entc, DO NOT EDIT.

package user

import (
	"time"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldNickname holds the string denoting the nickname field in the database.
	FieldNickname = "nickname"
	// FieldTOTPSecret holds the string denoting the totpsecret field in the database.
	FieldTOTPSecret = "totp_secret"
	// FieldOnline holds the string denoting the online field in the database.
	FieldOnline = "online"
	// FieldEnable holds the string denoting the enable field in the database.
	FieldEnable = "enable"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"

	// Table holds the table name of the user in the database.
	Table = "Users"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUsername,
	FieldPassword,
	FieldNickname,
	FieldTOTPSecret,
	FieldOnline,
	FieldEnable,
	FieldCreatedAt,
	FieldType,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)
