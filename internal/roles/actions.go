package roles

type Action string

const (
	CreateUser Action = "CREATE_USER"
	GetUser    Action = "GET_USER"
	ListUsers  Action = "LIST_USERS"
	UpdateUser Action = "UPDATE_USER"
	DeleteUser Action = "DELETE_USER"
	CreatePost Action = "CREATE_POST"
	GetPost    Action = "GET_POST"
	ListPosts  Action = "LIST_POSTS"
	UpdatePost Action = "UPDATE_POST"
	DeletePost Action = "DELETE_POST"
	CreatePage Action = "CREATE_PAGE"
	GetPage    Action = "GET_PAGE"
	ListPages  Action = "LIST_PAGES"
	UpdatePage Action = "UPDATE_PAGE"
	DeletePage Action = "DELETE_PAGE"
)

var RolePermissions = map[string][]Action{
	Admin: {
		CreateUser,
		GetUser,
		ListUsers,
		UpdateUser,
		DeleteUser,
		CreatePost,
		GetPost,
		ListPosts,
		UpdatePost,
		DeletePost,
		CreatePage,
		GetPage,
		ListPages,
		UpdatePage,
		DeletePage,
	},
	User: {
		GetUser,
		UpdateUser, // only him/herself
		CreatePost,
		GetPost,
		ListPosts,
		UpdatePost,
		DeletePost,
		CreatePage,
		GetPage,
		ListPages,
		UpdatePage,
		DeletePage,
	},
}
