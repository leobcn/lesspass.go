package tests

import (
	"math/big"
	"testing"

	lesspass ".."
)

func testCalcEntropyPbkdf2WithDefaultParams(t *testing.T) {
	var site = "example.org"
	var login = "contact@example.org"
	var masterPassword = "password"
	var passwordProfile = lesspass.PasswordProfile{}
	var entropy = lesspass.CalcEntropy(site, login, masterPassword, passwordProfile)
	if string(entropy) != "dc33d431bce2b01182c613382483ccdb0e2f66482cbba5e9d07dab34acc7eb1e" {
		t.Fatal("entropy != \"dc33d431bce2b01182c613382483ccdb0e2f66482cbba5e9d07dab34acc7eb1e\"")
	}
}

func testCalcEntropyWithDifferentOptions(t *testing.T) {
	var site = "example.org"
	var login = "contact@example.org"
	var masterPassword = "password"
	var passwordProfile = lesspass.PasswordProfile{"iterations": 8192, "keylen": 16, "digest": "sha512"}
	var entropy = lesspass.CalcEntropy(site, login, masterPassword, passwordProfile)
	if string(entropy) != "fff211c16a4e776b3574c6a5c91fd252" {
		t.Fatal("entropy != \"fff211c16a4e776b3574c6a5c91fd252\"")
	}
}

func testCalcEntropyWithCounterOne(t *testing.T) {
	var site = "example.org"
	var login = "contact@example.org"
	var masterPassword = "password"
	var passwordProfile = lesspass.PasswordProfile{"iterations": 1, "keylen": 16}
	var entropy = lesspass.CalcEntropy(site, login, masterPassword, passwordProfile)
	if string(entropy) != "d3ec1e988dd0b3640c7491cd2c2a88b5" {
		t.Fatal("entropy != \"d3ec1e988dd0b3640c7491cd2c2a88b5\"")
	}
}

func testCalcEntropyWithCounterTwo(t *testing.T) {
	var site = "example.org"
	var login = "contact@example.org"
	var masterPassword = "password"
	var passwordProfile = lesspass.PasswordProfile{"iterations": 1, "keylen": 16, "counter": 2}
	var entropy = lesspass.CalcEntropy(site, login, masterPassword, passwordProfile)
	if string(entropy) != "ddfb1136260f930c21f6d72f6eddbd40" {
		t.Fatal("entropy != \"ddfb1136260f930c21f6d72f6eddbd40\"")
	}
}

func testConsumeEntropy(t *testing.T) {
	var value, entropy = lesspass.ConsumeEntropy("", big.NewInt(4*4+2), "abcd", 2)
	if string(value) != "ca" {
		t.Fatal("value != \"ca\"")
	}
	if entropy.Text(16) != "1" {
		t.Fatal("entropy != \"1\"")
	}
}
