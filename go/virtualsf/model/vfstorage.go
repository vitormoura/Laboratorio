package model

//VFStorage representa um componente capaz de armazenar e localizar arquivos
type VFStorage interface {

	//Add adiciona um novo arquivo ao sistema de armazenamento,
	//preenchendo o identificador único do arquivo em caso de sucesso
	Add(f *File) error

	//Find localiza o arquivo identificado pelo identificador informado
	Find(id string) (*File, error)

	//Remove exclui arquivo identificado pelo id informado
	Remove(id string) error

	//List recupera relação de arquivos do storage
	List() ([]FileInfo, error)

	//Stats recupera estatísticas sobre a situação de armazenamento do storage
	Stats() (*VFStorageStats, error)

	//Refresh dispara rotinas de internas para redefinir estado do storage (Quando aplicável)
	Refresh() error
}
