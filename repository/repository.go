package repository

import "gorm.io/gorm"

type Repo struct {
}

func (r *Repo) Create(db *gorm.DB, i interface{}) error {
	if err := db.
		Set("gorm:association_autoupdate", false).
		Set("gorm:association_autocreate", false).
		Create(i).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repo) FindOneObjectByField(db *gorm.DB, field string, value interface{}, i interface{}) error {
	if err := db.Where(field+" = ?", value).First(i).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repo) Update(db *gorm.DB, i interface{}) error {
	if err := db.
		Set("gorm:association_autoupdate", false).
		Set("gorm:association_autocreate", false).
		Save(i).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repo) Delete(db *gorm.DB, i interface{}) error {
	if err := db.
		Set("gorm:association_autoupdate", false).
		Set("gorm:association_autocreate", false).
		Delete(i).Error; err != nil {
		return err
	}
	return nil
}
