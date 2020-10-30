package entities

// Book é um tipo no qual contém o nome do livro, a abreviação dele e todos os capítulos e versículos.
type Book struct {
	Name     string
	Abbrev   string
	Chapters [][]string
}
