package user

import "gorm.io/gorm"

// buat interface, Repository menggunakan R besar jadi bisa di akses dari package lain
type Repository interface{
	// buat fungsi save yang parameternya struct User, balikannya User dan error
	Save(user User)(User, error)
}

// buat struct Repository, pada struct respository r nya kecil berarti struct tersebut berdifat private atau tidak bisa di akses di package lain
type repository struct{
	// untuk mengakses koneksi ke database
	db *gorm.DB
}

// membuat function untuk menggunakan db (object baru dari struct repository)
func NewRepository(db *gorm.DB) *repository{
	return &repository{db}
}

// buat function save yang mengikuti kontrak dari interface 
func (r *repository) Save(user User) (User, error){
		// save user
		err := r.db.Create(&user).Error

		if err != nil{
			return user, err
		}

		return user, nil
}