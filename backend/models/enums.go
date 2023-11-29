package models

const (
	MAX_WALLET_COUNT int64  = 3
	MOMO_WALLET_TYPE string = "MOMO"
	CARD_WALLET_TYPE string = "CARD"
)

// User Success Enums
const (
	CREATE_USER_SUCCESS string = "User created successfully"
	UPDATE_USER_SUCCESS string = "User updated successfully"
	GET_USER_SUCCESS    string = "User(s) retrieved successfully"
	DELETE_USER_SUCCESS string = "User deleted successfully"
)

// User Error Enums
const (
	CREATE_USER_ERR string = "Something went wrong when creating user"
	UPDATE_USER_ERR string = "Something went wrong when updating user"
	GET_USER_ERR    string = "Something went wrong when retrieving user data"
	DELETE_USER_ERR string = "Something went wrong when deleting user"
	USER_NOT_FOUND  string = "User does not exist!"
)

// Wallet Success Enums
const (
	CREATE_WALLET_SUCCESS string = "Wallet created successfully"
	UPDATE_WALLET_SUCCESS string = "Wallet updated successfully"
	GET_WALLET_SUCCESS    string = "Wallet(s) retrieved successfully"
	DELETE_WALLET_SUCCESS string = "Wallet deleted successfully"
)

// Wallet Error Enums
const (
	CREATE_WALLET_ERR string = "Something went wrong when creating wallet"
	UPDATE_WALLET_ERR string = "Something went wrong when updating wallet"
	GET_WALLET_ERR    string = "Something went wrong when retrieving wallet data"
	DELETE_WALLET_ERR string = "Something went wrong when deleting wallet"
	WALLET_NOT_FOUND  string = "Wallet does not exist!"
)
