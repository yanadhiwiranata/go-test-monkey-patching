package human

import (
	"strconv"

	"github.com/yanadhiwiranata/go-test-monkey-patching/domain"
)

type Human struct {
	id   int
	name string
}

func NewHuman(id int, name string) domain.UserFunction {
	return &Human{
		id:   id,
		name: name,
	}
}

func (h *Human) GetId() int {
	return h.id
}
func (h *Human) GetName() string {
	return h.name
}
func (h *Human) GetIDAndName() string {
	return strconv.Itoa(h.GetId()) + h.GetName()
}
