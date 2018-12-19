// This file is generated by protoc-gen-persist
// Source File: pb/user.proto
// DO NOT EDIT !

// TXs, Queries, Hooks, TypeMappings, Handlers, Rows, Iters
package pb

import (
	"fmt"
  "time"
	io "io"

  spanner "cloud.google.com/go/spanner"
  "google.golang.org/api/iterator"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	context "golang.org/x/net/context"
	codes "google.golang.org/grpc/codes"
	gstatus "google.golang.org/grpc/status"
)

// WriteHandlers
// RestOf<S>Handlers
type RestOfUServHandlers interface {
  UpdateAllNames(*Empty, UServ_UpdateAllNamesServer) error
}

func (this *UServ_Impl) UpdateAllNames(req *Empty, stream UServ_UpdateAllNamesServer) error {
  return this.opts.HANDLERS.UpdateAllNames(req, stream)
}

// WriteTypeMappigns
type UServ_TypeMappings interface {
	TimestampTimestamp() TimestampTimestampMappingImpl
	SliceStringParam() SliceStringParamMappingImpl
}

type TimestampTimestampMappingImpl interface {
	ToProto(**timestamp.Timestamp) error
	Empty() TimestampTimestampMappingImpl
	ToSpanner(*timestamp.Timestamp) TimestampTimestampMappingImpl
	SpannerScan(src *spanner.GenericColumnValue) error
	SpannerValue() (interface{}, error)
}

type SliceStringParamMappingImpl interface {
	ToProto(**SliceStringParam) error
	Empty() SliceStringParamMappingImpl
  ToSpanner(*SliceStringParam) SliceStringParamMappingImpl
	SpannerScan(src *spanner.GenericColumnValue) error
	SpannerValue() (interface{}, error)
}

type UServ_Hooks interface {
	InsertUsersBeforeHook(context.Context, *User) (*Empty, error)
	InsertUsersAfterHook(context.Context, *User, *Empty) error
	GetAllUsersBeforeHook(context.Context, *Empty) ([]*User, error)
	GetAllUsersAfterHook(context.Context, *Empty, *User) error
}
type alwaysScanner struct {
	i *interface{}
}

func (s *alwaysScanner) Scan(src interface{}) error {
	s.i = &src
	return nil
}

type Result interface {
  LastInsertId() (int64, error)
  RowsAffected() (int64, error)
}
type SpannerResult struct {
  iter *spanner.RowIterator
}

func (sr *SpannerResult) LastInsertId() (int64, error) {
  // sr.iter.QueryStats or sr.iter.QueryPlan
  return -1, nil
}
func (sr *SpannerResult) RowsAffected() (int64, error) {
  // Execution statistics for the query. Available after RowIterator.Next returns iterator.Done
  return sr.iter.RowCount, nil
}

type Runnable interface {
  ReadWriteTransaction(context.Context, func(context.Context, *spanner.ReadWriteTransaction) error)(time.Time, error)
}

// func DefaultClientStreamingPersistTx(ctx context.Context, db *spanner.Client) (PersistTx, error) {
//   // TODO dont need if spanner is taking care of it
// 	return db.BeginTx(ctx, nil)
// }
// func DefaultServerStreamingPersistTx(ctx context.Context, db *spanner.Client) (PersistTx, error) {
// 	return NopPersistTx(db)
// }
// func DefaultUnaryPersistTx(ctx context.Context, db *spanner.Client) (PersistTx, error) {
// 	return NopPersistTx(db)
// }

// type ignoreTx struct {
// 	r Runnable
// }

// func (this *ignoreTx) Commit() error   { return nil }
// func (this *ignoreTx) Rollback() error { return nil }
// func (this *ignoreTx) QueryContext(ctx context.Context, x string, ys ...interface{}) (*spanner.RowIterator, error) {
// 	return this.r.QueryContext(ctx, x, ys...)
// }
// func (this *ignoreTx) ExecContext(ctx context.Context, x string, ys ...interface{}) (SpannerResult, error) {
// 	return this.r.ExecContext(ctx, x, ys...)
// }

type UServ_QueryOpts struct {
	MAPPINGS UServ_TypeMappings
	db       Runnable
	ctx      context.Context
}

func DefaultUServQueryOpts(db Runnable) UServ_QueryOpts {
	return UServ_QueryOpts{
		db: db,
	}
}

type UServ_Queries struct {
	opts UServ_QueryOpts
}
//type PersistTx interface {
//	Commit() error
//	Rollback() error
//	Runnable
//}

//func (tx *PersistTx) Commit() error {
//  //TODO 
//}
//func (tx *PersistTx) Rollback() error {
//  //TODO 
//}

// func NopPersistTx(r Runnable) (PersistTx, error) {
// 	return &ignoreTx{r}, nil
// }

type UServ_InsertUsersOut interface {
	GetId() int64
	GetName() string
	GetFriends() *Friends
	GetCreatedOn() *timestamp.Timestamp
}

type UServ_InsertUsersRow struct {
	item UServ_InsertUsersOut
	err  error
}

func newUServ_InsertUsersRow(item UServ_InsertUsersOut, err error) *UServ_InsertUsersRow {
	return &UServ_InsertUsersRow{item, err}
}

// Unwrap takes an address to a proto.Message as its only parameter
// Unwrap can only set into output protos of that match method return types + the out option on the query itself
func (this *UServ_InsertUsersRow) Unwrap(pointerToMsg proto.Message) error {
	if this.err != nil {
		return this.err
	}
	// for each known method result
	if o, ok := (pointerToMsg).(*User); ok {
		if o == nil {
			return fmt.Errorf("must initialize *User before giving to Unwrap()")
		}
		res, _ := this.User()
		// set shared fields
		o.Id = res.Id
		o.Name = res.Name
		o.Friends = res.Friends
		o.CreatedOn = res.CreatedOn
		return nil
	}
	if o, ok := (pointerToMsg).(*Friends); ok {
		if o == nil {
			return fmt.Errorf("must initialize *Friends before giving to Unwrap()")
		}
	}

	return nil
}

// one for each Output type of the methods that use this query + the output proto itself

func (this *UServ_InsertUsersRow) User() (*User, error) {
	return &User{
		Id:        this.item.GetId(),
		Name:      this.item.GetName(),
		Friends:   this.item.GetFriends(),
		CreatedOn: this.item.GetCreatedOn(),
	}, nil
}

// just for example
func (this *UServ_InsertUsersRow) Friends() (*Friends, error) {
	return nil, nil
}

// UServPersistQueries returns all the known 'SQL' queires for the 'UServ' service.
func UServPersistQueries(db Runnable, opts ...UServ_QueryOpts) *UServ_Queries {
	var myOpts UServ_QueryOpts
	if len(opts) > 0 {
		myOpts = opts[0]
    myOpts.db = db
	} else {
		myOpts = DefaultUServQueryOpts(db)
	}
	return &UServ_Queries{
		opts: myOpts,
	}
}

// camel case the services query name
// method for every query

// InsertUsersQuery returns a new struct wrapping the current UServ_QueryOpts
// that will perform 'UServ' services 'insert_users_query' on the database
// when executed
func (this *UServ_Queries) InsertUsersQuery(ctx context.Context) *UServ_InsertUsersQuery {
	return &UServ_InsertUsersQuery{
		opts: UServ_QueryOpts{
			MAPPINGS: this.opts.MAPPINGS,
			db:       this.opts.db,
			ctx:      ctx,
		},
	}
}

// I dont know this is a insert query, I only know this is a query
type UServ_InsertUsersQuery struct {
	opts UServ_QueryOpts
	ctx  context.Context
}

func (this *UServ_InsertUsersQuery) QueryInTypeUser()  {}
func (this *UServ_InsertUsersQuery) QueryOutTypeUser() {}

// the main execute function
func (this *UServ_InsertUsersQuery) Execute(x UServ_InsertUsersOut) *UServ_InsertUsersIter {
	var setupErr error
	params := []interface{}{
		func() (out interface{}) {
			out = x.GetId()
			return
		}(),
		func() (out interface{}) {
      out = fmt.Sprintf(`"%s"`, x.GetName())
			return
		}(),
		func() (out interface{}) {
			raw, err := proto.Marshal(x.GetFriends())
      fmt.Println("****",raw)
			if err != nil {
				setupErr = err
			}
			out = fmt.Sprintf(`CAST("%v" as BYTES)`, raw)
			return
		}(),
		func() (out interface{}) {
			mapper := this.opts.MAPPINGS.TimestampTimestamp()
      ts, err := mapper.ToSpanner(x.GetCreatedOn()).SpannerValue()
      if err != nil {
        setupErr = err
      }

      out = fmt.Sprintf(`"%s"`, ts)
			return
		}(),
	}
	result := &UServ_InsertUsersIter{
		tm:  this.opts.MAPPINGS,
		ctx: this.ctx,
	}
	if setupErr != nil {
		result.err = setupErr
		return result
	}

  fmt.Println("params: ", params)
  _, err := this.opts.db.ReadWriteTransaction(this.ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
    stmt := spanner.Statement{
      SQL: fmt.Sprintf("insert into users (id, name, friends, created_on) values (%v, %v, %v, %v);", params...)}
    iter := txn.QueryWithStats(ctx, stmt)
    // err, rowCount := txn.Update(ctx, stmt)
    result.rows = iter
    result.result = SpannerResult{
      iter: iter,
    }
    // TODO consider the effects of calling "defer iter.Stop()" here
    return nil
  })
  result.err = err
	return result
}

//<SERVICE><QUERY (camel)><MESSAGE TYPE>Iter
type UServ_InsertUsersIter struct {
	result SpannerResult
	rows   *spanner.RowIterator
	err    error
	tm     UServ_TypeMappings
	ctx    context.Context
}

func (this *UServ_InsertUsersIter) IterOutTypeUser() {}
func (this *UServ_InsertUsersIter) IterInTypeUser()  {}

// Each performs 'fun' on each row in the result set.
// Each respects the context passed to it.
// It will stop iteration, and returns ctx.Err() if encountered.
// TODO can just call the "Do" method on spanners iterator
func (this *UServ_InsertUsersIter) Each(fun func(*UServ_InsertUsersRow) error) error {
	for {
		select {
		case <-this.ctx.Done():
			return this.ctx.Err()
		default:
			if row, ok := this.Next(); !ok {
				return nil
			} else if err := fun(row); err != nil {
				return err
			}
		}
	}
	return nil
}

// One returns the sole row, or ensures an error if there was not one result when this row is converted
func (this *UServ_InsertUsersIter) One() *UServ_InsertUsersRow {
  first, hasFirst := this.Next()
  _, hasSecond := this.Next()
  if !hasFirst || hasSecond {
    amount := "none"
    if hasSecond {
      amount = "multiple"
    }
    return &UServ_InsertUsersRow{err: fmt.Errorf("expected exactly 1 result from query 'InsertUsers' found %s", amount)}
  }
  return first
}

// Zero returns an error if there were any rows in the result
func (this *UServ_InsertUsersIter) Zero() error {
	if _, ok := this.Next(); ok {
		return fmt.Errorf("expected exactly 0 results from query 'InsertUsers'")
	}
	return nil
}

// Next returns the next scanned row out of the database, or (nil, false) if there are no more rows
func (this *UServ_InsertUsersIter) Next() (*UServ_InsertUsersRow, bool) {
	if this.rows == nil || this.err == io.EOF {
		return nil, false
	} else if this.err != nil {
		err := this.err
		this.err = io.EOF
		return &UServ_InsertUsersRow{err: err}, true
	}
  row, err := this.rows.Next()
  if err == iterator.Done {
    if this.err == nil {
      this.err = io.EOF
      return nil, false
    }
  }
  if err != nil {
		return &UServ_InsertUsersRow{err: err}, true
  }

  var id int64
  if err := row.ColumnByName("id", &id); err != nil {
      return &UServ_InsertUsersRow{err: fmt.Errorf("cant convert db column id to protobuf go type int64")}, true
  }
  var name string
  if err := row.ColumnByName("name", &name); err != nil {
      return &UServ_InsertUsersRow{err: fmt.Errorf("cant convert db column name to protobuf go type string")}, true
  }
  var friends *Friends
  if err := row.ColumnByName("friends", &friends); err != nil {
      return &UServ_InsertUsersRow{err: fmt.Errorf("cant convert db column friends to protobuf go type *Friends")}, true
  }
  var created_on *timestamp.Timestamp
  if err := row.ColumnByName("created_on", &created_on); err != nil {
      return &UServ_InsertUsersRow{err: fmt.Errorf("could not convert mapped db column created_on to type on User.CreatedOn: %v", err)}, true
  }

  return &UServ_InsertUsersRow{item: &User{Id: id, Name: name, Friends: friends, CreatedOn: created_on}}, true
}

// Slice returns all rows found in the iterator as a Slice.
func (this *UServ_InsertUsersIter) Slice() []*UServ_InsertUsersRow {
	var results []*UServ_InsertUsersRow
	for {
		if i, ok := this.Next(); ok {
			results = append(results, i)
		} else {
			break
		}
	}
	return results
}

type UServ_ImplOpts struct {
	MAPPINGS UServ_TypeMappings
	HOOKS    UServ_Hooks
  HANDLERS RestOfUServHandlers
}

func DefaultUServImplOpts() UServ_ImplOpts {
	return UServ_ImplOpts{}
}

type UServ_Impl struct {
	opts    *UServ_ImplOpts
	QUERIES *UServ_Queries
	DB      *spanner.Client
}

func UServPersistImpl(db *spanner.Client, opts ...UServ_ImplOpts) *UServ_Impl {
	var myOpts UServ_ImplOpts
	if len(opts) > 0 {
		myOpts = opts[0]
	} else {
		myOpts = DefaultUServImplOpts()
	}
	return &UServ_Impl{
		opts:    &myOpts,
		QUERIES: UServPersistQueries(db, UServ_QueryOpts{MAPPINGS: myOpts.MAPPINGS}),
		DB:      db,
	}
}

// THIS is the grpc handler
func (this *UServ_Impl) InsertUsers(stream UServ_InsertUsersServer) error {
	// tx, err := DefaultClientStreamingPersistTx(stream.Context(), this.DB)
	// if err != nil {
	// 	return gstatus.Errorf(codes.Unknown, "error creating persist tx: %v", err)
	// }
	if err := this.InsertUsersTx(stream); err != nil {
		return gstatus.Errorf(codes.Unknown, "error executing 'insert_users' query: %v", err)
	}
	return nil
}

func (this *UServ_Impl) InsertUsersTx(stream UServ_InsertUsersServer) error {
	query := this.QUERIES.InsertUsersQuery(stream.Context())
	var first *User
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return gstatus.Errorf(codes.Unknown, "error receiving request: %v", err)
		}
		if first == nil {
			first = req
		}

		beforeRes, err := this.opts.HOOKS.InsertUsersBeforeHook(stream.Context(), req)
		if err != nil {
			return gstatus.Errorf(codes.Unknown, "error in before hook: %v", err)
		} else if beforeRes != nil {
			continue
		}

    // TODO had to do this to make sure context wasnt nil
    if query.ctx == nil {
      fmt.Println("context was nil")
      query.ctx = stream.Context()
    }

		result := query.Execute(req)
    // TODO this shouldnt be calling zero, it should have a row affected
		// if err := result.Zero(); err != nil {
    if err := result.One(); err != nil {
			return gstatus.Errorf(codes.InvalidArgument, "client streaming queries must return results")
		}
	}
  // TODO might need to handle commits and rollbacks uniquely with spanner.
	// if err := tx.Commit(); err != nil {
		// return fmt.Errorf("executed 'insert_users' query without error, but received error on commit: %v", err)
		// if rollbackErr := tx.Rollback(); rollbackErr != nil {
			// return fmt.Errorf("error executing 'insert_users' query :::AND COULD NOT ROLLBACK::: rollback err: %v, query err: %v", rollbackErr, err)
		// }
	// }
	res := &Empty{}
	if err := this.opts.HOOKS.InsertUsersAfterHook(stream.Context(), first, res); err != nil {
		return gstatus.Errorf(codes.Unknown, "error in after hook: %v", err)
	}

	if err := stream.SendAndClose(res); err != nil {
		return gstatus.Errorf(codes.Unknown, "error sending back response: %v", err)
	}

	return nil
}

// GetAllUsersQuery returns a new struct wrapping the current UServ_QueryOpts
// that will perform 'UServ' services 'get_all_users' on the database
// when executed
func (this *UServ_Queries) GetAllUsersQuery(ctx context.Context) *UServ_GetAllUsersQuery {
	return &UServ_GetAllUsersQuery{
		opts: UServ_QueryOpts{
			MAPPINGS: this.opts.MAPPINGS,
			db:       this.opts.db,
			ctx:      ctx,
		},
	}
}

type UServ_GetAllUsersQuery struct {
	opts UServ_QueryOpts
}

func (this *UServ_GetAllUsersQuery) QueryInTypeUser()  {}
func (this *UServ_GetAllUsersQuery) QueryOutTypeUser() {}

// Executes the query with parameters retrieved from x
func (this *UServ_GetAllUsersQuery) Execute(x UServ_GetAllUsersIn) *UServ_GetAllUsersIter {
	var setupErr error
	params := []interface{}{}
	result := &UServ_GetAllUsersIter{
		tm:  this.opts.MAPPINGS,
		ctx: this.opts.ctx,
	}
	if setupErr != nil {
		result.err = setupErr
		return result
	}
	// result.rows, result.err = this.opts.db.QueryContext(this.opts.ctx, "SELECT id, name, friends, created_on FROM users", params...)

  _, err := this.opts.db.ReadWriteTransaction(this.opts.ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
    stmt := spanner.Statement{
      SQL: fmt.Sprintf("SELECT id, name, friends, created_on FROM users", params...)}
    iter := txn.QueryWithStats(ctx, stmt)
    // err, rowCount := txn.Update(ctx, stmt)
    result.rows = iter
    result.result = SpannerResult{
      iter: iter,
    }
    // TODO consider the effects of calling "defer iter.Stop()" here
    return nil
  })
  result.err = err
  return result
}

type UServ_GetAllUsersIter struct {
  result SpannerResult
  rows   *spanner.RowIterator
  err    error
  tm     UServ_TypeMappings
  ctx    context.Context
}

func (this *UServ_GetAllUsersIter) IterOutTypeUser() {}
func (this *UServ_GetAllUsersIter) IterInTypeEmpty() {}

// Each performs 'fun' on each row in the result set.
// Each respects the context passed to it.
// It will stop iteration, and returns this.ctx.Err() if encountered.
func (this *UServ_GetAllUsersIter) Each(fun func(*UServ_GetAllUsersRow) error) error {
  for {
    select {
    case <-this.ctx.Done():
      return this.ctx.Err()
    default:
      if row, ok := this.Next(); !ok {
        return nil
      } else if err := fun(row); err != nil {
        return err
      }
    }
  }
}

// One returns the sole row, or ensures an error if there was not one result when this row is converted
func (this *UServ_GetAllUsersIter) One() *UServ_GetAllUsersRow {
  first, hasFirst := this.Next()
  _, hasSecond := this.Next()
  if !hasFirst || hasSecond {
    amount := "none"
    if hasSecond {
      amount = "multiple"
    }
    return &UServ_GetAllUsersRow{err: fmt.Errorf("expected exactly 1 result from query 'GetAllUsers' found %s", amount)}
  }
  return first
}

// Zero returns an error if there were any rows in the result
func (this *UServ_GetAllUsersIter) Zero() error {
  if _, ok := this.Next(); ok {
    return fmt.Errorf("expected exactly 0 results from query 'GetAllUsers'")
  }
  return nil
}

// Next returns the next scanned row out of the database, or (nil, false) if there are no more rows
func (this *UServ_GetAllUsersIter) Next() (*UServ_GetAllUsersRow, bool) {
  if this.rows == nil || this.err == io.EOF {
    return nil, false
  } else if this.err != nil {
    err := this.err
    this.err = io.EOF
    return &UServ_GetAllUsersRow{err: err}, true
  }
  row, err := this.rows.Next()
  if err == iterator.Done {
    if this.err == nil {
      this.err = io.EOF
      return nil, false
    }
  }
  if err != nil {
    return &UServ_GetAllUsersRow{err: err}, true
  }

  var id int64
  if err := row.ColumnByName("id",&id); err != nil {
      return &UServ_GetAllUsersRow{err: fmt.Errorf("cant convert db column id to protobuf go type ")}, true
  }

  var name string
  if err := row.ColumnByName("name",&name); err != nil {
      return &UServ_GetAllUsersRow{err: fmt.Errorf("cant convert db column name to protobuf go type ")}, true
  }

  var tmpFriends []byte
  if err := row.ColumnByName("friends", &tmpFriends); err != nil {
      return &UServ_GetAllUsersRow{err: fmt.Errorf("cant convert db column friends to protobuf go type *Friends")}, true
  }
  var friends = new(Friends)
  if err := proto.Unmarshal(tmpFriends, friends); err != nil {
    return &UServ_GetAllUsersRow{err: err}, true
  }

  var tmpCreatedOn spanner.GenericColumnValue
  var created_on *timestamp.Timestamp
  if err := row.ColumnByName("created_on", &tmpCreatedOn); err != nil {
    return &UServ_GetAllUsersRow{err: fmt.Errorf("could not convert mapped db column created_on to type *timestamp.Timestamp: %v", err)}, true
  }

  var converted = this.tm.TimestampTimestamp().Empty()
  if err := converted.SpannerScan(&tmpCreatedOn); err != nil {
      return &UServ_GetAllUsersRow{err: fmt.Errorf("could not convert mapped db column created_on to type on User.CreatedOn: %v", err)}, true
  }

  if err := converted.ToProto(&created_on); err != nil {
      return &UServ_GetAllUsersRow{err: fmt.Errorf("could not convert mapped db column created_onto type on User.CreatedOn: %v", err)}, true
  }

  return &UServ_GetAllUsersRow{item: &User{Id: id, Name: name, Friends: friends, CreatedOn: created_on}}, true
}

// Slice returns all rows found in the iterator as a Slice.
func (this *UServ_GetAllUsersIter) Slice() []*UServ_GetAllUsersRow {
  var results []*UServ_GetAllUsersRow
  for {
    if i, ok := this.Next(); ok {
      results = append(results, i)
    } else {
      break
    }
  }
  return results
}

type UServ_GetAllUsersIn interface {
}
type UServ_GetAllUsersOut interface {
  GetId() int64
  GetName() string
  GetFriends() *Friends
  GetCreatedOn() *timestamp.Timestamp
}
type UServ_GetAllUsersRow struct {
  item UServ_GetAllUsersOut
  err  error
}

func newUServ_GetAllUsersRow(item UServ_GetAllUsersOut, err error) *UServ_GetAllUsersRow {
  return &UServ_GetAllUsersRow{item, err}
}

// Unwrap takes an address to a proto.Message as its only parameter
// Unwrap can only set into output protos of that match method return types + the out option on the query itself
func (this *UServ_GetAllUsersRow) Unwrap(pointerToMsg proto.Message) error {
  if this.err != nil {
    return this.err
  }
  if o, ok := (pointerToMsg).(*User); ok {
    if o == nil {
      return fmt.Errorf("must initialize *User before giving to Unwrap()")
    }
    res, _ := this.User()
    _ = res
    o.Id = res.Id
    o.Name = res.Name
    o.Friends = res.Friends
    o.CreatedOn = res.CreatedOn
    return nil
  }
  return nil
}
func (this *UServ_GetAllUsersRow) User() (*User, error) {
  if this.err != nil {
    return nil, this.err
  }
  return &User{
    Id:        this.item.GetId(),
    Name:      this.item.GetName(),
    Friends:   this.item.GetFriends(),
    CreatedOn: this.item.GetCreatedOn(),
  }, nil
}

func (this *UServ_GetAllUsersRow) Proto() (*User, error) {
  if this.err != nil {
    return nil, this.err
  }
  return &User{
    Id:        this.item.GetId(),
    Name:      this.item.GetName(),
    Friends:   this.item.GetFriends(),
    CreatedOn: this.item.GetCreatedOn(),
  }, nil
}

// THIS is the grpc handler
func (this *UServ_Impl) GetAllUsers(req *Empty, stream UServ_GetAllUsersServer) error {
  // tx, err := DefaultServerStreamingPersistTx(stream.Context(), this.DB)
  // if err != nil {
  //   return gstatus.Errorf(codes.Unknown, "error creating persist tx: %v", err)
  // }
  if err := this.GetAllUsersTx(req, stream); err != nil {
    return gstatus.Errorf(codes.Unknown, "error executing 'get_all_users' query: %v", err)
  }
  return nil
}
func (this *UServ_Impl) GetAllUsersTx(req *Empty, stream UServ_GetAllUsersServer) error {
  query := this.QUERIES.GetAllUsersQuery(stream.Context())
  iter := query.Execute(req)
  return iter.Each(func(row *UServ_GetAllUsersRow) error {
    res, err := row.User()
    if err != nil {
      return err
    }
    return stream.Send(res)
  })
}

