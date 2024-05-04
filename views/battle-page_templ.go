// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.663
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"Battleships/client"
)

func MakeBattlePage(token string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<style>\r\n    body {\r\n        margin: 0;\r\n        height: 100vh;\r\n        display: flex;\r\n        flex-direction: column;\r\n        align-items: center;\r\n        justify-content: center;\r\n        background-color: #141526;\r\n        color: white;\r\n        font-family: Arial, sans-serif;\r\n    }\r\n    .footer {\r\n        background-color: black;\r\n        color: white;\r\n        padding: 10px;\r\n        position: fixed;\r\n        left: 0;\r\n        bottom: 0;\r\n        width: 100%;\r\n        text-align: center;\r\n        display: flex;\r\n        justify-content: center;\r\n        align-items: center;\r\n    }\r\n    .input-field {\r\n        margin-right: 10px;\r\n        padding: 5px;\r\n        font-size: 16px;\r\n        width: 60px; /* Set width to accommodate 3 characters */\r\n    }\r\n    .submit-button {\r\n        padding: 5px 15px;\r\n        font-size: 16px;\r\n        background-color: white;\r\n        color: black;\r\n        border: none;\r\n        cursor: pointer;\r\n    }\r\n    .boards {\r\n        display: flex;\r\n        justify-content: space-around;\r\n        align-items: center;\r\n        flex-grow: 1;\r\n        width: 100%;\r\n        margin-bottom: 50px;\r\n    }\r\n    .board {\r\n        background-color: #837777;\r\n        width: 40%;\r\n        aspect-ratio: 1;\r\n        display: grid;\r\n        grid-template-columns: repeat(10, 1fr);\r\n        grid-template-rows: repeat(10, 1fr);\r\n        gap: 3px;\r\n        margin: 12px;\r\n        border: 3px solid #837777;\r\n    }\r\n</style><!doctype html><html lang=\"en\"><head><title>Battle Page Temple</title></head><body><div class=\"boards\"><div class=\"board\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for col := 0; col < 10; col++ {
			for row := 0; row < 10; row++ {
				templ_7745c5c3_Err = MakeBoardCell(client.CalculateCellCoord(row, col)).Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div class=\"board\"></div></div><div class=\"footer\"><input type=\"text\" maxlength=\"3\" class=\"input-field\" placeholder=\"...\"> <button class=\"submit-button\">Fire</button></div></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func MakeBoardCell(coord string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<style>\r\n    .cell {\r\n        background-color: #222;\r\n        display: flex; /* Makes cell a flex container */\r\n        justify-content: center; /* Centers content horizontally */\r\n        align-items: center; /* Centers content vertically */\r\n        color: white; /* Sets the text color */\r\n        font-size: 24px; /* Sets the text size */\r\n    }\r\n</style><!doctype html><html lang=\"en\"><body><div class=\"cell\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(coord)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views\battle-page.templ`, Line: 109, Col: 29}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
