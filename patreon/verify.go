package patreon

import (
	"crypto"
	"crypto/hmac"
	"encoding/hex"
	"github.com/TicketsBot/whitelabelproxy/config"
)

func VerifyHmac(content, signature []byte) bool {
	mac := hmac.New(crypto.MD5.New, []byte(config.Conf.Patreon.Secret))
	mac.Write(content)

	var h []byte
	hex.Encode(mac.Sum(nil), h)

	return hmac.Equal(h, signature)
}
