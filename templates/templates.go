package templates

import "embed"

// FS expõe os arquivos de template embutidos para o resto da aplicação.
//
//go:embed oci/*.yaml
var FS embed.FS
