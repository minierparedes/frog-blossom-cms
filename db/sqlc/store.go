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
func (store *Store) CreatePostsTx(ctx context.Context, args CreatePostsTxParams) (CreatePostsTxResul, error) {
	var result CreatePostsTxResul

	err := store.executeTx(ctx, func(q *Queries) error {
		var err error

		user, err := q.GetUsers(ctx, args.UserId)
		if err != nil {
			return fmt.Errorf("get users err: %v", err)
		}
		result.User = user

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
