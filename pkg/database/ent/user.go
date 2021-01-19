// Code generated by entc, DO NOT EDIT.

package ent

import (
	"cloud-terminal/pkg/database/ent/user"
	"fmt"
	"strings"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/google/uuid"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Username holds the value of the "Username" field.
	Username string `json:"Username,omitempty"`
	// Password holds the value of the "Password" field.
	Password string `json:"Password,omitempty"`
	// Nickname holds the value of the "Nickname" field.
	Nickname string `json:"Nickname,omitempty"`
	// TOTPSecret holds the value of the "TOTPSecret" field.
	TOTPSecret string `json:"TOTPSecret,omitempty"`
	// Online holds the value of the "Online" field.
	Online bool `json:"Online,omitempty"`
	// Enable holds the value of the "Enable" field.
	Enable bool `json:"Enable,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Type holds the value of the "Type" field.
	Type string `json:"Type,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldOnline, user.FieldEnable:
			values[i] = &sql.NullBool{}
		case user.FieldUsername, user.FieldPassword, user.FieldNickname, user.FieldTOTPSecret, user.FieldType:
			values[i] = &sql.NullString{}
		case user.FieldCreatedAt:
			values[i] = &sql.NullTime{}
		case user.FieldID:
			values[i] = &uuid.UUID{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				u.ID = *value
			}
		case user.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Username", values[i])
			} else if value.Valid {
				u.Username = value.String
			}
		case user.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Password", values[i])
			} else if value.Valid {
				u.Password = value.String
			}
		case user.FieldNickname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Nickname", values[i])
			} else if value.Valid {
				u.Nickname = value.String
			}
		case user.FieldTOTPSecret:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field TOTPSecret", values[i])
			} else if value.Valid {
				u.TOTPSecret = value.String
			}
		case user.FieldOnline:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field Online", values[i])
			} else if value.Valid {
				u.Online = value.Bool
			}
		case user.FieldEnable:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field Enable", values[i])
			} else if value.Valid {
				u.Enable = value.Bool
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case user.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Type", values[i])
			} else if value.Valid {
				u.Type = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", Username=")
	builder.WriteString(u.Username)
	builder.WriteString(", Password=")
	builder.WriteString(u.Password)
	builder.WriteString(", Nickname=")
	builder.WriteString(u.Nickname)
	builder.WriteString(", TOTPSecret=")
	builder.WriteString(u.TOTPSecret)
	builder.WriteString(", Online=")
	builder.WriteString(fmt.Sprintf("%v", u.Online))
	builder.WriteString(", Enable=")
	builder.WriteString(fmt.Sprintf("%v", u.Enable))
	builder.WriteString(", created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", Type=")
	builder.WriteString(u.Type)
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}