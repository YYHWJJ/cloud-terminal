// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/willie-lin/cloud-terminal/ent/migrate"

	"github.com/willie-lin/cloud-terminal/ent/usergroup"

	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// UserGroup is the client for interacting with the UserGroup builders.
	UserGroup *UserGroupClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.UserGroup = NewUserGroupClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		ctx:       ctx,
		config:    cfg,
		UserGroup: NewUserGroupClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(*sql.Driver).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: &txDriver{tx: tx, drv: c.driver}, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		config:    cfg,
		UserGroup: NewUserGroupClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		UserGroup.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.UserGroup.Use(hooks...)
}

// UserGroupClient is a client for the UserGroup schema.
type UserGroupClient struct {
	config
}

// NewUserGroupClient returns a client for the UserGroup from the given config.
func NewUserGroupClient(c config) *UserGroupClient {
	return &UserGroupClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `usergroup.Hooks(f(g(h())))`.
func (c *UserGroupClient) Use(hooks ...Hook) {
	c.hooks.UserGroup = append(c.hooks.UserGroup, hooks...)
}

// Create returns a create builder for UserGroup.
func (c *UserGroupClient) Create() *UserGroupCreate {
	mutation := newUserGroupMutation(c.config, OpCreate)
	return &UserGroupCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of UserGroup entities.
func (c *UserGroupClient) CreateBulk(builders ...*UserGroupCreate) *UserGroupCreateBulk {
	return &UserGroupCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for UserGroup.
func (c *UserGroupClient) Update() *UserGroupUpdate {
	mutation := newUserGroupMutation(c.config, OpUpdate)
	return &UserGroupUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserGroupClient) UpdateOne(ug *UserGroup) *UserGroupUpdateOne {
	mutation := newUserGroupMutation(c.config, OpUpdateOne, withUserGroup(ug))
	return &UserGroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserGroupClient) UpdateOneID(id int) *UserGroupUpdateOne {
	mutation := newUserGroupMutation(c.config, OpUpdateOne, withUserGroupID(id))
	return &UserGroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for UserGroup.
func (c *UserGroupClient) Delete() *UserGroupDelete {
	mutation := newUserGroupMutation(c.config, OpDelete)
	return &UserGroupDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserGroupClient) DeleteOne(ug *UserGroup) *UserGroupDeleteOne {
	return c.DeleteOneID(ug.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserGroupClient) DeleteOneID(id int) *UserGroupDeleteOne {
	builder := c.Delete().Where(usergroup.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserGroupDeleteOne{builder}
}

// Query returns a query builder for UserGroup.
func (c *UserGroupClient) Query() *UserGroupQuery {
	return &UserGroupQuery{config: c.config}
}

// Get returns a UserGroup entity by its id.
func (c *UserGroupClient) Get(ctx context.Context, id int) (*UserGroup, error) {
	return c.Query().Where(usergroup.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserGroupClient) GetX(ctx context.Context, id int) *UserGroup {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *UserGroupClient) Hooks() []Hook {
	return c.hooks.UserGroup
}
