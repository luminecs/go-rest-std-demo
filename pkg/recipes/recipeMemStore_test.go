package recipes

import "testing"

func getHamCheeseToasties() Recipe {
	return Recipe{
		Name: "ham and cheese toastie",
		Ingredients: []Ingredient{
			{Name: "bread"},
			{Name: "ham"},
			{Name: "cheese"},
		},
	}
}

func TestMemStore_Add(t *testing.T) {
	type fields struct {
		list map[string]Recipe
	}
	type args struct {
		name   string
		recipe Recipe
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		wantLen int
	}{
		{
			name: "Add to empty map",
			fields: fields{
				map[string]Recipe{},
			},
			args: args{
				name:   "ham and cheese toastie",
				recipe: getHamCheeseToasties(),
			},
			wantLen: 1,
			wantErr: false,
		},
	}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		m := MemStore{
	//			list: tt.fields.list,
	//		}
	//		err := m.Add(tt.args.name, tt.args.recipe)
	//		if !tt.wantErr {
	//			assert.NoError(t, err)
	//		}
	//		assert.Len(t, tt.fields.list, tt.wantLen)
	//	})
	//}
}
