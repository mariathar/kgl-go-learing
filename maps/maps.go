package maps

// Хеш-таблица является одной из самых гениальных и универсальных из всех
// структур данных. Это неупорядоченная коллекция пар “ключ-значение”, в которой
// все ключи различны, а значение, связанное с заданным ключом, можно получить, обновить или удалить с использованием в среднем константного количества сравнений
// ключей, независимо от размера хеш-таблицы.
//just dictionary
const (
	ErrNotFound = DictionaryErr("could not find the word you were looking for")

	ErrWordExists = DictionaryErr("cannot add word because it already exists")

	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		//if not found - add item(not clear)
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		//if no errors - update(not clear)
		d[word] = definition
	default:
		return err

	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
