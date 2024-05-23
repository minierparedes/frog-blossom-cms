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

type CreatePostsTxParams struct {
	UserId   int64               `json:"user_id"`
	Username string              `json:"username"`
	Posts    []CreatePostsParams `json:"posts"`
	Metas    []CreateMetaParams  `json:"meta"`
}

type CreatePostsTxResul struct {
	User  User   `json:"user"`
	Posts []Post `json:"post"`
	Metas []Meta `json:"meta"`
}
