package enums

type ResponseMsg string

// User Success Enums
const (
	ResponseMsg_CREATE_USER_SUCCESS = "User created successfully"
	ResponseMsg_UPDATE_USER_SUCCESS = "User updated successfully"
	ResponseMsg_GET_USER_SUCCESS    = "User(s) retrieved successfully"
	ResponseMsg_DELETE_USER_SUCCESS = "User deleted successfully"

	ResponseMsg_CREATE_USER_ERR ResponseMsg = "Something went wrong when creating user"
	ResponseMsg_UPDATE_USER_ERR ResponseMsg = "Something went wrong when updating user"
	ResponseMsg_GET_USER_ERR    ResponseMsg = "Something went wrong when retrieving user data"
	ResponseMsg_DELETE_USER_ERR ResponseMsg = "Something went wrong when deleting user"
	ResponseMsg_USER_NOT_FOUND  ResponseMsg = "User does not exist!"

	ResponseMsg_AuthenticateUserErr ResponseMsg = "Something went wrong when authenticating user"
	ResponseMsg_InvalidCredentials  ResponseMsg = "Invalid credentials. Try again."
	ResponseMsg_ValidCredentials    ResponseMsg = "User authenticated successfully"
	ResponseMsg_UnAuthorizedUser    ResponseMsg = "Unauthorized User"
	ResponseMsg_UnAuthenticatedUser ResponseMsg = "User not authenticated"
	ResponseMsg_CreateWalletSuccess ResponseMsg = "Wallet created successfully"
	ResponseMsg_UpdateWalletSuccess ResponseMsg = "Wallet updated successfully"
	ResponseMsg_GetWalletSuccess    ResponseMsg = "Wallet(s) retrieved successfully"
	ResponseMsg_DeleteWalletSuccess ResponseMsg = "Wallet deleted successfully"
	ResponseMsg_CreateWalletErr     ResponseMsg = "Something went wrong when creating wallet"
	ResponseMsg_UpdateWalletErr     ResponseMsg = "Something went wrong when updating wallet"
	ResponseMsg_GetWalletErr        ResponseMsg = "Something went wrong when retrieving wallet data"
	ResponseMsg_DeleteWalletErr     ResponseMsg = "Something went wrong when deleting wallet"
	ResponseMsg_WalletNotFound      ResponseMsg = "Wallet does not exist!"

	ResponseMsg_SignUpUserErr    ResponseMsg = "Something went wrong when creating an account"
	ResponseMsg_SignUpUserSucess ResponseMsg = "Account created successfully"
)
