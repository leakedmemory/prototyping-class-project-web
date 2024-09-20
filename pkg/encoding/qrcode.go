package encoding

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/skip2/go-qrcode"
)

func GenerateQRCode(leashID string) ([]byte, error) {
	domain := os.Getenv("DOMAIN")
	url := fmt.Sprintf("%s/pet?leash-id=%s", domain, leashID)
	qr, err := qrcode.Encode(url, qrcode.Medium, 256)
	if err != nil {
		return nil, err
	}

	return qr, nil
}
