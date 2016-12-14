package sketches

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/raedatoui/learn-opengl/utils"
)

type HelloShaders struct {
	Window   *glfw.Window
	Vao, Vbo uint32
	Shader  uint32
}


func (sketch *HelloShaders) Setup() {
	var err error
	sketch.Shader, err = utils.Shader(
		"sketches/assets/3.shaders/basic.vs",
		"sketches/assets/3.shaders/basic.frag", "")

	if err != nil {
		panic(err)
	}

	var vertices = []float32{
		// Positions      // Colors
		0.5, -0.5, 0.0,   1.0, 0.0, 0.0, // Bottom Right
		-0.5, -0.5, 0.0,  0.0, 1.0, 0.0, // Bottom Left
		0.0, 0.5, 0.0,    0.0, 0.0, 1.0, // Top
	}
	gl.GenVertexArrays(1, &sketch.Vao)
	gl.GenBuffers(1, &sketch.Vbo)

	gl.BindVertexArray(sketch.Vao)

	gl.BindBuffer(gl.ARRAY_BUFFER, sketch.Vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices) * utils.GL_FLOAT32_SIZE, gl.Ptr(vertices), gl.STATIC_DRAW)

	// position uniform
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 6 * utils.GL_FLOAT32_SIZE, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)

	//color uniform
	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, 6 * utils.GL_FLOAT32_SIZE, gl.PtrOffset(3 * utils.GL_FLOAT32_SIZE))
	gl.EnableVertexAttribArray(1)

	gl.BindVertexArray(0)

	gl.Disable(gl.DEPTH_TEST)
}

func (sketch *HelloShaders) Update() {

}

func (sketch *HelloShaders) Draw() {
	gl.ClearColor(0.2, 0.3, 0.3, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT)

	// Draw the triangle
	gl.UseProgram(sketch.Shader)
	gl.BindVertexArray(sketch.Vao)
	gl.DrawArrays(gl.TRIANGLES, 0, 3)
	gl.BindVertexArray(0)
}

func (sketch *HelloShaders) Close() {
	gl.DeleteVertexArrays(1, &sketch.Vao)
	gl.DeleteBuffers(1, &sketch.Vbo)
	gl.UseProgram(0)
}

func (sketch *HelloShaders) HandleKeyboard(key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if key == glfw.KeyEscape && action == glfw.Press {
		sketch.Window.SetShouldClose(true)
	}
}