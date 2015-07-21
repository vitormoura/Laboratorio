package model

type VFStorageFactory interface {

	//Recupera relação de storages disponíveis para criação na fábrica
	GetAvaiableStorages() ([]string, error)

	//Recupera storage adequado para aplicação identificada pelo ID informado
	Create(appID string) (VFStorage, error)
}
