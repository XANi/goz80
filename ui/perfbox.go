package ui

import (
	"fmt"
	"math"
	"runtime"
	"runtime/debug"
	"strconv"
	"time"

	"github.com/inkyblackness/imgui-go/v2"
)
var frameGraphProbes = 256
type runtimeStats struct {
	frameStart time.Time
	frameTime time.Duration
	frameTimes []float32
	frameHist []float32
	frameMin, frameSum, frameMax, frameAvg  float32
	runtimeGC []float32
	heapHistory []float32
	frames int
	GCStats debug.GCStats
	pauseHistory []float32
	pauseMax float32
	prevGCCount uint32
}
var MemStats runtime.MemStats
var stats = runtimeStats{
	frameTimes:   make([]float32, frameGraphProbes),
	frameHist:    make([]float32, 32),
	runtimeGC:    make([]float32, frameGraphProbes),
	heapHistory:  make([]float32, frameGraphProbes),
	pauseHistory: make([]float32, len(MemStats.PauseNs)),
}
func init() {
	runtime.ReadMemStats(&MemStats)
	stats.prevGCCount =MemStats.NumGC
}
func perfboxFrameStart() {
	stats.frameStart = time.Now()
}
func perfboxFrameStop() {
	stats.frameTime = time.Since(stats.frameStart)
	perfboxUpdateGC()
}



func perfboxUpdateGC() {
	runtime.ReadMemStats(&MemStats)
	us :=  float32(stats.frameTime.Nanoseconds()/1000)
	stats.frameTimes[stats.frames % frameGraphProbes] = us
	stats.runtimeGC[stats.frames % frameGraphProbes] = float32(MemStats.NumGC - stats.prevGCCount)
	stats.heapHistory[stats.frames % frameGraphProbes] = float32(MemStats.HeapInuse)
	idx := int(us/10)
	if len(stats.frameHist) <= idx {
		for i := len(stats.frameHist)-1 ; i <= idx; i++ {
			stats.frameHist = append(stats.frameHist,0)
		}
	}
	stats.frameHist[idx]++
	stats.pauseMax = 0
	for idx, p := range MemStats.PauseNs {
		us := float32(p/1000)
		stats.pauseHistory[idx] = us
		if stats.pauseMax < us {
			stats.pauseMax=us
		}
	}
	stats.frameSum = float32(0)
	stats.frameMin = stats.frameTimes[0]
	for _, i := range (stats.frameTimes) {
		stats.frameSum +=  i
		if stats.frameMax < i { stats.frameMax = i }
		if stats.frameMin > i { stats.frameMin = i }
	}
	stats.frameAvg =  stats.frameSum/float32(len(stats.frameTimes))
	stats.frames++
	stats.prevGCCount=MemStats.NumGC
}

func Perfbox(open *bool) {
	imgui.BeginV("Perfbox",open,0)
	imgui.Text(fmt.Sprintf("frame time avg/min/max %0.2f/%0.2f/%0.0f us",
		stats.frameAvg,
		stats.frameMin,
		stats.frameMax))
	imgui.PlotLinesV("frameTime",stats.frameTimes,
		0,"",
		0,math.MaxFloat32,imgui.Vec2{})
	imgui.PlotHistogramV("frameTime/10",stats.frameHist,
		10,"",
		0,math.MaxFloat32,imgui.Vec2{})
	imgui.PlotLines("GC",stats.runtimeGC)
	imgui.PlotLines("PauseTime" , stats.pauseHistory)
	if MemStats.PauseNs[0] > 0  {
		imgui.Text(
			"last: " +
				strconv.Itoa(int(MemStats.PauseNs[(MemStats.NumGC+255)%256]/1000)) +
				"us, max: " +
				strconv.Itoa(int(stats.pauseMax))+ "us")
	} else {
		imgui.Text("No pause yet")
	}
	imgui.PushStyleColor(imgui.StyleColorText,imgui.Vec4{
		X: 0.6,
		Y: 0.6,
		Z: 0.8,
		W: 1,
	})
	imgui.ColumnsV(2,"meminfo",true)
	imgui.Separator()

	imgui.Text("HeapAlloc")
	imgui.NextColumn()
	imgui.Text(ByteCountBinary(MemStats.HeapAlloc))
	imgui.Separator()
	imgui.NextColumn()

	imgui.Text("HeapSys")
	imgui.NextColumn()
	imgui.Text(ByteCountBinary(MemStats.HeapSys))
	imgui.Separator()
	imgui.NextColumn()


	imgui.Text("HeapReleased")
	imgui.NextColumn()
	imgui.Text(ByteCountBinary(MemStats.HeapReleased))
	imgui.Separator()
	imgui.NextColumn()

	imgui.Text("HeapIdle")
	imgui.NextColumn()
	imgui.Text(ByteCountBinary(MemStats.HeapIdle))
	imgui.Separator()
	imgui.NextColumn()

	imgui.Text("HeapInUse")
	imgui.NextColumn()
	imgui.Text(ByteCountBinary(MemStats.HeapInuse))
	imgui.Separator()
	imgui.NextColumn()

	imgui.Text("StackSys")
	imgui.NextColumn()
	imgui.Text(ByteCountBinary(MemStats.StackSys))
	imgui.Separator()
	imgui.NextColumn()

	imgui.Text("StackInuse")
	imgui.NextColumn()
	imgui.Text(ByteCountBinary(MemStats.StackInuse))
	imgui.Separator()
	imgui.NextColumn()

	imgui.Text("OtherSys")
	imgui.NextColumn()
	imgui.Text(ByteCountBinary(MemStats.OtherSys))
	imgui.Separator()
	imgui.NextColumn()

	imgui.Text("HeapObjects")
	imgui.NextColumn()
	imgui.Text(Count(MemStats.HeapObjects))
	imgui.Separator()
	imgui.NextColumn()


	imgui.PopStyleColor()

	imgui.End()
}