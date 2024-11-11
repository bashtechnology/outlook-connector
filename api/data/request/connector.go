package request

type GetEmailFilterRequest struct {
	Filter                *string  `json:"filter,omitempty"`                // Condição de filtro
	Count                 *bool    `json:"count,omitempty"`                 // Flag para retornar a contagem total
	Expand                []string `json:"expand,omitempty"`                // Propriedades adicionais para expandir
	IncludeHiddenMessages *bool    `json:"includeHiddenMessages,omitempty"` // Flag para incluir mensagens ocultas
	Orderby               []string `json:"orderby,omitempty"`               // Ordenação dos resultados
	Search                *string  `json:"search,omitempty"`                // Termo de pesquisa
	Select                []string `json:"select,omitempty"`                // Seleção de campos específicos
	Skip                  *int32   `json:"skip,omitempty"`                  // Número de registros para ignorar
	Top                   *int32   `json:"top,omitempty"`                   // Número máximo de registros a retornar
}
