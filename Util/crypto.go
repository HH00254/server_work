package util

import "github.com/ProtonMail/gopenpgp/v2/helper"

type PgpDecryptor struct {
	privateKey string
}

func NewPgpDecrypter(privateKey string) PgpDecryptor {
	return PgpDecryptor{
		privateKey: privateKey,
	}
}

func (pgp PgpDecryptor) Decrypt(armorKey string) (string, error) {
	return helper.DecryptMessageArmored(pgp.privateKey, nil, armorKey)
}
