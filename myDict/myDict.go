package myDict

import "errors"

// Dictionary type
type Dictionary map[string]string

var (
  errNotFound = errors.New("Not Found.")
  errWordExists = errors.New("This word already exists.")
  errCantUpdate = errors.New("Can't update non-existing word.")
  )
// Search a word
func (d Dictionary) Search(word string) (string, error) {
  value, exists := d[word]
  if exists {
    return value, nil
  }
  return "", errNotFound
}

// Add a word to the Dictionary
func (d Dictionary) Add(word, def string) error {
  _, err := d.Search(word)
  switch err{
  case errNotFound:
    d[word] = def
  case nil:
    return errWordExists
  }
  return nil
}

func (d Dictionary) Update(word, newDef string) error {
  _, error := d.Search(word)
  switch error{
  case nil:
    d[word] =  newDef
  case errNotFound:
    return errCantUpdate
  }
  return nil
}

func (d Dictionary) Delete(word string)  {
  delete(d, word)
}