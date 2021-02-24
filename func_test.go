package protodoc

import (
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/compiler/protogen"
	"testing"
)

func contains(m *protogen.Message, ms []*protogen.Message) bool {
	for _, v := range ms {
		if v == m {
			return true
		}
	}
	return false
}

func TestAllMessages(t *testing.T) {
	var m1, m2 protogen.Message
	m3 := protogen.Message{Messages: []*protogen.Message{&m1}}

	f := &protogen.File{
		Messages: []*protogen.Message{&m3, &m2},
	}

	var ms []*protogen.Message

	for m := range getMessages(f) {
		ms = append(ms, m)
	}

	assert.Equal(t, 3, len(ms))
	assert.True(t, contains(&m1, ms))
	assert.True(t, contains(&m2, ms))
	assert.True(t, contains(&m3, ms))
}

func TestCommentString(t *testing.T) {
	cs := protogen.CommentSet{
		LeadingDetached: []protogen.Comments{
			"Leading detached 1",
			"Leading detached 2",
		},
		Leading: "Leading comment",
		Trailing: "Trailing Comment",
	}

	exp := `Leading detached 1
Leading detached 2
Leading comment
Trailing Comment`
	assert.Equal(t, exp, commentString(cs))
}
