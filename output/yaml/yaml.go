package yaml

import (
	"io"

	"github.com/goccy/go-yaml"
	"github.com/k1LoW/tbls/schema"
)

// YAML struct
type YAML struct{}

// OutputSchema output YAML format for full relation.
func (j *YAML) OutputSchema(wr io.Writer, s *schema.Schema) error {
	encoder := yaml.NewEncoder(wr)
	err := encoder.Encode(s)
	if err != nil {
		return err
	}
	return nil
}

// OutputTable output YAML format for table.
func (j *YAML) OutputTable(wr io.Writer, t *schema.Table) error {
	encoder := yaml.NewEncoder(wr)
	err := encoder.Encode(t)
	if err != nil {
		return err
	}
	return nil
}

func (j *YAML) OutputProcedure(wr io.Writer, s *schema.Procedure) error {
	/* TODO:
	ここで、プロシージャの出力のための結果を取得する。
	その結果をたぶん、tmpl.Execute(wr, templateData)する。
	*/

	return nil
}
