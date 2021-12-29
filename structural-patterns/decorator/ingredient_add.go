package decorator

type IngredientAdd interface {
	AddIngredient() (string, error)
}
