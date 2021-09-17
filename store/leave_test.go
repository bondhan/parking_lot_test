package store

import (
	"parking_lot/errors"
	"parking_lot/schema"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("leave a slot tests", func() {
	var (
		connection Store
	)
	connection = NewStore()
	It("Tear Down Store Data", func() {
		TearDown()
	})

	Context("leave store execute", func() {
		TearDown()

		It("leave help", func() {
			cmd := &schema.Command{
				Command:   "leave",
				Arguments: []string{"help"},
			}
			res, err := connection.Leave().Execute(cmd)
			Ω(err).ShouldNot(HaveOccurred())
			Expect(res).To(Equal(schema.CMDLeaveHint))
		})
		It("leave a slot but parking lot not yet created", func() {
			cmd := &schema.Command{
				Command:   "leave",
				Arguments: []string{"1"},
			}
			_, err := connection.Leave().Execute(cmd)
			Ω(err).Should(HaveOccurred())
			Expect(err).To(Equal(errors.ErrNoParkingLot))
		})
		//It("leave a slot but car is empty", func() {
		//	cmd := &schema.Command{
		//		Command: "create_parking_lot",
		//	}
		//
		//	cmd.Arguments = []string{"5"}
		//	res, err := connection.CreateParkingLot().Execute(cmd)
		//	Ω(err).ShouldNot(HaveOccurred())
		//	Expect(res).To(Equal(fmt.Sprintf(ParkinglotCreatedInfo, 5)))
		//
		//	cmd = &schema.Command{
		//		Command:   "leave",
		//		Arguments: []string{"1"},
		//	}
		//	res, err = connection.Leave().Execute(cmd)
		//	Ω(err).ShouldNot(HaveOccurred())
		//	Expect(res).To(fmt.Sprintf(SlotIsFreeInfo, slot.GetID()))
		//})
	})
})
