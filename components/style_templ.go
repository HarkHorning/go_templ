// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.857
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

// css
func flex() templ.CSSClass {
	templ_7745c5c3_CSSBuilder := templruntime.GetBuilder()
	templ_7745c5c3_CSSBuilder.WriteString(`display:flex;`)
	templ_7745c5c3_CSSBuilder.WriteString(`justify-content:center;`)
	templ_7745c5c3_CSSID := templ.CSSID(`flex`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func col() templ.CSSClass {
	templ_7745c5c3_CSSBuilder := templruntime.GetBuilder()
	templ_7745c5c3_CSSBuilder.WriteString(`flex-direction:column;`)
	templ_7745c5c3_CSSID := templ.CSSID(`col`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func sleeve() templ.CSSClass {
	templ_7745c5c3_CSSBuilder := templruntime.GetBuilder()
	templ_7745c5c3_CSSBuilder.WriteString(`width:1200px;`)
	templ_7745c5c3_CSSID := templ.CSSID(`sleeve`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func textColor() templ.CSSClass {
	templ_7745c5c3_CSSBuilder := templruntime.GetBuilder()
	templ_7745c5c3_CSSBuilder.WriteString(`color:white;`)
	templ_7745c5c3_CSSID := templ.CSSID(`textColor`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func spacing() templ.CSSClass {
	templ_7745c5c3_CSSBuilder := templruntime.GetBuilder()
	templ_7745c5c3_CSSBuilder.WriteString(`gap:1rem;`)
	templ_7745c5c3_CSSBuilder.WriteString(`justify-content:space-between;`)
	templ_7745c5c3_CSSBuilder.WriteString(`align-items:center;`)
	templ_7745c5c3_CSSID := templ.CSSID(`spacing`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func footerbox() templ.CSSClass {
	templ_7745c5c3_CSSBuilder := templruntime.GetBuilder()
	templ_7745c5c3_CSSBuilder.WriteString(`padding:1rem 2rem;`)
	templ_7745c5c3_CSSBuilder.WriteString(`border-radius:12px;`)
	templ_7745c5c3_CSSBuilder.WriteString(`background-color:#FFC000;`)
	templ_7745c5c3_CSSBuilder.WriteString(`color:black;`)
	templ_7745c5c3_CSSID := templ.CSSID(`footerbox`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

var _ = templruntime.GeneratedTemplate
