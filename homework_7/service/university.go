package service

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

//MD5加密
func CryptographyNow(data string) (string, int64) {
	Dc.Data=data
	temTime:=time.Now().Unix()
	Dc.Cryptography(temTime)
	return Dc.Result,temTime
}
func (dc *DataCryptographyMD5)Cryptography(salt int64){
	has:=md5.New()
	has.Write([]byte(dc.Data))
	tem:= has.Sum([]byte(string(salt)))
	dc.Result=hex.EncodeToString(tem)
}