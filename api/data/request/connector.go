package request

type GetEmailFilterRequest struct {
	FolderID              *string  `json:"folder_id,omitempty"`
	Filter                *string  `json:"filter,omitempty"`                  // Condição de filtro
	Count                 *bool    `json:"count,omitempty"`                   // Flag para retornar a contagem total
	Expand                []string `json:"expand,omitempty"`                  // Propriedades adicionais para expandir
	IncludeHiddenMessages *bool    `json:"include_hidden_messages,omitempty"` // Flag para incluir mensagens ocultas
	Orderby               []string `json:"orderby,omitempty"`                 // Ordenação dos resultados
	Search                *string  `json:"search,omitempty"`                  // Termo de pesquisa
	Select                []string `json:"select,omitempty"`                  // Seleção de campos específicos
	Skip                  *int32   `json:"skip,omitempty"`                    // Número de registros para ignorar
	Top                   *int32   `json:"top,omitempty"`                     // Número máximo de registros a retornar
}
type MarkEmailIDRequest struct {
	ID *[]string `json:"id,omitempty"`
}

type MoveToRequest struct {
	ID     *[]string `json:"id,omitempty"`
	Folder *string   `json:"folder,omitempty"`
}

type GetFoldersRequest struct {
	Folder *string `json:"folder,omitempty"`
}
