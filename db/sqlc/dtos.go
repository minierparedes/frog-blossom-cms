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
	User       User   `json:"user"`
	UserStatus string `json:"user_status"`
	Pages      []Page `json:"init_page"`
	PageStatus string `json:"page_status"`
	Posts      []Post `json:"post_id"`
	PostStatus string `json:"post_status"`
	Metas      []Meta `json:"meta_id"`
	MetaStatus string `json:"meta_status"`
}
