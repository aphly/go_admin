package helper

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"go_admin/app/core"
	"strings"
)

type JWT struct {
	header    string
	Payload   string
	signature string
}

type PayloadData struct {
	Expire    int64
	Uid       core.Uint
	TokenType string
}

func encodeBase64(data string) string {
	return base64.RawURLEncoding.EncodeToString([]byte(data))
}

func generateSignature(key []byte, data []byte) (string, error) {
	hash := hmac.New(sha256.New, key)
	_, err := hash.Write(data)
	if err != nil {
		return "", err
	}
	return encodeBase64(string(hash.Sum(nil))), nil
}

func CreateToken(key []byte, payloadData PayloadData) (string, error) {
	header := `{"alg":"HS256","typ":"JWT"}`
	payload, jsonErr := json.Marshal(payloadData)
	if jsonErr != nil {
		return "", errors.New("json错误")
	}
	encodedHeader := encodeBase64(header)
	encodedPayload := encodeBase64(string(payload))
	HeaderAndPayload := encodedHeader + "." + encodedPayload
	signature, err := generateSignature(key, []byte(HeaderAndPayload))
	if err != nil {
		return "", err
	}
	return HeaderAndPayload + "." + signature, nil
}

func ParseJwt(token string, key []byte) (PayloadData, error) {
	payloadData := PayloadData{}
	jwtParts := strings.Split(token, ".")
	if len(jwtParts) != 3 {
		return payloadData, errors.New("非法token")
	}
	encodedHeader := jwtParts[0]
	encodedPayload := jwtParts[1]
	signature := jwtParts[2]

	confirmSignature, err := generateSignature(key, []byte(encodedHeader+"."+encodedPayload))
	if err != nil {
		return payloadData, errors.New("生成签名错误")
	}
	if signature != confirmSignature {
		return payloadData, errors.New("token验证失败")
	}
	dstPayload, _ := base64.RawURLEncoding.DecodeString(encodedPayload)

	err = json.Unmarshal(dstPayload, &payloadData)
	if err != nil {
		return payloadData, err
	}
	//if payloadData.Expire < time.Now().Unix() {
	//	return payloadData, errors.New("token过期了")
	//}
	//return &JWT{encodedHeader, string(dstPayload), signature}, nil
	return payloadData, nil
}
