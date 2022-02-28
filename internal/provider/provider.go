package pokedex

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
    pokedex "github.com/yobooooi/terraform-provider-pokedex/internal/client"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"pokedex_pokemon": pokedex.resourcePokemon(),
		},
	}
}
