package dict

//import "errors"

type Dictionary map[string]string

type DictionaryErr string

const (
    ErrNotFound = DictionaryErr("could not find the key you are looking for")
    ErrExists = DictionaryErr("key already exists")
    ErrKeyDoesNotExist = DictionaryErr("input key does not exist")
)

func (e DictionaryErr) Error() string {
    return string(e)
}

func (d Dictionary) Delete(key string) {
    delete(d, key)
}

func (d Dictionary) Update(key, val string) error {

    _, err := d.Search(key)

    switch err {
    
    case ErrNotFound:
        return ErrKeyDoesNotExist
    case nil:
        d[key] = val
    default:
        return err
    }

    return nil
}
    
func (d Dictionary) Add(key, val string) error {

    _, err := d.Search(key)

    switch err {
    case ErrNotFound:
        d[key] = val
    case nil:
        return ErrExists
    default:
        return err
    }

    return nil
}

func (d Dictionary) Search(key string) (string, error) {

    value, ok := d[key]

    if !ok {
        return "", ErrNotFound
    }
   
    return value, nil
}
