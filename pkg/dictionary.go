package pkg

import "errors"

// Error Codes:
// 	101: dictionary hasn't been created
//  102: word has been in dictionary

var dictionary map[string]string

// NewDictionary creates new dictionary
func NewDictionary() {
	dictionary = make(map[string]string, 0)
}

// AddNewWord adds new word to dictionary
// Word needs to be unique and never played before
func AddNewWord(word string, meaning string) error {
	if err := validateNewWord(word); err != nil {
		return err
	}
	dictionary[word] = meaning
	return nil
}

// validateNewWord validates whether the new word added to dictionary is unique or have been existed
func validateNewWord(word string) error {
	if dictionary == nil {
		return errors.New("101")
	}
	_, existed := dictionary[word]
	if existed {
		return errors.New("102")
	}
	return nil
}

// GetDictionary returns list of words added
func GetDictionary() (map[string]string, bool) {
	// if no dictionary has been created, return false
	if dictionary == nil {
		return nil, false
	}
	return dictionary, true
}
