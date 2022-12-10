package human_test

import (
	"reflect"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
	"github.com/undefinedlabs/go-mpatch"
	_human "github.com/yanadhiwiranata/go-test-monkey-patching/implementation"
)

func TestGetId(t *testing.T) {
	human := _human.NewHuman(1, "yan")
	assert.Equal(t, 1, human.GetId())
}

func TestGetName(t *testing.T) {
	human := _human.NewHuman(1, "yan")
	assert.Equal(t, "yan", human.GetName())
}

func TestGetIDAndName(t *testing.T) {
	human := _human.NewHuman(1, "yan")
	assert.Equal(t, "1yan", human.GetIDAndName())
}

func TestMockGetId(t *testing.T) {
	human := _human.NewHuman(1, "yan")
	var patch *mpatch.Patch
	var err error
	patch, err = mpatch.PatchInstanceMethodByName(reflect.TypeOf(human), "GetId", func(m *_human.Human) int {
		patch.Unpatch()
		defer patch.Patch()
		return 2
	})
	assert.Equal(t, 2, human.GetId())
	assert.NoError(t, err)
}

func TestMockGetIDAndName(t *testing.T) {
	human := _human.NewHuman(1, "yan")
	var patch *mpatch.Patch
	var err error
	patch, err = mpatch.PatchInstanceMethodByName(reflect.TypeOf(human), "GetIDAndName", func(m *_human.Human) string {
		patch.Unpatch()
		defer patch.Patch()
		return "2adhi"
	})
	assert.Equal(t, "2adhi", human.GetIDAndName())
	assert.NoError(t, err)
}

func TestMockGoMonekyGetId(t *testing.T) {
	human := _human.NewHuman(1, "yan")
	var patch *mpatch.Patch
	var err error
	patch, err = mpatch.PatchInstanceMethodByName(reflect.TypeOf(human), "GetId", func(m *_human.Human) int {
		patch.Unpatch()
		defer patch.Patch()
		return 2
	})
	assert.Equal(t, 2, human.GetId())
	assert.NoError(t, err)
}

func TestMockGetID2(t *testing.T) {
	human := _human.NewHuman(1, "yan")
	gomonkey.ApplyMethod(reflect.TypeOf(human), "GetId", func(m *_human.Human) int {
		return 2
	})
	assert.Equal(t, 2, human.GetId())
}
