package frog_blossom_db

type InitSetupConfigTxParams struct {
	UserId       int64               `json:"user_id"`
	Username     string              `json:"username"`
	Email        string              `json:"email"`
	UserURl      string              `json:"user_url"`
	InitialPages []CreatePagesParams `json:"initial_pages"`
	InitialPosts []CreatePostsParams `json:"initial_posts"`
	InitialMeta  []CreateMetaParams  `json:"initial_meta"`
}

type InitSetupConfigTxResult struct {
	User  User   `json:"user"`
	Pages []Page `json:"init_page"`
	Posts []Post `json:"post_id"`
	Metas []Meta `json:"meta_id"`
}

type CreateContentTxParams struct {
	UserId   int64               `json:"user_id"`
	Username string              `json:"username"`
	PageId   *int64              `json:"page_id"`
	PostId   *int64              `json:"post_id"`
	Pages    []CreatePagesParams `json:"pages"`
	Posts    []CreatePostsParams `json:"posts"`
	Metas    []CreateMetaParams  `json:"meta"`
}

type CreateContentTxResult struct {
	User   User   `json:"user"`
	PageId *Page  `json:"pages_id"`
	PostId *Post  `json:"post_id"`
	Posts  []Post `json:"post"`
	Metas  []Meta `json:"meta"`
	Pages  []Page `json:"page"`
}

type CreatePageTxParams struct {
	UserId   int64              `json:"user_id"`
	Username string             `json:"username"`
	PageId   *int64             `json:"page_id"`
	Pages    *CreatePagesParams `json:"pages"`
	Metas    CreateMetaParams   `json:"meta"`
}

type CreatePageTxResult struct {
	Pages Page `json:"page"`
	Metas Meta `json:"meta"`
}

type CreatePostTxParams struct {
	UserId   int64              `json:"user_id"`
	Username string             `json:"username"`
	PostId   *int64             `json:"page_id"`
	Posts    *CreatePostsParams `json:"pages"`
	Metas    CreateMetaParams   `json:"meta"`
}

type CreatePostTxResult struct {
	Posts Post `json:"page"`
	Metas Meta `json:"meta"`
}

type UpdateContentTxParams struct {
	UserId   int64              `json:"user_id"`
	Username string             `json:"username"`
	PageId   *int64             `json:"page_id"`
	PostId   *int64             `json:"post_id"`
	Pages    *UpdatePagesParams `json:"pages"`
	Posts    *UpdatePostsParams `json:"posts"`
	Metas    UpdateMetaParams   `json:"meta"`
}

type UpdateContentTxResult struct {
	Pages Page `json:"page"`
	Posts Post `json:"post"`
	Metas Meta `json:"meta"`
}

type UpdatePageTxParams struct {
	UserId   int64              `json:"user_id"`
	Username string             `json:"username"`
	PageId   *int64             `json:"page_id"`
	Pages    *UpdatePagesParams `json:"pages"`
	Metas    UpdateMetaParams   `json:"meta"`
}

type UpdatePageTxResult struct {
	Pages Page `json:"page"`
	Metas Meta `json:"meta"`
}

type UpdatePostTxParams struct {
	UserId   int64              `json:"user_id"`
	Username string             `json:"username"`
	Post     *int64             `json:"page_id"`
	Posts    *UpdatePostsParams `json:"pages"`
	Metas    UpdateMetaParams   `json:"meta"`
}

type UpdatePostTxResult struct {
	Pages Post `json:"page"`
	Metas Meta `json:"meta"`
}

type DeleteContentTxParams struct {
	PageId *int64 `json:"page_id"`
	PostId *int64 `json:"post_id"`
}

type DeletePostTxParams struct {
	PostId *int64 `json:"post_id"`
}

type DeletePageTxParams struct {
	PageId *int64 `json:"page_id"`
}

type DeleteContentTxResult struct {
	DeletedPost bool `json:"deleted_post"`
	DeletedPage bool `json:"deleted_page"`
	DeletedMeta bool `json:"deleted_meta"`
}

type DeletePostTxResult struct {
	DeletedPost bool `json:"deleted_post"`
	DeletedMeta bool `json:"deleted_meta"`
}

type DeletePageTxResult struct {
	DeletedPage bool `json:"deleted_page"`
	DeletedMeta bool `json:"deleted_meta"`
}
