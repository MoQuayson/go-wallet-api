package enums

type ResponseMsg string

// User Success Enums
const (
	ResponseMsg_CREATE_USER_SUCCESS = "User created successfully"
	ResponseMsg_UPDATE_USER_SUCCESS = "User updated successfully"
	ResponseMsg_GET_USER_SUCCESS    = "User(s) retrieved successfully"
	ResponseMsg_DELETE_USER_SUCCESS = "User deleted successfully"

	ResponseMsg_CREATE_USER_ERR string = "Something went wrong when creating user"
	ResponseMsg_UPDATE_USER_ERR string = "Something went wrong when updating user"
	ResponseMsg_GET_USER_ERR    string = "Something went wrong when retrieving user data"
	ResponseMsg_DELETE_USER_ERR string = "Something went wrong when deleting user"
	ResponseMsg_USER_NOT_FOUND  string = "User does not exist!"
)
