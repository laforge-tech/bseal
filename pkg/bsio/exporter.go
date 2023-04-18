package bsio

import(
	"io"

	"github.com/pterm/pterm"
)

type Exporter struct {
	configuration Configuration
}

func NewExporter(configuration Configuration) * Exporter {
	return &Exporter{
		configuration: configuration,
	}
}

func (e *Exporter) Export(localFile io.Reader, archiveName string) error {

	pterm.Info.Prefix.Text =    "INFO    "
	pterm.Warning.Prefix.Text = "WARNING "
	pterm.Error.Prefix.Text =   "ERROR   "

	pterm.Info.Printf("Working\n")
	pterm.Warning.Printf("Working\n")
	pterm.Error.Printf("Working\n")
	return nil
}
