package unit

import (
	"testing"
	"github.com/ThanawitGZS/Backend-Test/entity"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestRoom(t *testing.T){
	g:= NewGomegaWithT(t)

	e := entity.Employee{
		Name:"Thanawit",
		Email: "A@aaa.com",
	}

	t.Run(`Room is valid`, func(t *testing.T){
		room := entity.Room {
			RoomName: "A",
			EmployeeID: uint(1),
			Employee: e,
		}

		ok, err := govalidator.ValidateStruct(room)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})

	t.Run(`room_name is required`, func(t *testing.T){
		room := entity.Room{
			RoomName: "",//ผิดตรงนี้
			EmployeeID: uint(1),
			Employee: e,
		}

		ok, err := govalidator.ValidateStruct(room)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("RoomName is required"))
	})

	t.Run(`employee_id is required`, func(t *testing.T){
		room := entity.Room{
			RoomName: "A",
			EmployeeID: uint(0),
			Employee: e,
		}

		ok, err := govalidator.ValidateStruct(room)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("EmployeeID is required"))
	})
}