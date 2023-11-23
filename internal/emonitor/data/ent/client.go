// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"vine-template-rpc/internal/emonitor/data/ent/migrate"

	"vine-template-rpc/internal/emonitor/data/ent/car"
	"vine-template-rpc/internal/emonitor/data/ent/customer"
	"vine-template-rpc/internal/emonitor/data/ent/equipment"
	"vine-template-rpc/internal/emonitor/data/ent/site"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Car is the client for interacting with the Car builders.
	Car *CarClient
	// Customer is the client for interacting with the Customer builders.
	Customer *CustomerClient
	// Equipment is the client for interacting with the Equipment builders.
	Equipment *EquipmentClient
	// Site is the client for interacting with the Site builders.
	Site *SiteClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Car = NewCarClient(c.config)
	c.Customer = NewCustomerClient(c.config)
	c.Equipment = NewEquipmentClient(c.config)
	c.Site = NewSiteClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
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

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:       ctx,
		config:    cfg,
		Car:       NewCarClient(cfg),
		Customer:  NewCustomerClient(cfg),
		Equipment: NewEquipmentClient(cfg),
		Site:      NewSiteClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:       ctx,
		config:    cfg,
		Car:       NewCarClient(cfg),
		Customer:  NewCustomerClient(cfg),
		Equipment: NewEquipmentClient(cfg),
		Site:      NewSiteClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Car.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
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
	c.Car.Use(hooks...)
	c.Customer.Use(hooks...)
	c.Equipment.Use(hooks...)
	c.Site.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Car.Intercept(interceptors...)
	c.Customer.Intercept(interceptors...)
	c.Equipment.Intercept(interceptors...)
	c.Site.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *CarMutation:
		return c.Car.mutate(ctx, m)
	case *CustomerMutation:
		return c.Customer.mutate(ctx, m)
	case *EquipmentMutation:
		return c.Equipment.mutate(ctx, m)
	case *SiteMutation:
		return c.Site.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// CarClient is a client for the Car schema.
type CarClient struct {
	config
}

// NewCarClient returns a client for the Car from the given config.
func NewCarClient(c config) *CarClient {
	return &CarClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `car.Hooks(f(g(h())))`.
func (c *CarClient) Use(hooks ...Hook) {
	c.hooks.Car = append(c.hooks.Car, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `car.Intercept(f(g(h())))`.
func (c *CarClient) Intercept(interceptors ...Interceptor) {
	c.inters.Car = append(c.inters.Car, interceptors...)
}

// Create returns a builder for creating a Car entity.
func (c *CarClient) Create() *CarCreate {
	mutation := newCarMutation(c.config, OpCreate)
	return &CarCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Car entities.
func (c *CarClient) CreateBulk(builders ...*CarCreate) *CarCreateBulk {
	return &CarCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *CarClient) MapCreateBulk(slice any, setFunc func(*CarCreate, int)) *CarCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &CarCreateBulk{err: fmt.Errorf("calling to CarClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*CarCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &CarCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Car.
func (c *CarClient) Update() *CarUpdate {
	mutation := newCarMutation(c.config, OpUpdate)
	return &CarUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CarClient) UpdateOne(ca *Car) *CarUpdateOne {
	mutation := newCarMutation(c.config, OpUpdateOne, withCar(ca))
	return &CarUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CarClient) UpdateOneID(id int) *CarUpdateOne {
	mutation := newCarMutation(c.config, OpUpdateOne, withCarID(id))
	return &CarUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Car.
func (c *CarClient) Delete() *CarDelete {
	mutation := newCarMutation(c.config, OpDelete)
	return &CarDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CarClient) DeleteOne(ca *Car) *CarDeleteOne {
	return c.DeleteOneID(ca.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CarClient) DeleteOneID(id int) *CarDeleteOne {
	builder := c.Delete().Where(car.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CarDeleteOne{builder}
}

// Query returns a query builder for Car.
func (c *CarClient) Query() *CarQuery {
	return &CarQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeCar},
		inters: c.Interceptors(),
	}
}

// Get returns a Car entity by its id.
func (c *CarClient) Get(ctx context.Context, id int) (*Car, error) {
	return c.Query().Where(car.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CarClient) GetX(ctx context.Context, id int) *Car {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CarClient) Hooks() []Hook {
	return c.hooks.Car
}

// Interceptors returns the client interceptors.
func (c *CarClient) Interceptors() []Interceptor {
	return c.inters.Car
}

func (c *CarClient) mutate(ctx context.Context, m *CarMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&CarCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&CarUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&CarUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&CarDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Car mutation op: %q", m.Op())
	}
}

// CustomerClient is a client for the Customer schema.
type CustomerClient struct {
	config
}

// NewCustomerClient returns a client for the Customer from the given config.
func NewCustomerClient(c config) *CustomerClient {
	return &CustomerClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `customer.Hooks(f(g(h())))`.
func (c *CustomerClient) Use(hooks ...Hook) {
	c.hooks.Customer = append(c.hooks.Customer, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `customer.Intercept(f(g(h())))`.
func (c *CustomerClient) Intercept(interceptors ...Interceptor) {
	c.inters.Customer = append(c.inters.Customer, interceptors...)
}

// Create returns a builder for creating a Customer entity.
func (c *CustomerClient) Create() *CustomerCreate {
	mutation := newCustomerMutation(c.config, OpCreate)
	return &CustomerCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Customer entities.
func (c *CustomerClient) CreateBulk(builders ...*CustomerCreate) *CustomerCreateBulk {
	return &CustomerCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *CustomerClient) MapCreateBulk(slice any, setFunc func(*CustomerCreate, int)) *CustomerCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &CustomerCreateBulk{err: fmt.Errorf("calling to CustomerClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*CustomerCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &CustomerCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Customer.
func (c *CustomerClient) Update() *CustomerUpdate {
	mutation := newCustomerMutation(c.config, OpUpdate)
	return &CustomerUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CustomerClient) UpdateOne(cu *Customer) *CustomerUpdateOne {
	mutation := newCustomerMutation(c.config, OpUpdateOne, withCustomer(cu))
	return &CustomerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CustomerClient) UpdateOneID(id int) *CustomerUpdateOne {
	mutation := newCustomerMutation(c.config, OpUpdateOne, withCustomerID(id))
	return &CustomerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Customer.
func (c *CustomerClient) Delete() *CustomerDelete {
	mutation := newCustomerMutation(c.config, OpDelete)
	return &CustomerDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CustomerClient) DeleteOne(cu *Customer) *CustomerDeleteOne {
	return c.DeleteOneID(cu.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CustomerClient) DeleteOneID(id int) *CustomerDeleteOne {
	builder := c.Delete().Where(customer.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CustomerDeleteOne{builder}
}

// Query returns a query builder for Customer.
func (c *CustomerClient) Query() *CustomerQuery {
	return &CustomerQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeCustomer},
		inters: c.Interceptors(),
	}
}

// Get returns a Customer entity by its id.
func (c *CustomerClient) Get(ctx context.Context, id int) (*Customer, error) {
	return c.Query().Where(customer.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CustomerClient) GetX(ctx context.Context, id int) *Customer {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CustomerClient) Hooks() []Hook {
	return c.hooks.Customer
}

// Interceptors returns the client interceptors.
func (c *CustomerClient) Interceptors() []Interceptor {
	return c.inters.Customer
}

func (c *CustomerClient) mutate(ctx context.Context, m *CustomerMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&CustomerCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&CustomerUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&CustomerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&CustomerDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Customer mutation op: %q", m.Op())
	}
}

// EquipmentClient is a client for the Equipment schema.
type EquipmentClient struct {
	config
}

// NewEquipmentClient returns a client for the Equipment from the given config.
func NewEquipmentClient(c config) *EquipmentClient {
	return &EquipmentClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `equipment.Hooks(f(g(h())))`.
func (c *EquipmentClient) Use(hooks ...Hook) {
	c.hooks.Equipment = append(c.hooks.Equipment, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `equipment.Intercept(f(g(h())))`.
func (c *EquipmentClient) Intercept(interceptors ...Interceptor) {
	c.inters.Equipment = append(c.inters.Equipment, interceptors...)
}

// Create returns a builder for creating a Equipment entity.
func (c *EquipmentClient) Create() *EquipmentCreate {
	mutation := newEquipmentMutation(c.config, OpCreate)
	return &EquipmentCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Equipment entities.
func (c *EquipmentClient) CreateBulk(builders ...*EquipmentCreate) *EquipmentCreateBulk {
	return &EquipmentCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *EquipmentClient) MapCreateBulk(slice any, setFunc func(*EquipmentCreate, int)) *EquipmentCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &EquipmentCreateBulk{err: fmt.Errorf("calling to EquipmentClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*EquipmentCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &EquipmentCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Equipment.
func (c *EquipmentClient) Update() *EquipmentUpdate {
	mutation := newEquipmentMutation(c.config, OpUpdate)
	return &EquipmentUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *EquipmentClient) UpdateOne(e *Equipment) *EquipmentUpdateOne {
	mutation := newEquipmentMutation(c.config, OpUpdateOne, withEquipment(e))
	return &EquipmentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *EquipmentClient) UpdateOneID(id uuid.UUID) *EquipmentUpdateOne {
	mutation := newEquipmentMutation(c.config, OpUpdateOne, withEquipmentID(id))
	return &EquipmentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Equipment.
func (c *EquipmentClient) Delete() *EquipmentDelete {
	mutation := newEquipmentMutation(c.config, OpDelete)
	return &EquipmentDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *EquipmentClient) DeleteOne(e *Equipment) *EquipmentDeleteOne {
	return c.DeleteOneID(e.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *EquipmentClient) DeleteOneID(id uuid.UUID) *EquipmentDeleteOne {
	builder := c.Delete().Where(equipment.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &EquipmentDeleteOne{builder}
}

// Query returns a query builder for Equipment.
func (c *EquipmentClient) Query() *EquipmentQuery {
	return &EquipmentQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeEquipment},
		inters: c.Interceptors(),
	}
}

// Get returns a Equipment entity by its id.
func (c *EquipmentClient) Get(ctx context.Context, id uuid.UUID) (*Equipment, error) {
	return c.Query().Where(equipment.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *EquipmentClient) GetX(ctx context.Context, id uuid.UUID) *Equipment {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *EquipmentClient) Hooks() []Hook {
	return c.hooks.Equipment
}

// Interceptors returns the client interceptors.
func (c *EquipmentClient) Interceptors() []Interceptor {
	return c.inters.Equipment
}

func (c *EquipmentClient) mutate(ctx context.Context, m *EquipmentMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&EquipmentCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&EquipmentUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&EquipmentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&EquipmentDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Equipment mutation op: %q", m.Op())
	}
}

// SiteClient is a client for the Site schema.
type SiteClient struct {
	config
}

// NewSiteClient returns a client for the Site from the given config.
func NewSiteClient(c config) *SiteClient {
	return &SiteClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `site.Hooks(f(g(h())))`.
func (c *SiteClient) Use(hooks ...Hook) {
	c.hooks.Site = append(c.hooks.Site, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `site.Intercept(f(g(h())))`.
func (c *SiteClient) Intercept(interceptors ...Interceptor) {
	c.inters.Site = append(c.inters.Site, interceptors...)
}

// Create returns a builder for creating a Site entity.
func (c *SiteClient) Create() *SiteCreate {
	mutation := newSiteMutation(c.config, OpCreate)
	return &SiteCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Site entities.
func (c *SiteClient) CreateBulk(builders ...*SiteCreate) *SiteCreateBulk {
	return &SiteCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *SiteClient) MapCreateBulk(slice any, setFunc func(*SiteCreate, int)) *SiteCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &SiteCreateBulk{err: fmt.Errorf("calling to SiteClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*SiteCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &SiteCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Site.
func (c *SiteClient) Update() *SiteUpdate {
	mutation := newSiteMutation(c.config, OpUpdate)
	return &SiteUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SiteClient) UpdateOne(s *Site) *SiteUpdateOne {
	mutation := newSiteMutation(c.config, OpUpdateOne, withSite(s))
	return &SiteUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SiteClient) UpdateOneID(id uuid.UUID) *SiteUpdateOne {
	mutation := newSiteMutation(c.config, OpUpdateOne, withSiteID(id))
	return &SiteUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Site.
func (c *SiteClient) Delete() *SiteDelete {
	mutation := newSiteMutation(c.config, OpDelete)
	return &SiteDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *SiteClient) DeleteOne(s *Site) *SiteDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *SiteClient) DeleteOneID(id uuid.UUID) *SiteDeleteOne {
	builder := c.Delete().Where(site.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SiteDeleteOne{builder}
}

// Query returns a query builder for Site.
func (c *SiteClient) Query() *SiteQuery {
	return &SiteQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeSite},
		inters: c.Interceptors(),
	}
}

// Get returns a Site entity by its id.
func (c *SiteClient) Get(ctx context.Context, id uuid.UUID) (*Site, error) {
	return c.Query().Where(site.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SiteClient) GetX(ctx context.Context, id uuid.UUID) *Site {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *SiteClient) Hooks() []Hook {
	return c.hooks.Site
}

// Interceptors returns the client interceptors.
func (c *SiteClient) Interceptors() []Interceptor {
	return c.inters.Site
}

func (c *SiteClient) mutate(ctx context.Context, m *SiteMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&SiteCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&SiteUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&SiteUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&SiteDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Site mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Car, Customer, Equipment, Site []ent.Hook
	}
	inters struct {
		Car, Customer, Equipment, Site []ent.Interceptor
	}
)
