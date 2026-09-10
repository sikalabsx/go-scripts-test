package gravatar_image_from_email

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

func GravatarImageURL(email string) string {
	hash := md5.Sum([]byte(strings.ToLower(strings.TrimSpace(email))))
	return fmt.Sprintf("https://www.gravatar.com/avatar/%s", hex.EncodeToString(hash[:]))
}

func PrintGravatarImageURL(email string) {
	fmt.Println(GravatarImageURL(email))
}
