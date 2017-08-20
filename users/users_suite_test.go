package users_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rtravitz/culture_knights/db"

	"testing"
)

var testDB *db.DB

func TestUsers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Users Suite")
}

var _ = BeforeSuite(func() {
	var err error
	testDB, err = db.New(os.Getenv("CULTURE_DB_TEST"))
	Expect(err).NotTo(HaveOccurred())
})
