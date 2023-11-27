package pakarbibackend

import (
	"fmt"
	"testing"

	"github.com/aiteung/atdb"
	"github.com/whatsauth/watoken"
	"go.mongodb.org/mongo-driver/bson"
)

func TestCreateNewUserRole(t *testing.T) {
	var userdata User
	userdata.Username = "pakarbi"
	userdata.NPM = "1214000"
	userdata.Password = "pakarbipass"
	userdata.PasswordHash = "pakarbipass"
	userdata.Email = "pakarbi2023@gmail.com"
	userdata.Role = "user"
	mconn := SetConnection("MONGOSTRING", "pakarbidbmongo")
	CreateNewUserRole(mconn, "user", userdata)
}

// func TestDeleteUser(t *testing.T) {
// 	mconn := SetConnection("MONGOSTRING", "pasabar13")
// 	var userdata User
// 	userdata.Username = "lolz"
// 	DeleteUser(mconn, "user", userdata)
// }

func CreateNewUserToken(t *testing.T) {
	var userdata User
	userdata.Username = "pakarbi"
	userdata.NPM = "1214000"
	userdata.Password = "pakarbipass"
	userdata.PasswordHash = "pakarbipass"
	userdata.Email = "pakarbi2023@gmail.com"
	userdata.Role = "user"

	// Create a MongoDB connection
	mconn := SetConnection("MONGOSTRING", "pakarbidbmongo")

	// Call the function to create a user and generate a token
	err := CreateUserAndAddToken("your_private_key_env", mconn, "user", userdata)

	if err != nil {
		t.Errorf("Error creating user and token: %v", err)
	}
}

func TestGFCPostHandlerUser(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "pakarbidbmongo")
	var userdata User
	userdata.Username = "pakarbi"
	userdata.NPM = "1214000"
	userdata.Password = "pakarbipass"
	userdata.PasswordHash = "pakarbipass"
	userdata.Email = "pakarbi2023@gmail.com"
	userdata.Role = "user"
	CreateNewUserRole(mconn, "user", userdata)
}

func TestGeneratePasswordHash(t *testing.T) {
	passwordhash := "pakarbipass"
	hash, _ := HashPassword(passwordhash) // ignore error for the sake of simplicity

	fmt.Println("Password:", passwordhash)
	fmt.Println("Hash:    ", hash)
	match := CheckPasswordHash(passwordhash, hash)
	fmt.Println("Match:   ", match)
}
func TestGeneratePrivateKeyPaseto(t *testing.T) {
	privateKey, publicKey := watoken.GenerateKey()
	fmt.Println(privateKey)
	fmt.Println(publicKey)
	hasil, err := watoken.Encode("pakarbipass", privateKey)
	fmt.Println(hasil, err)
}

func TestHashFunction(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "pakarbidbmongo")
	var userdata User
	userdata.Username = "pakarbi"
	userdata.PasswordHash = "pakarbipass"

	filter := bson.M{"username": userdata.Username}
	res := atdb.GetOneDoc[User](mconn, "user", filter)
	fmt.Println("Mongo User Result: ", res)
	hash, _ := HashPassword(userdata.PasswordHash)
	fmt.Println("Hash Password : ", hash)
	match := CheckPasswordHash(userdata.PasswordHash, res.PasswordHash)
	fmt.Println("Match:   ", match)

}

func TestIsPasswordValid(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "pakarbidbmongo")
	var userdata User
	userdata.Username = "pakarbi"
	userdata.PasswordHash = "pakarbipass"

	anu := IsPasswordValid(mconn, "user", userdata)
	fmt.Println(anu)
}

func TestUserFix(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "pakarbidbmongo")
	var userdata User
	userdata.Username = "pakarbi"
	userdata.NPM = "1214000"
	userdata.Password = "pakarbipass"
	userdata.PasswordHash = "pakarbipass"
	userdata.Email = "pakarbi2023@gmail.com"
	userdata.Role = "user"
	CreateUser(mconn, "user", userdata)
}
