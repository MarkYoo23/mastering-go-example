package entries

import (
	"context"
	"white-page/ent"
	"white-page/ent/entry"
	"white-page/internal/db"
)

type EntryRepository struct {
	dbContext *db.DbContext
}

func NewEntryRepository(dbContext *db.DbContext) *EntryRepository {
	return &EntryRepository{
		dbContext: dbContext,
	}
}

func (r *EntryRepository) Add(name string, surName string, tel string) (*ent.Entry, error) {
	client := r.dbContext.Client

	result, err := client.Entry.Create().
		SetName(name).
		SetSurname(surName).
		SetTel(tel).
		Save(context.Background())

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *EntryRepository) Get(id int) (*ent.Entry, error) {
	client := r.dbContext.Client

	result, err := client.Entry.Get(context.Background(), id)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *EntryRepository) GetAll() (ent.Entries, error) {
	client := r.dbContext.Client

	entries, err := client.Entry.Query().All(context.Background())

	if err != nil {
		return nil, err
	}

	return entries, nil
}

func (r *EntryRepository) Update(id int, name string, surName string, tel string) (*ent.Entry, error) {
	client := r.dbContext.Client

	result, err := client.Entry.UpdateOneID(id).
		SetName(name).
		SetSurname(surName).
		SetTel(tel).
		Save(context.Background())

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *EntryRepository) Delete(id int) error {
	client := r.dbContext.Client

	err := client.Entry.DeleteOneID(id).Exec(context.Background())

	if err != nil {
		return err
	}

	return nil
}

func (r *EntryRepository) DeleteAll() error {
	client := r.dbContext.Client

	_, err := client.Entry.Delete().Exec(context.Background())

	if err != nil {
		return err
	}

	return nil
}

func (r *EntryRepository) FindByPhone(tel string) (*ent.Entry, error) {
	client := r.dbContext.Client

	result, err := client.Entry.Query().
		Where(entry.TelIn(tel)).
		First(context.Background())

	if err != nil {
		return nil, err
	}

	return result, nil
}
