package ui

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/inkyblackness/imgui-go/v2"

	"github.com/XANi/goz80/core/z80"
)

// Platform covers mouse/keyboard/gamepad inputs, cursor shape, timing, windowing.
type Platform interface {
	// ShouldStop is regularly called as the abort condition for the program loop.
	ShouldStop() bool
	// ProcessEvents is called once per render loop to dispatch any pending events.
	ProcessEvents()
	// DisplaySize returns the dimension of the display.
	DisplaySize() [2]float32
	// FramebufferSize returns the dimension of the framebuffer.
	FramebufferSize() [2]float32
	// NewFrame marks the begin of a render pass. It must update the imgui IO state according to user input (mouse, keyboard, ...)
	NewFrame()
	// PostRender marks the completion of one render pass. Typically this causes the display buffer to be swapped.
	PostRender()
	// ClipboardText returns the current text of the clipboard, if available.
	ClipboardText() (string, error)
	// SetClipboardText sets the text as the current text of the clipboard.
	SetClipboardText(text string)
}

type clipboard struct {
	platform Platform
}

func (board clipboard) Text() (string, error) {
	return board.platform.ClipboardText()
}

func (board clipboard) SetText(text string) {
	board.platform.SetClipboardText(text)
}

// Renderer covers rendering imgui draw data.
type Renderer interface {
	// PreRender causes the display buffer to be prepared for new output.
	PreRender(clearColor [4]float32)
	// Render draws the provided imgui draw data.
	Render(displaySize [2]float32, framebufferSize [2]float32, drawData imgui.DrawData)
}



var mt time.Time
var mtIntervals = make(map[string]time.Duration,0)
var mtIntervalsOrder = make([]string,0)

func measureInterval(name string) {
	if _, ok := mtIntervals[name]; !ok {
		mtIntervalsOrder = append(mtIntervalsOrder,name)
	}

	mtIntervals[name]=time.Since(mt)
	mt = time.Now()

}

// Run implements the main program loop of the demo. It returns when the platform signals to stop.
// This demo application shows some basic features of ImGui, as well as exposing the standard demo window.
func Run(p Platform, r Renderer) {
	imgui.CurrentIO().SetClipboard(clipboard{platform: p})
	buffer := make([]byte,65535)
	rand.Read(buffer)

	showDemoWindow := false
	clearColor := [4]float32{0.0, 0.0, 0.0, 1.0}
	showPerfbox := false
	runZ80 := false
	stepsPerCycle:=int32(1)
	start := int32(0)
	cpu := z80.InitCPU()
	prog := []byte{
		0x0e,0x0b,
		0x06,0x01,
		0x3c,
		0x3c,
		0x0c,
		0x0c,
		0x0c,
		0x02,
		0xa9,
		0xC3,0x00,0x04}

	copy(cpu.Data[:],prog)
	measureInterval("start")
	frame := 0
	for !p.ShouldStop() {
		measureInterval("init")
		frame++
		s := time.Now()

		p.ProcessEvents()
		perfboxFrameStart()
		// Signal start of a new frame
		p.NewFrame()
		imgui.NewFrame()
		measureInterval("frame-start")
		// 1. Show the big demo window (Most of the sample code is in ImGui::ShowDemoWindow()!
		// You can browse its code to learn more about Dear ImGui!).
		if showDemoWindow {
			imgui.ShowDemoWindow(&showDemoWindow)
		}
		measureInterval("demo-window")
		// 2. Show a simple window that we create ourselves. We use a Begin/End pair to created a named window.
		{
			imgui.Begin("Hello, world!") // Create a window called "Hello, world!" and append into it.

			imgui.Text("This is some useful text.") // Display some text

			imgui.Checkbox("Demo Window", &showDemoWindow) // Edit bools storing our window open/close state
			imgui.Checkbox("Perfbox", &showPerfbox)
			imgui.Checkbox("Run Z80", &runZ80)
			imgui.SliderIntV("steps/frame", &stepsPerCycle,1,655350,"%d")


			imgui.SliderIntV("start", &start, 0, int32(len(buffer)),"%d")
			Hexview(&cpu.Data,16,int(start),512)

			imgui.End()
		}
		measureInterval("Hexview")
		if runZ80 {
				for i :=int32(0) ; i <=stepsPerCycle; i++ {
					cpu.Step()
				}
		}
		measureInterval("z80")


		if showPerfbox {
			Perfbox(&showPerfbox)
		}
		measureInterval("perfbox")
		// Rendering
		imgui.Render() // This call only creates the draw data list. Actual rendering to framebuffer is done below.
		measureInterval("imgui-render")
		perfboxFrameStop()

		r.PreRender(clearColor)
		measureInterval("pre-render")
		// A this point, the application could perform its own rendering...
		// app.RenderScene()

		r.Render(p.DisplaySize(), p.FramebufferSize(), imgui.RenderedDrawData())
		measureInterval("render")

		p.PostRender()
		measureInterval("post-render")
		ft := time.Since(s)

		if ft > time.Millisecond * 20 {
			fmt.Printf("frame time exceeded: %s\n", ft)
		} else if ft < time.Millisecond * 10 {
			//// sleep to avoid 100% CPU usage if frame is fast
			<-time.After((time.Millisecond * 16) - ft)
		}
		if (frame % 60 ) == 30 {
			fmt.Println("-- timing --\n")
			for _, step := range mtIntervalsOrder[1:] {
				fmt.Printf("%s:\t%s\n",step, mtIntervals[step])
			}
		}
	}
}

func ByteCountBinary(b uint64) string {
        const unit = 1024
        if b < unit {
                return fmt.Sprintf("%d B", b)
        }
        div, exp := int64(unit), 0
        for n := b / unit; n >= unit; n /= unit {
                div *= unit
                exp++
        }
        return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "KMGTPE"[exp])
}
func Count(b uint64) string {
        const unit = 1000
        if b < unit {
                return fmt.Sprintf("%d B", b)
        }
        div, exp := int64(unit), 0
        for n := b / unit; n >= unit; n /= unit {
                div *= unit
                exp++
        }
        return fmt.Sprintf("%.1f %c", float64(b)/float64(div), "KMGTPE"[exp])
}