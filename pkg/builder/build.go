package builder

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"os"
	"path"
	"path/filepath"
	"text/template"

	"gopkg.in/yaml.v3"
)

var (
	defaultHTMLTemplate = template.Must(template.New(path.Base("index.gotmpl")).ParseFiles("index.gotmpl"))
)

func Handler(tpl string, input *string, output string) error {
	var htmlTemplate *template.Template

	cwd, err := os.Getwd()
	if err != nil {
		log.Warn().Msg(err.Error())
		return err
	}

	if _, err := os.Stat(tpl); os.IsNotExist(err) {
		log.Info().Msg("Template file was not found, using the default template")
		htmlTemplate = defaultHTMLTemplate
	} else {
		templatePath := filepath.Join(cwd, "index.gotmpl")
		htmlTemplate = template.Must(template.New(path.Base(templatePath)).ParseFiles(templatePath))
	}

	data, err := os.ReadFile(*input)

	if err != nil {
		log.Warn().Msg(err.Error())
		return ErrFailReadChartConfig
	}

	charts := Charts{}

	err = yaml.Unmarshal(data, &charts)
	if err != nil {
		log.Warn().Msg(err.Error())
		return ErrFailProcessChartFile
	}

	err = Render(output, htmlTemplate, &charts)

	if err != nil {
		log.Warn().Msg(err.Error())
		return ErrFailRenderTemplate
	}

	return err
}

func Render(output string, htmlTemplate *template.Template, charts *Charts) error {
	var outputHandle *os.File
	var err error

	if output == "-" {
		outputHandle = os.Stdout
	} else {
		outputHandle, err = os.Create(output)
		if err != nil {
			return err
		}
		log.Info().Msg(fmt.Sprintf("Rendering the template %s", output))
	}

	err = htmlTemplate.Execute(outputHandle, charts)

	if err != nil {
		return err
	}

	return err
}
