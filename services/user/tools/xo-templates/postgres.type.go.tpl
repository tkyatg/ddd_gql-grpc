{{- $short := (shortname .Name "err" "res" "sqlstr" "db" "XOLog") -}}
{{- $table := (.Table.TableName) -}}

// {{ .Name }}InsertArgs represents a row from '{{ $table }}'.
type {{ .Name }}InsertArgs struct {
{{- range .Fields }}{{ if ne .Name "ID" }}
	{{ .Name }} {{ retype .Type }} `db:"{{ .Col.ColumnName }}"` // {{ .Col.ColumnName }}
{{- end }}{{- end }}
}
// {{ .Name }} represents a row from '{{ $table }}'.
type {{ .Name }} struct {
{{- range .Fields }}
	{{ .Name }} {{ retype .Type }} `db:"{{ .Col.ColumnName }}"` // {{ .Col.ColumnName }}
{{- end }}
}

// Insert{{ .Name }} is insert model to {{ $table }}
func (t *{{ .Name }}InsertArgs) Insert(dbAccessor acallsql.DBAccessor) (*{{ .Name }},error) {
	const sql = `
		INSERT INTO {{ .Schema }}.{{ $table }} ({{$first := true}}
{{- range .Fields }}{{ if ne .Name "ID" }}
			{{ if eq $first false }}, {{- end }}{{$first = false}}"{{ .Col.ColumnName }}"
{{- end }}{{- end }}{{$first = true}}
		) VALUES (
{{- range .Fields }}{{ if ne .Name "ID" }}
			{{ if eq $first false }}, {{- end }}{{$first = false}}:{{ .Col.ColumnName }}
{{- end }}{{- end }}
		)
		RETURNING{{$first = true}}
{{- range .Fields }}
			{{ if eq $first false }}, {{- end }}{{$first = false}}"{{ .Col.ColumnName }}"
{{- end }}
	`
	var obj {{ .Name }}
	if err := dbAccessor.ExecuteCallback(sql, t, func(rows acallsql.DBRows) error {
		for rows.Next() {
			rows.StructScan(&obj)
			return nil
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return &obj, nil
}
