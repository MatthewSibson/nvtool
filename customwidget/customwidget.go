package customwidget

import (
	g "github.com/AllenDang/giu"
	"github.com/AllenDang/giu/imgui"
	theme "github.com/Nicify/nvtool/theme"
)

func InputTextVHiPDIFont(label string, width float32, value *string, flags g.InputTextFlags, font imgui.Font, cb imgui.InputTextCallback, onChange func()) g.Layout {
	useFont := theme.UseFont(font)
	return g.Layout{
		g.Custom(useFont.Push),
		g.InputTextV(label, width, value, flags, cb, onChange),
		g.Custom(useFont.Pop),
	}
}

func InputTextMultilineHiDPIFont(label string, text *string, width float32, height float32, flags g.InputTextFlags, font imgui.Font, cb imgui.InputTextCallback, onChange func()) g.Layout {
	useFont := theme.UseFont(font)
	return g.Layout{
		g.Custom(useFont.Push),
		g.InputTextMultiline(label, text, width, height, flags, cb, onChange),
		g.Custom(useFont.Pop),
	}
}

func WithHiDPIFont(hDPIFont imgui.Font, lDPIFont imgui.Font, layout g.Layout) g.Layout {
	return g.Layout{
		g.Custom(func() {
			if imgui.DPIScale == 1 {
				g.PushFont(lDPIFont)
			} else {
				g.PushFont(hDPIFont)
			}
		}),
		layout,
		g.Custom(func() {
			g.PopFont()
		}),
	}
}