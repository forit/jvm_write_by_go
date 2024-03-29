package classpath

import (
	"errors"
	"fmt"
	"strings"
)

type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	compositeEntry := []Entry{}
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}


func (self CompositeEntry) readClass(className string) ([]byte, Entry, error)  {
	for _, entry := range self {
		data, entryOther, err := entry.readClass(className);
		if(err == nil) {
			return data, entryOther, err
		}
	}

	return nil, nil, errors.New(fmt.Sprintf("class nit found %s", className))
}

func (self CompositeEntry) String() string {
	strs := make([]string, len(self))
	for i, entry := range self {
		strs[i] = entry.String()
	}
	return strings.Join(strs, pathListSeparator)
}