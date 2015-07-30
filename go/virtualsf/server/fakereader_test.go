package server

import (
	"bufio"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestFakeFileReaderInterface(t *testing.T) {

	Convey("FAKE FILE", t, func() {

		Convey("gera conteúdo do tamanho desejado", func() {
			fake := NewFakeReader(1024)
			reader := bufio.NewReader(fake)
			content, _ := reader.ReadString(' ')

			So(len(content), ShouldEqual, 1024)
		})

	})

}
