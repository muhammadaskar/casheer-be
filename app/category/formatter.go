package category

import "github.com/muhammadaskar/casheer-be/domains"

type CategoryFormatter struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func FormatCategory(category domains.Category) CategoryFormatter {
	categoryFormatter := CategoryFormatter{}
	categoryFormatter.ID = category.ID
	categoryFormatter.Name = category.Name

	return categoryFormatter
}

func FormatCategories(categories []domains.Category) []CategoryFormatter {
	categoriesFormatter := []CategoryFormatter{}

	for _, category := range categories {
		categoryFormatter := FormatCategory(category)
		categoriesFormatter = append(categoriesFormatter, categoryFormatter)
	}
	return categoriesFormatter
}
