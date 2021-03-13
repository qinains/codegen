const Mock = require('mockjs')

const data = Mock.mock({
  'items|100': [{
{{- range $k,$v := .table.columns -}}
{{- if (Contains $v.columnName "id") -}}
{{$v.columnName | Camel}}: '@id', // {{$v.columnComment}}
{{- else if (or (Contains $v.columnName "_url")) -}}
{{$v.columnName | Camel}}: '@image', // {{$v.columnComment}}
{{- else if (or (Contains $v.columnName "content")) -}}
{{$v.columnName | Camel}}: '@cparagraph', // {{$v.columnComment}}
{{- else if (or (Contains $v.columnName "_time") (Contains $v.columnName "_at")) -}}
{{$v.columnName | Camel}}: '@datetime("T")', // {{$v.columnComment}}
{{- else if (or (Contains $v.columnName "status")) -}}
'{{$v.columnName | Camel}}|1': [10, 20, 30, 40, 50], // {{$v.columnComment}}
{{- else if Contains $v.columnName "sort" -}}
{{$v.columnName | Camel}}: '@integer(0, 50000)', // {{$v.columnComment}}
{{- else if IsNumberDataType $v.dataType -}}
{{$v.columnName | Camel}}: '@integer(0, 50000)', // {{$v.columnComment}}
{{- else -}}
{{$v.columnName | Camel}}: '@csentence(2, 30)', // {{$v.columnComment}}
{{- end}}
{{end}}
  }]
})

module.exports = [
  {
    url: '/{{.table.tableName | Dash}}/create-{{.table.tableName | Dash}}',
    type: 'post',
    response: config => {
      const items = data.items
      return {
        code: 20000,
        data: items[0]
      }
    }
  },
  {
    url: '/{{.table.tableName | Dash}}/delete-{{.table.tableName | Dash}}',
    type: 'post',
    response: config => {
      const items = data.items
      return {
        code: 20000,
        data: items[0]
      }
    }
  },
  {
    url: '/{{.table.tableName | Dash}}/update-{{.table.tableName | Dash}}',
    type: 'post',
    response: config => {
      const items = data.items
      return {
        code: 20000,
        data: items[0]
      }
    }
  },
  {
    url: '/{{.table.tableName | Dash}}/find-{{.table.tableName | Dash}}-list',
    type: 'post',
    response: config => {
      const items = data.items
      return {
        code: 20000,
        data: {
          total: items.length,
          items: items
        }
      }
    }
  },
  {
    url: '/{{.table.tableName | Dash}}/find-{{.table.tableName | Dash}}',
    type: 'post',
    response: config => {
      const items = data.items
      return {
        code: 20000,
        data: items[0]
      }
    }
  }
]
