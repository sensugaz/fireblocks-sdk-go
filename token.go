package fireblockssdk

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/averagesecurityguy/random"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type claims struct {
	Uri      string `json:"uri"`
	BodyHash string `json:"bodyHash"`
	Nonce    int64  `json:"nonce"`
	jwt.StandardClaims
}

func (r request) signJwt(path string, body interface{}) string {
	var (
		timestamp = time.Now()
		nonce, _  = random.Int64()
	)

	bodyJson, _ := json.Marshal(body)
	bodyHash := sha256.Sum256(bodyJson)

	claims := claims{
		Uri:      path,
		Nonce:    nonce,
		BodyHash: hex.EncodeToString(bodyHash[:]),
	}
	claims.Subject = r.options.apiKey
	claims.IssuedAt = timestamp.Unix()
	claims.ExpiresAt = timestamp.Add(5 * time.Second).Unix()

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(r.options.privateKey))
	if err != nil {
		return ""
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(signKey)
	if err != nil {
		return ""
	}
	
	return token
}
