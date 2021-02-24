package protodoc

import (
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"sort"
	"strings"
)

func _getMessages(m *protogen.Message, c chan *protogen.Message) {
	c <- m
	for _, msg := range m.Messages {
		_getMessages(msg, c)
	}
}

func getMessages(f *protogen.File) chan *protogen.Message {
	c := make(chan *protogen.Message)
	go func() {
		defer close(c)
		for _, msg := range f.Messages {
			_getMessages(msg, c)
		}
	}()
	return c
}

func getEnums(f *protogen.File) chan *protogen.Enum {
	c := make(chan *protogen.Enum)

	go func() {
		defer close(c)
		for _, e := range f.Enums {
			c <- e
		}
		for m := range getMessages(f) {
			for _, e := range m.Enums {
				c <- e
			}
		}
	}()
	return c
}

func commentString(cs protogen.CommentSet) string {
	var b strings.Builder
	for _, l := range cs.LeadingDetached {
		if l != "" {
			fmt.Fprintln(&b, string(l))
		}
	}
	if cs.Leading != "" {
		fmt.Fprintln(&b, string(cs.Leading))
	}
	if cs.Trailing != "" {
		fmt.Fprintln(&b, string(cs.Trailing))
	}

	return strings.TrimSpace(b.String())
}

func messageId(m protoreflect.Descriptor) string {
	n := strings.ToLower(string(m.FullName()))
	return strings.Replace(n, ".", "-", -1)
}

func fileLoc(f protoreflect.FileDescriptor) string {
	name := f.Path()
	withoutExt := strings.Split(name, ".")[0]
	v := fmt.Sprintf("%s.html", withoutExt)
	if v[0] != '/' {
		return "/" + v
	}
	return v
}

type Files []*protogen.File

func (fs Files) Len() int {
	return len(fs)
}

func (fs Files) Less(i, j int) bool {
	return fs[i].Desc.Path() < fs[j].Desc.Path()
}

func (fs Files) Swap(i, j int) {
	tmp := fs[i]
	fs[i] = fs[j]
	fs[j] = tmp
}

func sortedFiles(files []*protogen.File) []*protogen.File {

	fs := make(Files, len(files))
	copy(fs, files)
	sort.Sort(fs)

	return fs
}
