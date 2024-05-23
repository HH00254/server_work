package server

import util "github.com/HH00254/server_work/Util"

type Controller struct {
	publicKey string
	decryptor util.PgpDecryptor
}

func NewController(publicKey string, decryptor util.PgpDecryptor) Controller {
	return Controller{
		publicKey: publicKey,
		decryptor: decryptor,
	}
}
