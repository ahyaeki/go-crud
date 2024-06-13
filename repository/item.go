package repository

import (
	"database/sql"
	"go-crud/model"
)

func GetItems(db *sql.DB) ([]model.Item, error) {
	rows, err := db.Query("SELECT id, name, description FROM items")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []model.Item
	for rows.Next() {
		var item model.Item
		if err := rows.Scan(&item.ID, &item.Name, &item.Description); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

func CreateItem(db *sql.DB, item *model.Item) error {
	if item.ID != 0 {
		_, err := db.Exec("INSERT INTO items (id, name, description) VALUES ($1, $2, $3)", item.ID, item.Name, item.Description)
		return err
	} else {
		return db.QueryRow("INSERT INTO items (name, description) VALUES ($1, $2) RETURNING id", item.Name, item.Description).Scan(&item.ID)
	}
}

func GetItemByID(db *sql.DB, id string) (model.Item, error) {
	var item model.Item
	err := db.QueryRow("SELECT id, name, description FROM items WHERE id = $1", id).Scan(&item.ID, &item.Name, &item.Description)
	return item, err
}

func UpdateItem(db *sql.DB, id string, updates map[string]interface{}) error {
	// Assuming only name and description can be updated
	var query string
	var args []interface{}
	if name, ok := updates["name"]; ok {
		query += "name = $1"
		args = append(args, name)
	}
	if description, ok := updates["description"]; ok {
		if len(args) > 0 {
			query += ", "
		}
		query += "description = $2"
		args = append(args, description)
	}
	if len(args) == 0 {
		return nil
	}
	args = append(args, id)
	query = "UPDATE items SET " + query + " WHERE id = $" + string(rune(len(args)))

	_, err := db.Exec(query, args...)
	return err
}

func DeleteItem(db *sql.DB, id string) error {
	_, err := db.Exec("DELETE FROM items WHERE id = $1", id)
	return err
}
