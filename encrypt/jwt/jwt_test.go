package jwt

import (
	"testing"
	"log"
)

func TestGenJwtTokenForOnlineSupport(t *testing.T){

	if token, err := GenJwtTokenForOnlineSupport(); err != nil {
		t.Fatal(err)
	} else {
		if err = ValidateJwtTokenForOnlineSupport(token); err != nil {
			t.Fatal(err)
		}
		log.Printf("token is:  %s", token)
	}

}

/*func TestValidateJwtTokenForOnlineSupport(t *testing.T){
	var tokenString = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjMzMjg2NDh9.nDvQOyxt6XGVkgxjbUDIZYTRumkkcr8IZUIeq7G72kk"
	if err := ValidateJwtTokenForOnlineSupport(tokenString); err != nil {
		t.Fatal(err)
	}
}*/
