// This file is generated by protoc-gen-persist
// Source File: pb/user.proto
// DO NOT EDIT !
package pb

import (
	sql "database/sql"
	fmt "fmt"
	io "io"

	proto "github.com/golang/protobuf/proto"
	persist_lib "github.com/tcncloud/protoc-gen-persist/test_service/user_sql/pb/persist_lib"
	context "golang.org/x/net/context"
	codes "google.golang.org/grpc/codes"
	gstatus "google.golang.org/grpc/status"
)

type UServImpl struct {
	PERSIST   *persist_lib.UServMethodReceiver
	FORWARDED RestOfUServHandlers
}
type RestOfUServHandlers interface {
	UpdateAllNames(req *Empty, stream UServ_UpdateAllNamesServer) error
}
type UServImplBuilder struct {
	err           error
	rest          RestOfUServHandlers
	queryHandlers *persist_lib.UServQueryHandlers
	i             *UServImpl
	db            sql.DB
}

func NewUServBuilder() *UServImplBuilder {
	return &UServImplBuilder{i: &UServImpl{}}
}
func (b *UServImplBuilder) WithRestOfGrpcHandlers(r RestOfUServHandlers) *UServImplBuilder {
	b.rest = r
	return b
}
func (b *UServImplBuilder) WithPersistQueryHandlers(p *persist_lib.UServQueryHandlers) *UServImplBuilder {
	b.queryHandlers = p
	return b
}
func (b *UServImplBuilder) WithDefaultQueryHandlers() *UServImplBuilder {
	accessor := persist_lib.NewSqlClientGetter(&b.db)
	queryHandlers := &persist_lib.UServQueryHandlers{
		CreateTableHandler:     persist_lib.DefaultCreateTableHandler(accessor),
		InsertUsersHandler:     persist_lib.DefaultInsertUsersHandler(accessor),
		GetAllUsersHandler:     persist_lib.DefaultGetAllUsersHandler(accessor),
		SelectUserByIdHandler:  persist_lib.DefaultSelectUserByIdHandler(accessor),
		UpdateUserNamesHandler: persist_lib.DefaultUpdateUserNamesHandler(accessor),
		UpdateNameToFooHandler: persist_lib.DefaultUpdateNameToFooHandler(accessor),
		GetFriendsHandler:      persist_lib.DefaultGetFriendsHandler(accessor),
		DropTableHandler:       persist_lib.DefaultDropTableHandler(accessor),
	}
	b.queryHandlers = queryHandlers
	return b
}

// set the custom handlers you want to use in the handlers
// this method will make sure to use a default handler if
// the handler is nil.
func (b *UServImplBuilder) WithNilAsDefaultQueryHandlers(p *persist_lib.UServQueryHandlers) *UServImplBuilder {
	accessor := persist_lib.NewSqlClientGetter(&b.db)
	if p.CreateTableHandler == nil {
		p.CreateTableHandler = persist_lib.DefaultCreateTableHandler(accessor)
	}
	if p.InsertUsersHandler == nil {
		p.InsertUsersHandler = persist_lib.DefaultInsertUsersHandler(accessor)
	}
	if p.GetAllUsersHandler == nil {
		p.GetAllUsersHandler = persist_lib.DefaultGetAllUsersHandler(accessor)
	}
	if p.SelectUserByIdHandler == nil {
		p.SelectUserByIdHandler = persist_lib.DefaultSelectUserByIdHandler(accessor)
	}
	if p.UpdateUserNamesHandler == nil {
		p.UpdateUserNamesHandler = persist_lib.DefaultUpdateUserNamesHandler(accessor)
	}
	if p.UpdateNameToFooHandler == nil {
		p.UpdateNameToFooHandler = persist_lib.DefaultUpdateNameToFooHandler(accessor)
	}
	if p.GetFriendsHandler == nil {
		p.GetFriendsHandler = persist_lib.DefaultGetFriendsHandler(accessor)
	}
	if p.DropTableHandler == nil {
		p.DropTableHandler = persist_lib.DefaultDropTableHandler(accessor)
	}
	b.queryHandlers = p
	return b
}
func (b *UServImplBuilder) WithSqlClient(c *sql.DB) *UServImplBuilder {
	b.db = *c
	return b
}
func (b *UServImplBuilder) WithNewSqlDb(driverName, dataSourceName string) *UServImplBuilder {
	db, err := sql.Open(driverName, dataSourceName)
	b.err = err
	b.db = *db
	return b
}
func (b *UServImplBuilder) Build() (*UServImpl, error) {
	if b.err != nil {
		return nil, b.err
	}
	b.i.PERSIST = &persist_lib.UServMethodReceiver{Handlers: *b.queryHandlers}
	b.i.FORWARDED = b.rest
	return b.i, nil
}
func (b *UServImplBuilder) MustBuild() *UServImpl {
	s, err := b.Build()
	if err != nil {
		panic("error in builder: " + err.Error())
	}
	return s
}
func EmptyToUServPersistType(req *Empty) (*persist_lib.EmptyForUServ, error) {
	params := &persist_lib.EmptyForUServ{}
	return params, nil
}
func EmptyFromUServDatabaseRow(row persist_lib.Scanable) (*Empty, error) {
	res := &Empty{}
	if err := row.Scan(); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return res, nil
}
func IterUServEmptyProto(iter *persist_lib.Result, next func(i *Empty) error) error {
	return iter.Do(func(r persist_lib.Scanable) error {
		item, err := EmptyFromUServDatabaseRow(r)
		if err != nil {
			return fmt.Errorf("error converting Empty row to protobuf message: %s", err)
		}
		return next(item)
	})
}
func UserToUServPersistType(req *User) (*persist_lib.UserForUServ, error) {
	params := &persist_lib.UserForUServ{}
	params.Id = req.Id
	params.Name = req.Name
	if req.Friends == nil {
		req.Friends = new(Friends)
	}
	{
		raw, err := proto.Marshal(req.Friends)
		if err != nil {
			return nil, err
		}
		params.Friends = raw
	}
	params.CreatedOn = (TimeString{}).ToSql(req.CreatedOn)
	return params, nil
}
func UserFromUServDatabaseRow(row persist_lib.Scanable) (*User, error) {
	res := &User{}
	var Id_ int64
	var Name_ string
	var Friends_ []byte
	var CreatedOn_ TimeString
	if err := row.Scan(
		&Id_,
		&Name_,
		&Friends_,
		&CreatedOn_,
	); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	res.Id = Id_
	res.Name = Name_
	{
		var converted = new(Friends)
		if err := proto.Unmarshal(Friends_, converted); err != nil {
			return nil, err
		}
		res.Friends = converted
	}
	res.CreatedOn = CreatedOn_.ToProto()
	return res, nil
}
func IterUServUserProto(iter *persist_lib.Result, next func(i *User) error) error {
	return iter.Do(func(r persist_lib.Scanable) error {
		item, err := UserFromUServDatabaseRow(r)
		if err != nil {
			return fmt.Errorf("error converting User row to protobuf message: %s", err)
		}
		return next(item)
	})
}
func FriendsQueryToUServPersistType(req *FriendsQuery) (*persist_lib.FriendsQueryForUServ, error) {
	params := &persist_lib.FriendsQueryForUServ{}
	params.Names = (SliceStringConverter{}).ToSql(req.Names)
	return params, nil
}
func (s *UServImpl) CreateTable(ctx context.Context, req *Empty) (*Empty, error) {
	var err error
	var res = &Empty{}
	_ = err
	_ = res
	params, err := EmptyToUServPersistType(req)
	if err != nil {
		return nil, err
	}
	var iterErr error
	err = s.PERSIST.CreateTable(ctx, params, func(row persist_lib.Scanable) {
		if row == nil { // there was no return data
			return
		}
		res, err = EmptyFromUServDatabaseRow(row)
		if err != nil {
			iterErr = err
			return
		}
	})
	if err != nil {
		return nil, gstatus.Errorf(codes.Unknown, "error calling persist service: %v", err)
	} else if iterErr != nil {
		return nil, iterErr
	}
	return res, nil
}
func (s *UServImpl) InsertUsers(stream UServ_InsertUsersServer) error {
	var err error
	_ = err
	res := &Empty{}
	feed, stop := s.PERSIST.InsertUsers(stream.Context())
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return gstatus.Errorf(codes.Unknown, "error receiving request: %v", err)
		}
		beforeRes, err := IncId(req)
		if err != nil {
			return gstatus.Errorf(codes.Unknown, "error in before hook: %v", err)
		} else if beforeRes != nil {
			continue
		}
		params, err := UserToUServPersistType(req)
		if err != nil {
			return err
		}
		feed(params)
	}
	row, err := stop()
	if err != nil {
		return gstatus.Errorf(codes.Unknown, "error receiving result row: %v", err)
	}
	if row != nil {
		res, err = EmptyFromUServDatabaseRow(row)
		if err != nil {
			return err
		}
	}
	if err := stream.SendAndClose(res); err != nil {
		return gstatus.Errorf(codes.Unknown, "error sending back response: %v", err)
	}
	return nil
}
func (s *UServImpl) GetAllUsers(req *Empty, stream UServ_GetAllUsersServer) error {
	var err error
	_ = err
	params, err := EmptyToUServPersistType(req)
	if err != nil {
		return err
	}
	var iterErr error
	err = s.PERSIST.GetAllUsers(stream.Context(), params, func(row persist_lib.Scanable) {
		if row == nil { // there was no return data
			return
		}
		res, err := UserFromUServDatabaseRow(row)
		if err != nil {
			iterErr = err
			return
		}
		if err := stream.Send(res); err != nil {
			iterErr = gstatus.Errorf(codes.Unknown, "error during iteration: %v", err)
		}
	})
	if err != nil {
		return gstatus.Errorf(codes.Unknown, "error during iteration: %v", err)
	} else if iterErr != nil {
		return iterErr
	}
	return nil
}
func (s *UServImpl) SelectUserById(ctx context.Context, req *User) (*User, error) {
	var err error
	var res = &User{}
	_ = err
	_ = res
	params, err := UserToUServPersistType(req)
	if err != nil {
		return nil, err
	}
	var iterErr error
	err = s.PERSIST.SelectUserById(ctx, params, func(row persist_lib.Scanable) {
		if row == nil { // there was no return data
			return
		}
		res, err = UserFromUServDatabaseRow(row)
		if err != nil {
			iterErr = err
			return
		}
	})
	if err != nil {
		return nil, gstatus.Errorf(codes.Unknown, "error calling persist service: %v", err)
	} else if iterErr != nil {
		return nil, iterErr
	}
	return res, nil
}
func (s *UServImpl) UpdateUserNames(stream UServ_UpdateUserNamesServer) error {
	var err error
	_ = err
	feed, stop := s.PERSIST.UpdateUserNames(stream.Context())
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return gstatus.Errorf(codes.Unknown, "error receiving request: %v", err)
		}
		params, err := UserToUServPersistType(req)
		if err != nil {
			return err
		}
		row, err := feed(params)
		if err != nil {
			return gstatus.Errorf(codes.Unknown, "error receiving result row: %v", err)
		}
		if row != nil {
			res, err := UserFromUServDatabaseRow(row)
			if err != nil {
				return err
			}
			if err := stream.Send(res); err != nil {
				return err
			}
		}
	}
	return stop()
}
func (s *UServImpl) UpdateNameToFoo(ctx context.Context, req *User) (*Empty, error) {
	var err error
	var res = &Empty{}
	_ = err
	_ = res
	params, err := UserToUServPersistType(req)
	if err != nil {
		return nil, err
	}
	var iterErr error
	err = s.PERSIST.UpdateNameToFoo(ctx, params, func(row persist_lib.Scanable) {
		if row == nil { // there was no return data
			return
		}
		res, err = EmptyFromUServDatabaseRow(row)
		if err != nil {
			iterErr = err
			return
		}
	})
	if err != nil {
		return nil, gstatus.Errorf(codes.Unknown, "error calling persist service: %v", err)
	} else if iterErr != nil {
		return nil, iterErr
	}
	return res, nil
}
func (s *UServImpl) UpdateAllNames(req *Empty, stream UServ_UpdateAllNamesServer) error {
	return s.FORWARDED.UpdateAllNames(req, stream)
}
func (s *UServImpl) GetFriends(req *FriendsQuery, stream UServ_GetFriendsServer) error {
	var err error
	_ = err
	params, err := FriendsQueryToUServPersistType(req)
	if err != nil {
		return err
	}
	var iterErr error
	err = s.PERSIST.GetFriends(stream.Context(), params, func(row persist_lib.Scanable) {
		if row == nil { // there was no return data
			return
		}
		res, err := UserFromUServDatabaseRow(row)
		if err != nil {
			iterErr = err
			return
		}
		if err := stream.Send(res); err != nil {
			iterErr = gstatus.Errorf(codes.Unknown, "error during iteration: %v", err)
		}
	})
	if err != nil {
		return gstatus.Errorf(codes.Unknown, "error during iteration: %v", err)
	} else if iterErr != nil {
		return iterErr
	}
	return nil
}
func (s *UServImpl) DropTable(ctx context.Context, req *Empty) (*Empty, error) {
	var err error
	var res = &Empty{}
	_ = err
	_ = res
	params, err := EmptyToUServPersistType(req)
	if err != nil {
		return nil, err
	}
	var iterErr error
	err = s.PERSIST.DropTable(ctx, params, func(row persist_lib.Scanable) {
		if row == nil { // there was no return data
			return
		}
		res, err = EmptyFromUServDatabaseRow(row)
		if err != nil {
			iterErr = err
			return
		}
	})
	if err != nil {
		return nil, gstatus.Errorf(codes.Unknown, "error calling persist service: %v", err)
	} else if iterErr != nil {
		return nil, iterErr
	}
	return res, nil
}
