// Code generated by templ@(devel) DO NOT EDIT.

package main

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "strconv"

func counts(global, user int) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_1 := templ.GetChildren(ctx)
		if var_1 == nil {
			var_1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<div>")
		if err != nil {
			return err
		}
		var_2 := `Global: `
		_, err = templBuffer.WriteString(var_2)
		if err != nil {
			return err
		}
		var var_3 string = strconv.Itoa(global)
		_, err = templBuffer.WriteString(templ.EscapeString(var_3))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</div><div>")
		if err != nil {
			return err
		}
		var_4 := `User: `
		_, err = templBuffer.WriteString(var_4)
		if err != nil {
			return err
		}
		var var_5 string = strconv.Itoa(user)
		_, err = templBuffer.WriteString(templ.EscapeString(var_5))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</div>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}

func form() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_6 := templ.GetChildren(ctx)
		if var_6 == nil {
			var_6 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<form action=\"/\" method=\"POST\"><div><button type=\"submit\" name=\"global\" value=\"global\">")
		if err != nil {
			return err
		}
		var_7 := `Global`
		_, err = templBuffer.WriteString(var_7)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</button></div><div><button type=\"submit\" name=\"user\" value=\"user\">")
		if err != nil {
			return err
		}
		var_8 := `User`
		_, err = templBuffer.WriteString(var_8)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</button></div></form>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}

func page(global, user int) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_9 := templ.GetChildren(ctx)
		if var_9 == nil {
			var_9 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<html><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>")
		if err != nil {
			return err
		}
		var_10 := `Counts`
		_, err = templBuffer.WriteString(var_10)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</title><link rel=\"stylesheet\" href=\"/assets/bulma.min.css\"><link rel=\"apple-touch-icon\" sizes=\"180x180\" href=\"/assets/favicon/apple-touch-icon.png\"><link rel=\"icon\" type=\"image/png\" sizes=\"32x32\" href=\"/assets/favicon/favicon-32x32.png\"><link rel=\"icon\" type=\"image/png\" sizes=\"16x16\" href=\"/assets/favicon/favicon-16x16.png\"><link rel=\"manifest\" href=\"/assets/favicon/site.webmanifest\"></head><body class=\"bg-gray-100\"><header class=\"hero is-primary\"><div class=\"hero-body\"><div class=\"container\"><h1 class=\"title\">")
		if err != nil {
			return err
		}
		var_11 := `Counts`
		_, err = templBuffer.WriteString(var_11)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</h1></div></div></header><section class=\"section\"><div class=\"container\"><div class=\"columns is-centered\"><div class=\"column is-half\">")
		if err != nil {
			return err
		}
		err = counts(global, user).Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		err = form().Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</div></div></div></section></body></html>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}