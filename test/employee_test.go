package uint

import (
	"testing"
	"github.com/asaskevich/govalidator"
	"github.com/ThanawitGZS/Backend-Test/entity"
	. "github.com/onsi/gomega"
)

func TestEmployee(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`employee is valid`, func(t *testing.T){
		employee := entity.Employee{
			Name: "A",
		}

		ok, err := govalidator.ValidateStruct(employee)
		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})

	t.Run(`name is required`, func(t *testing.T){
		employee := entity.Employee{
			Name: "",
		}

		ok, err := govalidator.ValidateStruct(employee)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Name is required"))
	})

}