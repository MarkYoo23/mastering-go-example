package entries

import "white-page/ent"

type Entry struct {
	ent.Entry
}

func NewEntry(entry ent.Entry) *Entry {
	return &Entry{Entry: entry}
}

type Entries []*Entry

func (e Entries) Len() int {
	return len(e)
}

func (e Entries) Less(i, j int) bool {
	EntryI := e[i]
	EntryJ := e[j]

	if EntryI.Surname == EntryJ.Surname {
		return EntryI.Name < EntryJ.Name
	}

	return EntryI.Surname < EntryJ.Surname
}

func (e Entries) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}
