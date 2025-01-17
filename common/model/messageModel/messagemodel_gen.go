// Code generated by goctl. DO NOT EDIT.

package messageModel

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	messageFieldNames          = builder.RawFieldNames(&Message{})
	messageRows                = strings.Join(messageFieldNames, ",")
	messageRowsExpectAutoSet   = strings.Join(stringx.Remove(messageFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	messageRowsWithPlaceHolder = strings.Join(stringx.Remove(messageFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheTiktokMessageIdPrefix = "cache:tiktok:message:id:"
)

type (
	messageModel interface {
		Insert(ctx context.Context, data *Message) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Message, error)
		Update(ctx context.Context, data *Message) error
		Delete(ctx context.Context, id int64) error
		FindMessageListByUserID(ctx context.Context, userid int64, to_userid int64) ([]*Message, error)
	}

	defaultMessageModel struct {
		sqlc.CachedConn
		table string
	}

	Message struct {
		Id         int64     `db:"id"`
		Content    string    `db:"content"`
		FromUserId int64     `db:"from_user_id"`
		ToUserId   int64     `db:"to_user_id"`
		CreateTime time.Time `db:"create_time"`
	}
)

func newMessageModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultMessageModel {
	return &defaultMessageModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`message`",
	}
}

func (m *defaultMessageModel) Delete(ctx context.Context, id int64) error {
	tiktokMessageIdKey := fmt.Sprintf("%s%v", cacheTiktokMessageIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, tiktokMessageIdKey)
	return err
}

func (m *defaultMessageModel) FindOne(ctx context.Context, id int64) (*Message, error) {
	tiktokMessageIdKey := fmt.Sprintf("%s%v", cacheTiktokMessageIdPrefix, id)
	var resp Message
	err := m.QueryRowCtx(ctx, &resp, tiktokMessageIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", messageRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultMessageModel) FindMessageListByUserID(ctx context.Context, from_user_id int64, to_userid int64) ([]*Message, error) {
	fmt.Printf("findall:::::::::::::::::::::::::")
	tiktokCommentIdKey := fmt.Sprintf("%s%v", cacheTiktokMessageIdPrefix, from_user_id+to_userid)
	var resp []*Message
	err := m.QueryRowCtx(ctx, &resp, tiktokCommentIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `from_user_id` = ? and `to_user_id` = ? order by `create_time` DESC", messageRows, m.table)
		fmt.Printf("sql:------------->%v", query)
		return conn.QueryRowsCtx(ctx, v, query, from_user_id+to_userid)
	})
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultMessageModel) Insert(ctx context.Context, data *Message) (sql.Result, error) {
	tiktokMessageIdKey := fmt.Sprintf("%s%v", cacheTiktokMessageIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, messageRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Content, data.FromUserId, data.ToUserId)
	}, tiktokMessageIdKey)
	return ret, err
}

func (m *defaultMessageModel) Update(ctx context.Context, data *Message) error {
	tiktokMessageIdKey := fmt.Sprintf("%s%v", cacheTiktokMessageIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, messageRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Content, data.FromUserId, data.ToUserId, data.Id)
	}, tiktokMessageIdKey)
	return err
}

func (m *defaultMessageModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheTiktokMessageIdPrefix, primary)
}

func (m *defaultMessageModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", messageRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultMessageModel) tableName() string {
	return m.table
}
