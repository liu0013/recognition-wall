// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"recognition-wall/internal/data/ent/feedback"
	"recognition-wall/internal/data/ent/like"
	"recognition-wall/internal/data/ent/predicate"
	"sync"
	"time"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeFeedback = "Feedback"
	TypeLike     = "Like"
)

// FeedbackMutation represents an operation that mutates the Feedback nodes in the graph.
type FeedbackMutation struct {
	config
	op            Op
	typ           string
	id            *int64
	name          *string
	content       *string
	created_at    *time.Time
	updated_at    *time.Time
	clearedFields map[string]struct{}
	like          *int64
	clearedlike   bool
	done          bool
	oldValue      func(context.Context) (*Feedback, error)
	predicates    []predicate.Feedback
}

var _ ent.Mutation = (*FeedbackMutation)(nil)

// feedbackOption allows management of the mutation configuration using functional options.
type feedbackOption func(*FeedbackMutation)

// newFeedbackMutation creates new mutation for the Feedback entity.
func newFeedbackMutation(c config, op Op, opts ...feedbackOption) *FeedbackMutation {
	m := &FeedbackMutation{
		config:        c,
		op:            op,
		typ:           TypeFeedback,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withFeedbackID sets the ID field of the mutation.
func withFeedbackID(id int64) feedbackOption {
	return func(m *FeedbackMutation) {
		var (
			err   error
			once  sync.Once
			value *Feedback
		)
		m.oldValue = func(ctx context.Context) (*Feedback, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Feedback.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withFeedback sets the old Feedback of the mutation.
func withFeedback(node *Feedback) feedbackOption {
	return func(m *FeedbackMutation) {
		m.oldValue = func(context.Context) (*Feedback, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m FeedbackMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m FeedbackMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Feedback entities.
func (m *FeedbackMutation) SetID(id int64) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *FeedbackMutation) ID() (id int64, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetName sets the "name" field.
func (m *FeedbackMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *FeedbackMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Feedback entity.
// If the Feedback object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FeedbackMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *FeedbackMutation) ResetName() {
	m.name = nil
}

// SetContent sets the "content" field.
func (m *FeedbackMutation) SetContent(s string) {
	m.content = &s
}

// Content returns the value of the "content" field in the mutation.
func (m *FeedbackMutation) Content() (r string, exists bool) {
	v := m.content
	if v == nil {
		return
	}
	return *v, true
}

// OldContent returns the old "content" field's value of the Feedback entity.
// If the Feedback object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FeedbackMutation) OldContent(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldContent is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldContent requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldContent: %w", err)
	}
	return oldValue.Content, nil
}

// ResetContent resets all changes to the "content" field.
func (m *FeedbackMutation) ResetContent() {
	m.content = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *FeedbackMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *FeedbackMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Feedback entity.
// If the Feedback object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FeedbackMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *FeedbackMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *FeedbackMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *FeedbackMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Feedback entity.
// If the Feedback object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FeedbackMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *FeedbackMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetLikeID sets the "like" edge to the Like entity by id.
func (m *FeedbackMutation) SetLikeID(id int64) {
	m.like = &id
}

// ClearLike clears the "like" edge to the Like entity.
func (m *FeedbackMutation) ClearLike() {
	m.clearedlike = true
}

// LikeCleared reports if the "like" edge to the Like entity was cleared.
func (m *FeedbackMutation) LikeCleared() bool {
	return m.clearedlike
}

// LikeID returns the "like" edge ID in the mutation.
func (m *FeedbackMutation) LikeID() (id int64, exists bool) {
	if m.like != nil {
		return *m.like, true
	}
	return
}

// LikeIDs returns the "like" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// LikeID instead. It exists only for internal usage by the builders.
func (m *FeedbackMutation) LikeIDs() (ids []int64) {
	if id := m.like; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetLike resets all changes to the "like" edge.
func (m *FeedbackMutation) ResetLike() {
	m.like = nil
	m.clearedlike = false
}

// Where appends a list predicates to the FeedbackMutation builder.
func (m *FeedbackMutation) Where(ps ...predicate.Feedback) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *FeedbackMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Feedback).
func (m *FeedbackMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *FeedbackMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.name != nil {
		fields = append(fields, feedback.FieldName)
	}
	if m.content != nil {
		fields = append(fields, feedback.FieldContent)
	}
	if m.created_at != nil {
		fields = append(fields, feedback.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, feedback.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *FeedbackMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case feedback.FieldName:
		return m.Name()
	case feedback.FieldContent:
		return m.Content()
	case feedback.FieldCreatedAt:
		return m.CreatedAt()
	case feedback.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *FeedbackMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case feedback.FieldName:
		return m.OldName(ctx)
	case feedback.FieldContent:
		return m.OldContent(ctx)
	case feedback.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case feedback.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Feedback field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *FeedbackMutation) SetField(name string, value ent.Value) error {
	switch name {
	case feedback.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case feedback.FieldContent:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetContent(v)
		return nil
	case feedback.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case feedback.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Feedback field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *FeedbackMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *FeedbackMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *FeedbackMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Feedback numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *FeedbackMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *FeedbackMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *FeedbackMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Feedback nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *FeedbackMutation) ResetField(name string) error {
	switch name {
	case feedback.FieldName:
		m.ResetName()
		return nil
	case feedback.FieldContent:
		m.ResetContent()
		return nil
	case feedback.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case feedback.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown Feedback field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *FeedbackMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.like != nil {
		edges = append(edges, feedback.EdgeLike)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *FeedbackMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case feedback.EdgeLike:
		if id := m.like; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *FeedbackMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *FeedbackMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *FeedbackMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedlike {
		edges = append(edges, feedback.EdgeLike)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *FeedbackMutation) EdgeCleared(name string) bool {
	switch name {
	case feedback.EdgeLike:
		return m.clearedlike
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *FeedbackMutation) ClearEdge(name string) error {
	switch name {
	case feedback.EdgeLike:
		m.ClearLike()
		return nil
	}
	return fmt.Errorf("unknown Feedback unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *FeedbackMutation) ResetEdge(name string) error {
	switch name {
	case feedback.EdgeLike:
		m.ResetLike()
		return nil
	}
	return fmt.Errorf("unknown Feedback edge %s", name)
}

// LikeMutation represents an operation that mutates the Like nodes in the graph.
type LikeMutation struct {
	config
	op              Op
	typ             string
	id              *int64
	count           *int32
	addcount        *int32
	clearedFields   map[string]struct{}
	feedback        *int64
	clearedfeedback bool
	done            bool
	oldValue        func(context.Context) (*Like, error)
	predicates      []predicate.Like
}

var _ ent.Mutation = (*LikeMutation)(nil)

// likeOption allows management of the mutation configuration using functional options.
type likeOption func(*LikeMutation)

// newLikeMutation creates new mutation for the Like entity.
func newLikeMutation(c config, op Op, opts ...likeOption) *LikeMutation {
	m := &LikeMutation{
		config:        c,
		op:            op,
		typ:           TypeLike,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withLikeID sets the ID field of the mutation.
func withLikeID(id int64) likeOption {
	return func(m *LikeMutation) {
		var (
			err   error
			once  sync.Once
			value *Like
		)
		m.oldValue = func(ctx context.Context) (*Like, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Like.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withLike sets the old Like of the mutation.
func withLike(node *Like) likeOption {
	return func(m *LikeMutation) {
		m.oldValue = func(context.Context) (*Like, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m LikeMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m LikeMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Like entities.
func (m *LikeMutation) SetID(id int64) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *LikeMutation) ID() (id int64, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetCount sets the "count" field.
func (m *LikeMutation) SetCount(i int32) {
	m.count = &i
	m.addcount = nil
}

// Count returns the value of the "count" field in the mutation.
func (m *LikeMutation) Count() (r int32, exists bool) {
	v := m.count
	if v == nil {
		return
	}
	return *v, true
}

// OldCount returns the old "count" field's value of the Like entity.
// If the Like object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *LikeMutation) OldCount(ctx context.Context) (v int32, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCount is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCount requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCount: %w", err)
	}
	return oldValue.Count, nil
}

// AddCount adds i to the "count" field.
func (m *LikeMutation) AddCount(i int32) {
	if m.addcount != nil {
		*m.addcount += i
	} else {
		m.addcount = &i
	}
}

// AddedCount returns the value that was added to the "count" field in this mutation.
func (m *LikeMutation) AddedCount() (r int32, exists bool) {
	v := m.addcount
	if v == nil {
		return
	}
	return *v, true
}

// ResetCount resets all changes to the "count" field.
func (m *LikeMutation) ResetCount() {
	m.count = nil
	m.addcount = nil
}

// SetFeedbackID sets the "feedback" edge to the Feedback entity by id.
func (m *LikeMutation) SetFeedbackID(id int64) {
	m.feedback = &id
}

// ClearFeedback clears the "feedback" edge to the Feedback entity.
func (m *LikeMutation) ClearFeedback() {
	m.clearedfeedback = true
}

// FeedbackCleared reports if the "feedback" edge to the Feedback entity was cleared.
func (m *LikeMutation) FeedbackCleared() bool {
	return m.clearedfeedback
}

// FeedbackID returns the "feedback" edge ID in the mutation.
func (m *LikeMutation) FeedbackID() (id int64, exists bool) {
	if m.feedback != nil {
		return *m.feedback, true
	}
	return
}

// FeedbackIDs returns the "feedback" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// FeedbackID instead. It exists only for internal usage by the builders.
func (m *LikeMutation) FeedbackIDs() (ids []int64) {
	if id := m.feedback; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetFeedback resets all changes to the "feedback" edge.
func (m *LikeMutation) ResetFeedback() {
	m.feedback = nil
	m.clearedfeedback = false
}

// Where appends a list predicates to the LikeMutation builder.
func (m *LikeMutation) Where(ps ...predicate.Like) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *LikeMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Like).
func (m *LikeMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *LikeMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.count != nil {
		fields = append(fields, like.FieldCount)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *LikeMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case like.FieldCount:
		return m.Count()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *LikeMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case like.FieldCount:
		return m.OldCount(ctx)
	}
	return nil, fmt.Errorf("unknown Like field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *LikeMutation) SetField(name string, value ent.Value) error {
	switch name {
	case like.FieldCount:
		v, ok := value.(int32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCount(v)
		return nil
	}
	return fmt.Errorf("unknown Like field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *LikeMutation) AddedFields() []string {
	var fields []string
	if m.addcount != nil {
		fields = append(fields, like.FieldCount)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *LikeMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case like.FieldCount:
		return m.AddedCount()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *LikeMutation) AddField(name string, value ent.Value) error {
	switch name {
	case like.FieldCount:
		v, ok := value.(int32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddCount(v)
		return nil
	}
	return fmt.Errorf("unknown Like numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *LikeMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *LikeMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *LikeMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Like nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *LikeMutation) ResetField(name string) error {
	switch name {
	case like.FieldCount:
		m.ResetCount()
		return nil
	}
	return fmt.Errorf("unknown Like field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *LikeMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.feedback != nil {
		edges = append(edges, like.EdgeFeedback)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *LikeMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case like.EdgeFeedback:
		if id := m.feedback; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *LikeMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *LikeMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *LikeMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedfeedback {
		edges = append(edges, like.EdgeFeedback)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *LikeMutation) EdgeCleared(name string) bool {
	switch name {
	case like.EdgeFeedback:
		return m.clearedfeedback
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *LikeMutation) ClearEdge(name string) error {
	switch name {
	case like.EdgeFeedback:
		m.ClearFeedback()
		return nil
	}
	return fmt.Errorf("unknown Like unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *LikeMutation) ResetEdge(name string) error {
	switch name {
	case like.EdgeFeedback:
		m.ResetFeedback()
		return nil
	}
	return fmt.Errorf("unknown Like edge %s", name)
}