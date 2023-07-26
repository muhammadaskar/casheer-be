package category

type CategoryFormatter struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func FormatCategory(category Category) CategoryFormatter {
	categoryFormatter := CategoryFormatter{}
	categoryFormatter.ID = category.ID
	categoryFormatter.Name = category.Name

	return categoryFormatter
}

func FormatCategories(categories []Category) []CategoryFormatter {
	categoriesFormatter := []CategoryFormatter{}

	for _, category := range categories {
		categoryFormatter := FormatCategory(category)
		categoriesFormatter = append(categoriesFormatter, categoryFormatter)
	}
	return categoriesFormatter
}
