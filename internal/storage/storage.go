package storage
import "github.com/EvaCaufmann/storage/internal/file"

type Storage struct {}

func NewStorage() *Storage{
	return &Storage	{}
	}

func (s *Storage) Upload(filename string, blob []byte) (*file.File, error){
	
	return file.NewFile(filename, blob)
	// if err != nil {
	// 	return nil, err
	// }
	// //если ошибка то стараться возвращать остальные аргументы пустыми
	// // если возвращается указатель и ошибка нил по этому указателю на этот аргумент
	// // нужно вернуть не нил объект
	// return newFile
}
