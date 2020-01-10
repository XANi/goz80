package demo

import (
	"fmt"
	"math"
	"runtime"
	"runtime/debug"
	"strconv"
	"time"

	"github.com/inkyblackness/imgui-go/v2"
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

// Run implements the main program loop of the demo. It returns when the platform signals to stop.
// This demo application shows some basic features of ImGui, as well as exposing the standard demo window.
func Run(p Platform, r Renderer) {
	imgui.CurrentIO().SetClipboard(clipboard{platform: p})

	showDemoWindow := false
	clearColor := [4]float32{0.0, 0.0, 0.0, 1.0}
	f := float32(0.0)
	counter := 0
	showAnotherWindow := false
	var frameTime time.Duration
	var frameGraphProbes = 256
	var frameTimes = make([]float32,frameGraphProbes)
	var frameHist = make([]float32,32)
	var runtimeGC = make([]float32,frameGraphProbes)
	var heapHistory = make([]float32,frameGraphProbes)
	var pauseHistory = make([]float32,256)
	var frames int
	var GCStats debug.GCStats
	var MemStats runtime.MemStats
	gcp := int32(debug.SetGCPercent(100))
	debug.SetGCPercent(int(gcp))
	debug.ReadGCStats(&GCStats)
	runtime.ReadMemStats(&MemStats)
	prevGCCount := MemStats.NumGC
	for !p.ShouldStop() {
		frames++
		sum := float32(0)
		max := float32(0)
		min := frameTimes[0]
		for _, i := range (frameTimes) {
			sum +=  i
			if max < i { max = i }
			if min > i { min = i }
		}
		for i := range(frameHist) {
			frameHist[i] *= 0.99
		}
		avg :=  sum/float32(len(frameTimes))

		if len(frameHist) > 100 && (frameHist[len(frameHist)-1]) < 0.5 {
			frameHist = frameHist[:len(frameHist) - 2]
		}

		if len(GCStats.Pause)  > len(pauseHistory) {
			pauseHistory = make([]float32,len(GCStats.Pause))
		}
		var pauseMax float32
		for idx, p := range MemStats.PauseNs {
			us := float32(p/1000)
			pauseHistory[idx] = us
			if pauseMax < us {
				pauseMax=us
			}
		}
		start := time.Now()
		p.ProcessEvents()

		// Signal start of a new frame
		p.NewFrame()
		imgui.NewFrame()

		// 1. Show the big demo window (Most of the sample code is in ImGui::ShowDemoWindow()!
		// You can browse its code to learn more about Dear ImGui!).
		if showDemoWindow {
			imgui.ShowDemoWindow(&showDemoWindow)
		}

		// 2. Show a simple window that we create ourselves. We use a Begin/End pair to created a named window.
		{
			imgui.Begin("Hello, world!") // Create a window called "Hello, world!" and append into it.

			imgui.Text("This is some useful text.") // Display some text

			imgui.Checkbox("Demo Window", &showDemoWindow) // Edit bools storing our window open/close state
			imgui.Checkbox("Another Window", &showAnotherWindow)

			imgui.SliderFloat("float", &f, 0.0, 1.0) // Edit one float using a slider from 0.0f to 1.0f
			// TODO add example of ColorEdit3 for clearColor

			if imgui.Button("Button") { // Buttons return true when clicked (most widgets return true when edited/activated)
				counter++
			}
			imgui.SameLine()
			imgui.Text(fmt.Sprintf("counter = %d", counter))

			// TODO add text of FPS based on IO.Framerate()

			imgui.End()
		}
		{
			imgui.Begin("Perfbox")

			imgui.Text(fmt.Sprintf("frame time avg/min/max %0.2f/%0.2f/%0.0f us", avg,min,max))
			imgui.PlotLinesV("frameTime",frameTimes,0,"",0,math.MaxFloat32,imgui.Vec2{})
			imgui.PlotHistogramV("frameTime/10",frameHist,10,"",0,math.MaxFloat32,imgui.Vec2{})
			imgui.PlotLines("GC",runtimeGC)
			if len(GCStats.Pause) >0  {
				imgui.PlotLines(
					"PauseTime, last: " +
						strconv.Itoa(int(MemStats.PauseNs[(MemStats.NumGC+255)%256]/1000)) +
						"us, max: " +
						strconv.Itoa(int(pauseMax))+ "us",
						pauseHistory)
			} else {
				imgui.PlotLines("PauseTime" , pauseHistory)
			}
			imgui.DragIntV("GC percent",&gcp,1,10,1000,"%d %")
			imgui.ColumnsV(2,"meminfo",true)
			imgui.Text("HeapAlloc")
			imgui.Text("HeapSys")
			imgui.Text("HeapReleased")
			imgui.Text("HeapIdle")
			imgui.Text("HeapInUse")
			imgui.Text("StackSys")
			imgui.Text("StackInuse")
			imgui.Text("OtherSys")
			imgui.Text("HeapObjects")
			imgui.NextColumn()
			imgui.Text(ByteCountBinary(MemStats.HeapAlloc))
			imgui.Text(ByteCountBinary(MemStats.HeapSys))
			imgui.Text(ByteCountBinary(MemStats.HeapReleased))
			imgui.Text(ByteCountBinary(MemStats.HeapIdle))
			imgui.Text(ByteCountBinary(MemStats.HeapInuse))
			imgui.Text(ByteCountBinary(MemStats.StackSys))
			imgui.Text(ByteCountBinary(MemStats.StackInuse))
			imgui.Text(ByteCountBinary(MemStats.OtherSys))
			imgui.Text(Count(MemStats.HeapObjects))

		imgui.End()
		}
		// 3. Show another simple window.
		if showAnotherWindow {
			// Pass a pointer to our bool variable (the window will have a closing button that will clear the bool when clicked)
			imgui.BeginV("Another window", &showAnotherWindow, 0)

			imgui.Text("Hello from another window!")
			if imgui.Button("Close Me") {
				showAnotherWindow = false
			}
			imgui.End()
		}

		// Rendering
		imgui.Render() // This call only creates the draw data list. Actual rendering to framebuffer is done below.

		r.PreRender(clearColor)
		// A this point, the application could perform its own rendering...
		// app.RenderScene()

		r.Render(p.DisplaySize(), p.FramebufferSize(), imgui.RenderedDrawData())
		p.PostRender()
		frameTime = time.Since(start)
		debug.SetGCPercent(int(gcp))
		debug.ReadGCStats(&GCStats)
		runtime.ReadMemStats(&MemStats)
		us :=  float32(frameTime.Nanoseconds()/1000)
		frameTimes[frames % frameGraphProbes] = us
		runtimeGC[frames % frameGraphProbes] = float32(MemStats.NumGC - prevGCCount)
		heapHistory[frames % frameGraphProbes] = float32(MemStats.HeapInuse)
		idx := int(us/10)
		if len(frameHist) <= idx {
			for i := len(frameHist)-1 ; i <= idx; i++ {
				frameHist = append(frameHist,0)
			}
		}
		frameHist[idx]++

		_=frameHist
		prevGCCount=MemStats.NumGC

		// sleep to avoid 100% CPU usage for this demo
		<-time.After(time.Millisecond * 25)
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