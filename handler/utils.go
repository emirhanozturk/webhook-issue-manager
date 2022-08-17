package handlers

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/webhook-issue-manager/model"
)

var (
	secretKey       = "secretKey"
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("expired Token")
)

func verifyToken(token string) (*model.Payload, error) {

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &model.Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*model.Payload)
	if !ok {
		return nil, ErrInvalidToken
	}
	return payload, nil

}

// type JSON json.RawMessage

// // Scan scan value into Jsonb, implements sql.Scanner interface
// func (j *JSON) Scan(value interface{}) error {
//   bytes, ok := value.([]byte)
//   if !ok {
//     return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
//   }

//   result := json.RawMessage{}
//   err := json.Unmarshal(bytes, &result)
//   *j = JSON(result)
//   return err
// }

// // Value return json value, implement driver.Valuer interface
// func (j JSON) Value() (driver.Value, error) {
//   if len(j) == 0 {
//     return nil, nil
//   }
//   return json.RawMessage(j).MarshalJSON()
// }
