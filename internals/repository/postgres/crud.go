package postgres

func (rep *Repository) Create(ent interface{}) error {
	return rep.Db.Connection.Create(ent).Error
}

func (rep *Repository) List(filters map[string]interface{}, order interface{}, ent interface{}) error {
	return rep.Db.Connection.Where(filters).Order(order).Find(ent).Error
}

func (rep *Repository) Read(filters interface{}, ent interface{}) error {
	res := rep.Db.Connection.Where(filters).Find(ent)
	return res.Error
}

func (rep *Repository) Update(filters interface{}, data interface{}) error {
	res := rep.Db.Connection.Where(filters).Updates(data)
	return res.Error
}

func (rep *Repository) Delete(filters interface{}, soft bool, ent interface{}) error {
	if soft {
		res := rep.Db.Connection.Where(filters).Updates(ent)
		return res.Error
	}
	res := rep.Db.Connection.Where(filters).Delete(ent)
	return res.Error
}