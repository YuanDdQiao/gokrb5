package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/jcmturner/goidentity.v1"
)

func TestImplementsInterface(t *testing.T) {
	//s := new(SPNEGOAuthenticator)
	var s SPNEGOAuthenticator
	a := new(goidentity.Authenticator)
	assert.Implements(t, a, s, "SPNEGOAuthenticator type does not implement the goidentity.Authenticator interface")
}
