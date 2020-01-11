package ui

import (
	"encoding/binary"
	"fmt"

	"github.com/inkyblackness/imgui-go/v2"
	"github.com/lucasb-eyer/go-colorful"
)

var pallete []imgui.Vec4
var hoverIdx = -1;
//var pallete = colorful.FastWarmPalette(256)
func init() {
	p := colorful.FastWarmPalette(256)
	p[0]=colorful.Color{R:0.4, G:0.4, B: 0.4}
	p[255]=colorful.Color{R:0.8, G:0.8, B: 0.8}
	p=append(p,colorful.Color{R:0.6, G:1, B: 0.6})
	p=append(p,colorful.Color{R:0.6, G:1, B: 0.6})
	p=append(p,colorful.Color{R:1, G:0.6, B: 1})
	p=append(p,colorful.Color{R:1, G:1, B: 1})
	pallete = make([]imgui.Vec4,len(p))

	for idx, c := range p {
		pallete[idx].X = float32(c.R)
		pallete[idx].Y = float32(c.G)
		pallete[idx].Z = float32(c.B)
		pallete[idx].W = 1
	}
}

func Hexview(buffer *[]byte, columns int) {
	imgui.Text(fmt.Sprintf("%d",hoverIdx))

	imgui.ColumnsV(columns+1,"hex",true)
	size := imgui.CalcTextSize(fmt.Sprintf(" %04X ",len(*buffer)),false,128)
	imgui.SetColumnWidth(0,size.X)
	imgui.NextColumn()
	imgui.PushStyleColor(imgui.StyleColorText,imgui.Vec4{
		X: 0.8,
		Y: 0.6,
		Z: 0.4,
		W: 1,
	})
	for i := 0; i < columns ; i++ {
		imgui.Text(fmt.Sprintf("%02X",i))
		size := imgui.CalcTextSize(" 00 ",false,0)
		imgui.SetColumnWidth(i+1,size.X)
		imgui.NextColumn()
	}
	imgui.PopStyleColor()
	//imgui.Separator()
	resetHover := true
	for idx, b := range *buffer{
		if imgui.ColumnIndex() == 0 {
			imgui.PushStyleColor(imgui.StyleColorText,imgui.Vec4{
				X: 0.6,
				Y: 0.8,
				Z: 0.4,
				W: 1,
			})
			imgui.Text(fmt.Sprintf("%04X", idx))

			imgui.NextColumn()
			imgui.PopStyleColor()
		}
		blockStart := hoverIdx - 3
		if blockStart < 0 {blockStart=0}
		blockEnd := hoverIdx

		if  idx >= blockStart && idx <= blockEnd  {
			imgui.PushStyleColor(imgui.StyleColorText, pallete[256 + idx-blockStart])
		} else {
			imgui.PushStyleColor(imgui.StyleColorText, pallete[int(b)])
		}
		imgui.Text(fmt.Sprintf("%02X",b))
		if imgui.IsItemHovered() {
			hoverIdx = idx
			resetHover = false
		}
		if  idx == blockEnd {
			sl := (*buffer)[blockStart:blockEnd+1]
			hexTooltip(sl)

		}
		imgui.PopStyleColor()
		imgui.NextColumn()
	}
	if resetHover { hoverIdx = -1}
}

func hexTooltip(b []byte) {
	imgui.BeginTooltip()

	defer imgui.EndTooltip()
	imgui.Text(fmt.Sprintf("l: %d",len(b)))
	imgui.PushStyleColor(imgui.StyleColorText,pallete[255+4])
	imgui.Text("byte:")
	imgui.SameLine()
	imgui.Text(fmt.Sprintf("%d",int(b[len(b)-1])))
	imgui.SameLine()
	imgui.Text(fmt.Sprintf("`%q`",string(b[len(b)-1])))
	imgui.SameLine()
	imgui.Text(fmt.Sprintf("%08b",b[len(b)-1]))
	imgui.PopStyleColor()
	if len(b) >= 2 {
		imgui.PushStyleColor(imgui.StyleColorText,pallete[255+3])

		w := b[len(b)-2:]
		imgui.Text("word:")
		imgui.SameLine()
		imgui.Text(fmt.Sprintf("%d", binary.BigEndian.Uint16(w)))
		imgui.SameLine()
		imgui.Text(fmt.Sprintf("`%q`", string(w)))
		imgui.SameLine()
		imgui.Text(fmt.Sprintf("%08b", w))
			imgui.PopStyleColor()

	}

	if len(b) >= 4 {
		dw := b[len(b)-4:]
		imgui.PushStyleColor(imgui.StyleColorText,pallete[255+2])

		imgui.Text("dword")
		imgui.SameLine()
		imgui.Text(fmt.Sprintf("%d", binary.BigEndian.Uint32(dw)))
		imgui.SameLine()
		imgui.Text(fmt.Sprintf("'%q'", string(dw)))
		imgui.SameLine()
		imgui.Text(fmt.Sprintf("%08b", dw))
		imgui.SameLine()
		imgui.PopStyleColor()

	}

}