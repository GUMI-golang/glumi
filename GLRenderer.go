package glumi

//#include <memory.h>
import "C"

import (
	"fmt"
	"strings"
	"unsafe"
	"github.com/go-gl/gl/all-core/gl"
)

type GLRender struct {
	glumi *GLUMIFullScreen
	//
	program    uint32
	frameimage uint32
	vao, vbo   uint32
	//
	pbo      [2]uint32
	pboIndex int
}

func (s *GLRender) init() error {
	err := s.newProgram()
	if err != nil {
		return err
	}
	s.newImage()
	s.newVAO()
	s.newVBO()
	s.newPBO()
	//
	gl.BindVertexArray(s.vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, s.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(ScreenModeling)*4*5, gl.Ptr(ScreenModeling), gl.STATIC_DRAW)
	s.attrib()
	return nil
}
func (s *GLRender) newProgram() error {
	program := gl.CreateProgram()
	gl.AttachShader(program, DefaultShader.Compiled.Vertex)
	gl.AttachShader(program, DefaultShader.Compiled.Fragment)
	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		return fmt.Errorf("failed to link program: %v", log)
	}
	s.program = program
	return nil
}
func (s *GLRender) newImage() {
	gl.GenTextures(1, &s.frameimage)
	gl.BindTexture(gl.TEXTURE_2D, s.frameimage)
	defer gl.BindTexture(gl.TEXTURE_2D, 0)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_READ_COLOR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
}
func (s *GLRender) newPBO() {
	gl.GenBuffers(2, &s.pbo[0])
	img := s.glumi.screen.RGBA()
	defer gl.BindBuffer(gl.PIXEL_UNPACK_BUFFER, 0)
	gl.BindBuffer(gl.PIXEL_UNPACK_BUFFER, s.pbo[0])
	gl.BufferData(gl.PIXEL_UNPACK_BUFFER, len(img.Pix), gl.PtrOffset(0), gl.STREAM_DRAW)
	gl.BindBuffer(gl.PIXEL_UNPACK_BUFFER, s.pbo[1])
	gl.BufferData(gl.PIXEL_UNPACK_BUFFER, len(img.Pix), gl.PtrOffset(0), gl.STREAM_DRAW)
}
func (s *GLRender) newVAO() {
	gl.GenVertexArrays(1, &s.vao)

}
func (s *GLRender) newVBO() {
	gl.GenBuffers(1, &s.vbo)
}
func (s *GLRender) attrib() {
	vertAttrib := uint32(gl.GetAttribLocation(s.program, gl.Str("vert\x00")))
	gl.EnableVertexAttribArray(vertAttrib)
	gl.VertexAttribPointer(vertAttrib, 3, gl.FLOAT, false, 5*4, gl.PtrOffset(0))

	texCoordAttrib := uint32(gl.GetAttribLocation(s.program, gl.Str("vertTexCoord\x00")))
	gl.EnableVertexAttribArray(texCoordAttrib)
	gl.VertexAttribPointer(texCoordAttrib, 2, gl.FLOAT, false, 5*4, gl.PtrOffset(3*4))
}
func (s *GLRender) Upload() {
	gl.UseProgram(s.program)
	img := s.glumi.screen.RGBA()
	curridx := s.pboIndex
	nextidx := (s.pboIndex + 1) % 2
	s.pboIndex = nextidx
	//texture binding
	gl.BindTexture(gl.TEXTURE_2D, s.frameimage)
	// downloading
	gl.BindBuffer(gl.PIXEL_UNPACK_BUFFER, s.pbo[curridx])
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(img.Rect.Dx()),
		int32(img.Rect.Dy()),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.PtrOffset(0),
	)
	// uploading
	gl.BindBuffer(gl.PIXEL_UNPACK_BUFFER, s.pbo[nextidx])
	//gl.BufferData(gl.PIXEL_UNPACK_BUFFER, len(img.Pix), gl.PtrOffset(0), gl.STREAM_DRAW)
	mapbuffer := gl.MapBuffer(gl.PIXEL_UNPACK_BUFFER, gl.WRITE_ONLY)
	if mapbuffer != nil {
		C.memcpy(mapbuffer, unsafe.Pointer(&img.Pix[0]), C.size_t(len(img.Pix)))
		if !gl.UnmapBuffer(gl.PIXEL_UNPACK_BUFFER) {
			panic("unmap error")
		}
	} else {
		panic("mem nil")
	}
	gl.BindBuffer(gl.PIXEL_UNPACK_BUFFER, 0)

}
func (s *GLRender) Draw() {
	gl.UseProgram(s.program)
	gl.BindVertexArray(s.vao)
	gl.BindTexture(gl.TEXTURE_2D, s.frameimage)
	gl.DrawArrays(gl.TRIANGLES, 0, 6*2*3)
}
