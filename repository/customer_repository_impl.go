package repository

import (
	"context"
	"database/sql"
	"insectdie/go-restapi-mux/helper"
	"insectdie/go-restapi-mux/model/domain"
)

type CustomerRepositoryImpl struct {
}

func NewCustomerRepository() CustomerRepository {
	return &CustomerRepositoryImpl{}
}

func (repository *CustomerRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Customer {

	var families []domain.Family
	SQL_FAM := "SELECT fl_relation, fl_name, fl_dob FROM `family_list`"
	rows_fam, err_fam := tx.QueryContext(ctx, SQL_FAM)
	helper.PanicIfError(err_fam)
	defer rows_fam.Close()
	for rows_fam.Next() {
		family := domain.Family{}
		err_fam := rows_fam.Scan(&family.Relation, &family.Name, &family.Dob)
		families = append(families, domain.Family{Relation: family.Relation, Name: family.Name, Dob: family.Dob})
		helper.PanicIfError(err_fam)
	}
	// fmt.Printf("%v", families)

	SQL := "select a.cst_id, a.cst_name, a.cst_dob, a.cst_phone_num, CONCAT(b.nationality_name, ' (',b.nationality_code, ')') as nationality_name, a.cst_email from customer a inner join nationality b on a.nationality_id = b.nationality_id"

	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	// families = append(families, domain.Family{Relation: "Bapak", Name: "Bornok", Dob: "2023-05-22"})
	// families = append(families, domain.Family{Relation: "Adik", Name: "Risa", Dob: "2000-05-09"})

	var customers []domain.Customer

	// customers = append(customers, domain.Customer{Id: 1, Name: "Andry", Dob: "1997-05-23", Nationality: "Indo", Email: "ompusunggu@gmail.com", Family: families})

	for rows.Next() {
		customer := domain.Customer{}
		err := rows.Scan(&customer.Id, &customer.Name, &customer.Dob, &customer.Phone, &customer.Nationality, &customer.Email)
		helper.PanicIfError(err)

		customers = append(customers, domain.Customer{Id: customer.Id, Name: customer.Name, Dob: customer.Dob, Phone: customer.Phone, Nationality: customer.Nationality, Email: customer.Email, Family: families})
	}
	return customers
}
