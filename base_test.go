package annie

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

func TestAnnie(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Annie Suit")
}
