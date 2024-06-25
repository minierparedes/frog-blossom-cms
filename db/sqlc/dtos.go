package db

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

type CreatePageTxParams struct {
	UserId   int64              `json:"user_id"`
	Username string             `json:"username"`
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
	Posts    *CreatePostsParams `json:"posts"`
	Metas    CreateMetaParams   `json:"meta"`
}

type CreatePostTxResult struct {
	Posts Post `json:"posts"`
	Metas Meta `json:"meta"`
}

type CreateMetaTxParams struct {
	PostsID         *int64  `json:"posts_id"`
	MetaTitle       *string `json:"meta_title"`
	MetaDescription *string `json:"meta_description"`
	MetaRobots      *string `json:"meta_robots"`
	MetaOgImage     *string `json:"meta_og_image"`
	Locale          *string `json:"locale"`
	PageAmount      int64   `json:"page_amount"`
	SiteLanguage    *string `json:"site_language"`
	MetaKey         string  `json:"meta_key"`
	MetaValue       string  `json:"meta_value"`
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
	PostId   *int64             `json:"page_id"`
	Posts    *UpdatePostsParams `json:"posts"`
	Metas    UpdateMetaParams   `json:"meta"`
}

type UpdatePostTxResult struct {
	Posts Post `json:"posts"`
	Metas Meta `json:"meta"`
}

type UpdateMetaTxParams struct {
	ID              int64   `json:"id"`
	PageID          *int64  `json:"page_id"`
	PostsID         *int64  `json:"posts_id"`
	MetaTitle       *string `json:"meta_title"`
	MetaDescription *string `json:"meta_description"`
	MetaRobots      *string `json:"meta_robots"`
	MetaOgImage     *string `json:"meta_og_image"`
	Locale          *string `json:"locale"`
	PageAmount      int64   `json:"page_amount"`
	SiteLanguage    *string `json:"site_language"`
	MetaKey         string  `json:"meta_key"`
	MetaValue       string  `json:"meta_value"`
}

type DeletePostTxParams struct {
	PostId *int64 `json:"post_id"`
}

type DeletePageTxParams struct {
	PageId *int64 `json:"page_id"`
}

type DeletePostTxResult struct {
	DeletedPost bool `json:"deleted_post"`
	DeletedMeta bool `json:"deleted_meta"`
}

type DeletePageTxResult struct {
	DeletedPage bool `json:"deleted_page"`
	DeletedMeta bool `json:"deleted_meta"`
}
