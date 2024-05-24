package frog_blossom_db

import (
	"context"
	"database/sql"
	"fmt"
)

// Functions for executing db queries and transactions

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// executes a function within a db transaction

func (store *Store) executeTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("transaction err: %v, rollback err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

// InitSetupConfigTx populates db tables with initial site-specific config data
// Use user info to populate the `posts`, `pages`, and `meta` tables
func (store *Store) InitSetupConfigTx(ctx context.Context, args InitSetupConfigTxParams) (InitSetupConfigTxResult, error) {
	var result InitSetupConfigTxResult

	err := store.executeTx(ctx, func(q *Queries) error {
		var err error

		user, err := q.GetUsers(ctx, args.UserId)
		if err != nil {
			return fmt.Errorf("get users err: %v", err)
		}
		result.User = user

		for _, postsParams := range args.InitialPosts {
			post, err := q.CreatePosts(ctx, postsParams)
			if err != nil {
				return fmt.Errorf("create posts err: %v", err)
			}
			result.Posts = append(result.Posts, post)
		}

		for _, pageParams := range args.InitialPages {
			page, err := q.CreatePages(ctx, pageParams)
			if err != nil {
				return fmt.Errorf("create pages err: %v", err)
			}
			result.Pages = append(result.Pages, page)
		}

		for _, metaParas := range args.InitialMeta {
			meta, err := q.CreateMeta(ctx, metaParas)
			if err != nil {
				return fmt.Errorf("create meta err: %v", err)
			}
			result.Metas = append(result.Metas, meta)
		}

		return nil
	})
	return result, err
}

// CreatePostsTx creates new posts content based on user information
// It utilizes user info(users.id, users.username) to create the `posts` and its respective `meta`.
func (store *Store) CreatePostsTx(ctx context.Context, args CreateContentTxParams) (CreateContentTxResult, error) {
	var result CreateContentTxResult

	err := store.executeTx(ctx, func(q *Queries) error {
		var err error

		user, err := q.GetUsers(ctx, args.UserId)
		if err != nil {
			return fmt.Errorf("get users err: %v", err)
		}
		result.User = user

		if args.PageId != nil {
			page, err := q.GetPages(ctx, *args.PageId)
			if err != nil {
				return fmt.Errorf("get pages err: %v", err)
			}
			result.PageId = &page
		}

		for _, postParams := range args.Posts {
			post, err := q.CreatePosts(ctx, postParams)
			if err != nil {
				return fmt.Errorf("create posts err: %v", err)
			}
			result.Posts = append(result.Posts, post)
		}

		for _, metaParas := range args.Metas {
			meta, err := q.CreateMeta(ctx, metaParas)
			if err != nil {
				return fmt.Errorf("create meta err: %v", err)
			}
			result.Metas = append(result.Metas, meta)
		}

		return nil
	})
	return result, err
}

// CreatePageTx creates new pages content based on user information
// It utilizes user info(users.id, users.username) to create the `page` and its respective `meta`.
func (store *Store) CreatePageTx(ctx context.Context, args CreateContentTxParams) (CreateContentTxResult, error) {
	var result CreateContentTxResult

	err := store.executeTx(ctx, func(q *Queries) error {
		var err error

		user, err := q.GetUsers(ctx, args.UserId)
		if err != nil {
			return fmt.Errorf("get users err: %v", err)
		}
		result.User = user

		posts, err := q.GetPosts(ctx, *args.PostId)
		if err != nil {
			return fmt.Errorf("get posts err: %v", err)
		}
		result.PostId = &posts

		for _, pageParams := range args.Pages {
			page, err := q.CreatePages(ctx, pageParams)
			if err != nil {
				return fmt.Errorf("create pages err: %v", err)
			}
			result.Pages = append(result.Pages, page)
		}

		for _, metaParas := range args.Metas {
			meta, err := q.CreateMeta(ctx, metaParas)
			if err != nil {
				return fmt.Errorf("create meta err: %v", err)
			}
			result.Metas = append(result.Metas, meta)
		}

		return nil
	})

	return result, err
}

// UpdatePostsTx updates existing content in the `posts` table and its respective `meta` table.
// It utilizes user info (users.id, users.username) to update the content and its associated metadata.
//func (store *Store) UpdatePostsTx(ctx context.Context, args UpdateContentTxParams) (UpdateContentTxResult, error) {
//	var result UpdateContentTxResult
//
//	err := store.executeTx(ctx, func(q *Queries) error {
//		var err error
//
//		user, err := q.GetUsers(ctx, args.UserId)
//		if err != nil {
//			return fmt.Errorf("get user err: %v", err)
//		}
//		result.User = user
//
//		post, err := q.GetPosts(ctx, args.PostId)
//		if err != nil {
//			return fmt.Errorf("get post err: %v", err)
//		}
//		result.PostId = post
//
//		meta, err := q.GetMeta(ctx, args.)
//
//		for _, postParams := range args.Posts {
//			post, err := q.UpdatePosts(ctx, postParams)
//			if err != nil {
//				return fmt.Errorf("update post err: %v", err)
//			}
//			result.Posts = append(result.Posts, post)
//		}
//	})
//
//}
