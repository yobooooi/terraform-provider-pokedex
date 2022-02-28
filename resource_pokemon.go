package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const pokemonAPI string = "https://pokeapi.co/api/v2/pokemon/"

func resourcePokemon() *schema.Resource {
	return &schema.Resource{
		Create: resourcePokemonCreate,
		Read:   resourcePokemonRead,
		Update: resourcePokemonUpdate,
		Delete: resourcePokemonDelete,

		Schema: map[string]*schema.Schema{
			"pokemon_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourcePokemonCreate(d *schema.ResourceData, m interface{}) error {
	pokemon_id := d.Get("pokemon_id").(string)

	d.SetId(pokemon_id)

	resp, err := http.Get(pokemonAPI + pokemon_id)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	var data map[string]interface{}

	err = json.Unmarshal([]byte(body), &data)

	if err != nil {
		panic(err.Error())
	}
	d.Set("name", data["name"])
	return resourcePokemonRead(d, m)
}

func resourcePokemonRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourcePokemonUpdate(d *schema.ResourceData, m interface{}) error {
	return resourcePokemonRead(d, m)
}

func resourcePokemonDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
