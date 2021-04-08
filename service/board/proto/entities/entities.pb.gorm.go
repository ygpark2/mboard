// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: service/board/proto/entities/entities.proto

/*
Package entities is a generated protocol buffer package.

It is generated from these files:
	service/board/proto/entities/entities.proto

It has these top-level messages:
	Board
*/
package entities

import context "context"
import fmt "fmt"
import time "time"

import errors1 "github.com/infobloxopen/protoc-gen-gorm/errors"
import field_mask1 "google.golang.org/genproto/protobuf/field_mask"
import go_uuid1 "github.com/satori/go.uuid"
import gorm1 "github.com/jinzhu/gorm"
import gorm2 "github.com/infobloxopen/atlas-app-toolkit/gorm"
import pq1 "github.com/lib/pq"
import ptypes1 "github.com/golang/protobuf/ptypes"
import types1 "github.com/infobloxopen/protoc-gen-gorm/types"

import math "math"
import _ "github.com/golang/protobuf/ptypes/timestamp"
import google_protobuf1 "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/envoyproxy/protoc-gen-validate/validate"

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = math.Inf

type BoardORM struct {
	CreatedAt   *time.Time      `gorm:"not null"`
	DeletedAt   *time.Time      `gorm:"index:idx_boards_deleted_at"`
	Description string          `gorm:"not null"`
	Id          go_uuid1.UUID   `gorm:"type:uuid;primary_key;unique;not null"`
	MobileTitle *string         `gorm:"size:255;not null"`
	Notices     pq1.StringArray `gorm:"type:text[]"`
	Order       *uint32         `gorm:"not null"`
	Search      *bool           `gorm:"not null"`
	Title       *string         `gorm:"size:255;not null"`
	UpdatedAt   *time.Time      `gorm:"not null"`
}

// TableName overrides the default tablename generated by GORM
func (BoardORM) TableName() string {
	return "boards"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *Board) ToORM(ctx context.Context) (BoardORM, error) {
	to := BoardORM{}
	var err error
	if prehook, ok := interface{}(m).(BoardWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	if m.Id != nil {
		to.Id, err = go_uuid1.FromString(m.Id.Value)
		if err != nil {
			return to, err
		}
	} else {
		to.Id = go_uuid1.Nil
	}
	if m.CreatedAt != nil {
		var t time.Time
		if t, err = ptypes1.Timestamp(m.CreatedAt); err != nil {
			return to, err
		}
		to.CreatedAt = &t
	}
	if m.UpdatedAt != nil {
		var t time.Time
		if t, err = ptypes1.Timestamp(m.UpdatedAt); err != nil {
			return to, err
		}
		to.UpdatedAt = &t
	}
	if m.DeletedAt != nil {
		var t time.Time
		if t, err = ptypes1.Timestamp(m.DeletedAt); err != nil {
			return to, err
		}
		to.DeletedAt = &t
	}
	if m.Title != nil {
		v := m.Title.Value
		to.Title = &v
	}
	if m.MobileTitle != nil {
		v := m.MobileTitle.Value
		to.MobileTitle = &v
	}
	if m.Order != nil {
		v := m.Order.Value
		to.Order = &v
	}
	if m.Search != nil {
		v := m.Search.Value
		to.Search = &v
	}
	to.Description = m.Description
	if m.Notices != nil {
		to.Notices = make(pq1.StringArray, len(m.Notices))
		copy(to.Notices, m.Notices)
	}
	if posthook, ok := interface{}(m).(BoardWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *BoardORM) ToPB(ctx context.Context) (Board, error) {
	to := Board{}
	var err error
	if prehook, ok := interface{}(m).(BoardWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = &types1.UUID{Value: m.Id.String()}
	if m.CreatedAt != nil {
		if to.CreatedAt, err = ptypes1.TimestampProto(*m.CreatedAt); err != nil {
			return to, err
		}
	}
	if m.UpdatedAt != nil {
		if to.UpdatedAt, err = ptypes1.TimestampProto(*m.UpdatedAt); err != nil {
			return to, err
		}
	}
	if m.DeletedAt != nil {
		if to.DeletedAt, err = ptypes1.TimestampProto(*m.DeletedAt); err != nil {
			return to, err
		}
	}
	if m.Title != nil {
		to.Title = &google_protobuf1.StringValue{Value: *m.Title}
	}
	if m.MobileTitle != nil {
		to.MobileTitle = &google_protobuf1.StringValue{Value: *m.MobileTitle}
	}
	if m.Order != nil {
		to.Order = &google_protobuf1.UInt32Value{Value: *m.Order}
	}
	if m.Search != nil {
		to.Search = &google_protobuf1.BoolValue{Value: *m.Search}
	}
	to.Description = m.Description
	if m.Notices != nil {
		to.Notices = make(pq1.StringArray, len(m.Notices))
		copy(to.Notices, m.Notices)
	}
	if posthook, ok := interface{}(m).(BoardWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type Board the arg will be the target, the caller the one being converted from

// BoardBeforeToORM called before default ToORM code
type BoardWithBeforeToORM interface {
	BeforeToORM(context.Context, *BoardORM) error
}

// BoardAfterToORM called after default ToORM code
type BoardWithAfterToORM interface {
	AfterToORM(context.Context, *BoardORM) error
}

// BoardBeforeToPB called before default ToPB code
type BoardWithBeforeToPB interface {
	BeforeToPB(context.Context, *Board) error
}

// BoardAfterToPB called after default ToPB code
type BoardWithAfterToPB interface {
	AfterToPB(context.Context, *Board) error
}

// DefaultCreateBoard executes a basic gorm create call
func DefaultCreateBoard(ctx context.Context, in *Board, db *gorm1.DB) (*Board, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(BoardORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(BoardORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type BoardORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type BoardORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm1.DB) error
}

// DefaultReadBoard executes a basic gorm read call
func DefaultReadBoard(ctx context.Context, in *Board, db *gorm1.DB) (*Board, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if ormObj.Id == go_uuid1.Nil {
		return nil, errors1.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(BoardORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if db, err = gorm2.ApplyFieldSelection(ctx, db, nil, &BoardORM{}); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(BoardORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := BoardORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(BoardORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type BoardORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type BoardORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type BoardORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm1.DB) error
}

func DefaultDeleteBoard(ctx context.Context, in *Board, db *gorm1.DB) error {
	if in == nil {
		return errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.Id == go_uuid1.Nil {
		return errors1.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(BoardORMWithBeforeDelete_); ok {
		if db, err = hook.BeforeDelete_(ctx, db); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&BoardORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(BoardORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type BoardORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type BoardORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm1.DB) error
}

func DefaultDeleteBoardSet(ctx context.Context, in []*Board, db *gorm1.DB) error {
	if in == nil {
		return errors1.NilArgumentError
	}
	var err error
	keys := []go_uuid1.UUID{}
	for _, obj := range in {
		ormObj, err := obj.ToORM(ctx)
		if err != nil {
			return err
		}
		if ormObj.Id == go_uuid1.Nil {
			return errors1.EmptyIdError
		}
		keys = append(keys, ormObj.Id)
	}
	if hook, ok := (interface{}(&BoardORM{})).(BoardORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	err = db.Where("id in (?)", keys).Delete(&BoardORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&BoardORM{})).(BoardORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type BoardORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*Board, *gorm1.DB) (*gorm1.DB, error)
}
type BoardORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*Board, *gorm1.DB) error
}

// DefaultStrictUpdateBoard clears / replaces / appends first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateBoard(ctx context.Context, in *Board, db *gorm1.DB) (*Board, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateBoard")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &BoardORM{}
	db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(BoardORMWithBeforeStrictUpdateCleanup); ok {
		if db, err = hook.BeforeStrictUpdateCleanup(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(BoardORMWithBeforeStrictUpdateSave); ok {
		if db, err = hook.BeforeStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(BoardORMWithAfterStrictUpdateSave); ok {
		if err = hook.AfterStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	return &pbResponse, err
}

type BoardORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type BoardORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type BoardORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm1.DB) error
}

// DefaultPatchBoard executes a basic gorm update call with patch behavior
func DefaultPatchBoard(ctx context.Context, in *Board, updateMask *field_mask1.FieldMask, db *gorm1.DB) (*Board, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	var pbObj Board
	var err error
	if hook, ok := interface{}(&pbObj).(BoardWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := DefaultReadBoard(ctx, &Board{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(BoardWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskBoard(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(BoardWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateBoard(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(BoardWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type BoardWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *Board, *field_mask1.FieldMask, *gorm1.DB) (*gorm1.DB, error)
}
type BoardWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *Board, *field_mask1.FieldMask, *gorm1.DB) (*gorm1.DB, error)
}
type BoardWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *Board, *field_mask1.FieldMask, *gorm1.DB) (*gorm1.DB, error)
}
type BoardWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *Board, *field_mask1.FieldMask, *gorm1.DB) error
}

// DefaultPatchSetBoard executes a bulk gorm update call with patch behavior
func DefaultPatchSetBoard(ctx context.Context, objects []*Board, updateMasks []*field_mask1.FieldMask, db *gorm1.DB) ([]*Board, error) {
	if len(objects) != len(updateMasks) {
		return nil, fmt.Errorf(errors1.BadRepeatedFieldMaskTpl, len(updateMasks), len(objects))
	}

	results := make([]*Board, 0, len(objects))
	for i, patcher := range objects {
		pbResponse, err := DefaultPatchBoard(ctx, patcher, updateMasks[i], db)
		if err != nil {
			return nil, err
		}

		results = append(results, pbResponse)
	}

	return results, nil
}

// DefaultApplyFieldMaskBoard patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskBoard(ctx context.Context, patchee *Board, patcher *Board, updateMask *field_mask1.FieldMask, prefix string, db *gorm1.DB) (*Board, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors1.NilArgumentError
	}
	var err error
	for _, f := range updateMask.Paths {
		if f == prefix+"Id" {
			patchee.Id = patcher.Id
			continue
		}
		if f == prefix+"CreatedAt" {
			patchee.CreatedAt = patcher.CreatedAt
			continue
		}
		if f == prefix+"UpdatedAt" {
			patchee.UpdatedAt = patcher.UpdatedAt
			continue
		}
		if f == prefix+"DeletedAt" {
			patchee.DeletedAt = patcher.DeletedAt
			continue
		}
		if f == prefix+"Title" {
			patchee.Title = patcher.Title
			continue
		}
		if f == prefix+"MobileTitle" {
			patchee.MobileTitle = patcher.MobileTitle
			continue
		}
		if f == prefix+"Order" {
			patchee.Order = patcher.Order
			continue
		}
		if f == prefix+"Search" {
			patchee.Search = patcher.Search
			continue
		}
		if f == prefix+"Description" {
			patchee.Description = patcher.Description
			continue
		}
		if f == prefix+"Notices" {
			patchee.Notices = patcher.Notices
			continue
		}
		if f == prefix+"TotalPosts" {
			patchee.TotalPosts = patcher.TotalPosts
			continue
		}
		if f == prefix+"TotalComments" {
			patchee.TotalComments = patcher.TotalComments
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListBoard executes a gorm list call
func DefaultListBoard(ctx context.Context, db *gorm1.DB) ([]*Board, error) {
	in := Board{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(BoardORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm2.ApplyCollectionOperators(ctx, db, &BoardORM{}, &Board{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(BoardORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("id")
	ormResponse := []BoardORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(BoardORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*Board{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type BoardORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type BoardORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type BoardORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm1.DB, *[]BoardORM) error
}