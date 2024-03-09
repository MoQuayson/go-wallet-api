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

	ResponseMsg_AuthenticateUserErr string = "Something went wrong when authenticating user"
	ResponseMsg_InvalidCredentials  string = "Invalid credentials. Try again."
	ResponseMsg_ValidCredentials    string = "User authenticated successfully"
	ResponseMsg_UnAuthorizedUser    string = "Unauthorized User"
	ResponseMsg_UnAuthenticatedUser string = "User not authenticated"
	ResponseMsg_CreateWalletSuccess string = "Wallet created successfully"
	ResponseMsg_UpdateWalletSuccess string = "Wallet updated successfully"
	ResponseMsg_GetWalletSuccess    string = "Wallet(s) retrieved successfully"
	ResponseMsg_DeleteWalletSuccess string = "Wallet deleted successfully"
	ResponseMsg_CreateWalletErr     string = "Something went wrong when creating wallet"
	ResponseMsg_UpdateWalletErr     string = "Something went wrong when updating wallet"
	ResponseMsg_GetWalletErr        string = "Something went wrong when retrieving wallet data"
	ResponseMsg_DeleteWalletErr     string = "Something went wrong when deleting wallet"
	ResponseMsg_WalletNotFound      string = "Wallet does not exist!"
)
