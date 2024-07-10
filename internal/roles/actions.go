package roles

type Action string

const (
	CreateUser Action = "CREATE_USER"
	GetUser    Action = "GET_USER"
	ListUsers  Action = "LIST_USERS"
	CreatePost Action = "CREATE_POST"
	CreatePage Action = "CREATE_PAGE"
	GetPost    Action = "GET_POST"
	ListPosts  Action = "LIST_POSTS"
	GetPage    Action = "GET_PAGE"
	ListPages  Action = "LIST_PAGES"
)

var RolePermissions = map[string][]Action{
	Admin: {
		CreateUser,
		GetUser,
		ListUsers,
		CreatePost,
		CreatePage,
		GetPost,
		ListPosts,
		GetPage,
		ListPages,
	},
	User: {
		GetUser,
		CreatePost,
		CreatePage,
		GetPost,
		ListPosts,
		GetPage,
		ListPages,
	},
}
