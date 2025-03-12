package postgres

var (
	queryCreateProduct = `
		INSERT INTO products (
			id, 
			category_id, 
			supplier_id, 
			unit_id, 
			name, 
			description, 
			base_price, 
			stock, 
			created_at, 
			created_by
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`

	queryListProduct = `
		SELECT
			p.id,
			p.category_id,
			p.supplier_id,
			p.unit_id,
			p.name,
			p.description,
			p.base_price,
			p.stock,
			p.created_at,
			p.created_by,
			p.updated_at,
			p.updated_by
		FROM products p		
	`

	queryGetListProduct = queryListProduct + `
		JOIN categories c on p.category_id = c.id
		WHERE 1=1
	`

	queryGetProductByID = queryListProduct + `
		WHERE p.id = $1
	`

	queryGetProductByName = queryListProduct + `
		WHERE 
			p.category_id = $1 AND 
			p.name = $2
	`
)
