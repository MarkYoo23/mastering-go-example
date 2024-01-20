package entries

import "sort"

type EntryService struct {
	entryRepository *EntryRepository
}

func NewEntryService(entryRepository *EntryRepository) *EntryService {
	return &EntryService{entryRepository}
}

func (s *EntryService) GetSortByFullName() (Entries, error) {
	entities, err := s.entryRepository.GetAll()
	if err != nil {
		return nil, err
	}

	var entries Entries

	for _, entity := range entities {
		entry := NewEntry(*entity)
		entries = append(entries, entry)
	}

	sort.Sort(entries)

	return entries, nil
}
