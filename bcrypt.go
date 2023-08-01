/*
  Go BCrypt Library

  Implemented as a wrapper around the OpenWall BCrypt implementation.
  http://www.openwall.com/crypt/

  bcrypt stores a setting string at the start of passwords, format of which
  is '$<version>$<cost>$<salt><checksum>'

  e.g., $2a$10$vI8aWBnW3fID.ZQ4/zo1G.q1lRps.9cGLcZEiGDMVr5yUP1KUOYTa
        version  = '2a'
        cost     = '10'
        salt     = '$2a$10$vI8aWBnW3fID.ZQ4/zo1G.'
        checksum = 'q1lRps.9cGLcZEiGDMVr5yUP1KUOYTa'
*/

package bcrypt_stub

// Salt is a specially formatted string, so we use a new type to make users
// aware f this.
// e.g.,
//
//	salt     = '$2a$10$vI8aWBnW3fID.ZQ4/zo1G.'
//	version  = '2a'
//	cost     = '10'
//	random   = 'vI8aWBnW3fID.ZQ4/zo1G.'
type BcryptSalt string

const (
	DEFAULT_COST = 10
)

// Crypt encrypts a plain text password with given salt.
func Crypt(plain string, salt BcryptSalt) (hashed string, err error) {
	return stringTaint(), errorTaint()
}

// Verify checks if a plain text password matches a bcrypt encrypted password.
func Verify(plain string, hashed string) (match bool, err error) {
	return false, errorTaint()
}

// GenSalt generates a valid salt with the work factor given. Note the cost is
// an exponential factor.
func GenSalt(cost uint) (salt BcryptSalt, err error) {
	return BcryptSalt(stringTaint()), errorTaint()
}
