package pakarbibackend

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/whatsauth/watoken"
	"go.mongodb.org/mongo-driver/bson"
)

// <--- ini parkiran --->

// parkiran post
func GCFInsertParkiran(publickey, MONGOCONNSTRINGENV, dbname, colluser, collparkiran string, r *http.Request) string {
	var response Credential
	response.Status = false
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var userdata User
	gettoken := r.Header.Get("token")
	if gettoken == "" {
		response.Message = "Missing token in headers"
	} else {
		// Process the request with the "Login" token
		checktoken := watoken.DecodeGetId(os.Getenv(publickey), gettoken)
		userdata.NPM = checktoken
		if checktoken == "" {
			response.Message = "Invalid token"
		} else {
			user2 := FindUser(mconn, colluser, userdata)
			if user2.Role == "user" {
				var dataparkiran Parkiran
				err := json.NewDecoder(r.Body).Decode(&dataparkiran)
				if err != nil {
					response.Message = "Error parsing application/json: " + err.Error()
				} else {
					insertParkiran(mconn, collparkiran, Parkiran{
						ParkiranId:     dataparkiran.ParkiranId,
						Nama:           dataparkiran.Nama,
						NPM:            dataparkiran.NPM,
						Jurusan:        dataparkiran.Jurusan,
						NamaKendaraan:  dataparkiran.NamaKendaraan,
						NomorKendaraan: dataparkiran.NomorKendaraan,
						JenisKendaraan: dataparkiran.JenisKendaraan,
						Status:         dataparkiran.Status,
					})
					response.Status = true
					response.Message = "Berhasil Insert Parkiran"
				}
			} else {
				response.Message = "Anda tidak bisa Insert data karena bukan admin"
			}
		}
	}
	return GCFReturnStruct(response)
}

// delete parkiran
func GCFDeleteParkiran(publickey, MONGOCONNSTRINGENV, dbname, colluser, collparkiran string, r *http.Request) string {

	var respon Credential
	respon.Status = false
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var userdata User

	gettoken := r.Header.Get("token")
	if gettoken == "" {
		respon.Message = "Missing token in headers"
	} else {
		// Process the request with the "Login" token
		checktoken := watoken.DecodeGetId(os.Getenv(publickey), gettoken)
		userdata.Email = checktoken
		if checktoken == "" {
			respon.Message = "Invalid token"
		} else {
			user2 := FindUser(mconn, colluser, userdata)
			if user2.Role == "user" {
				var dataparkiran Parkiran
				err := json.NewDecoder(r.Body).Decode(&dataparkiran)
				if err != nil {
					respon.Message = "Error parsing application/json: " + err.Error()
				} else {
					DeleteParkiran(mconn, collparkiran, dataparkiran)
					respon.Status = true
					respon.Message = "Berhasil Delete Parkiran"
				}
			} else {
				respon.Message = "Anda tidak bisa Delete data karena bukan admin"
			}
		}
	}
	return GCFReturnStruct(respon)
}

// update parkiran
func GCFUpdateParkiran(publickey, MONGOCONNSTRINGENV, dbname, colluser, collparkiran string, r *http.Request) string {
	var response Credential
	response.Status = false
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var userdata User

	gettoken := r.Header.Get("token")
	if gettoken == "" {
		response.Message = "Missing token in Headers"
	} else {
		checktoken := watoken.DecodeGetId(os.Getenv(publickey), gettoken)
		userdata.NPM = checktoken
		if checktoken == "" {
			response.Message = "Invalid token"
		} else {
			user2 := FindUser(mconn, colluser, userdata)
			if user2.Role == "user" {
				var dataparkiran Parkiran
				err := json.NewDecoder(r.Body).Decode(&dataparkiran)
				if err != nil {
					response.Message = "Error parsing application/json: " + err.Error()

				} else {
					UpdatedParkiran(mconn, collparkiran, bson.M{"id": dataparkiran.ID}, dataparkiran)
					response.Status = true
					response.Message = "Berhasil Update Parkiran"
					GCFReturnStruct(CreateResponse(true, "Success Update Parkiran", dataparkiran))
				}
			} else {
				response.Message = "Anda tidak bisa Update data karena bukan admin"
			}

		}
	}
	return GCFReturnStruct(response)
}

// get all parkiran
func GCFGetAllParkiran(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	dataparkiran := GetAllParkiran(mconn, collectionname)
	if dataparkiran != nil {
		return GCFReturnStruct(CreateResponse(true, "success Get All Parkiran", dataparkiran))
	} else {
		return GCFReturnStruct(CreateResponse(false, "Failed Get All Parkiran", dataparkiran))
	}
}

// get all parkiran by id
func GCFGetAllParkiranID(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)

	var dataparkiran Parkiran
	err := json.NewDecoder(r.Body).Decode(&dataparkiran)
	if err != nil {
		return err.Error()
	}

	parkiran := GetAllParkiranID(mconn, collectionname, dataparkiran)
	if parkiran != (Parkiran{}) {
		return GCFReturnStruct(CreateResponse(true, "Success: Get ID Parkiran", dataparkiran))
	} else {
		return GCFReturnStruct(CreateResponse(false, "Failed to Get ID Parkiran", dataparkiran))
	}
}
