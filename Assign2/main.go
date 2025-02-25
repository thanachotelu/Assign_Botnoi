package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PokemonRequest struct {
	ID string `json:"id"`
}

type PokemonStat struct {
	BaseStat int  `json:"base_stat"`
	Effort   int  `json:"effort"`
	Stat     Stat `json:"stat"`
}

type Stat struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Pokemonsprites struct {
	BackDefault      *string `json:"back_default"`
	BackFemale       *string `json:"back_female"`
	BackShiny        *string `json:"back_shiny"`
	BackShinyFemale  *string `json:"back_shiny_female"`
	FrontDefault     *string `json:"front_default"`
	FrontFemale      *string `json:"front_female"`
	FrontShiny       *string `json:"front_shiny"`
	FrontShinyFemale *string `json:"front_shiny_female"`
}

type PokemonResponse struct {
	Stats   []PokemonStat  `json:"stats"`
	Name    string         `json:"name"`
	Sprites Pokemonsprites `json:"sprites"`
}

type PokemonAPIResponse struct {
	Stats   []PokemonStat  `json:"stats"`
	Name    string         `json:"name"`
	Sprites Pokemonsprites `json:"sprites"`
}

func emptywithnull(str *string) *string {
	if str == nil || *str == "" {
		return nil
	}
	return str
}

func fetchAPI(url string, data interface{}) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("failed to fetch: %s", url)
	}
	return json.NewDecoder(res.Body).Decode(data)
}

func getPokemon(c *gin.Context) {
	var id PokemonRequest
	if err := c.ShouldBindJSON(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// ดึงข้อมูลจาก API `/pokemon/{id}/`
	pokemonAPIURL := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", id.ID)
	var pokeData PokemonAPIResponse
	if err := fetchAPI(pokemonAPIURL, &pokeData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch pokemon data"})
		return
	}

	// ดึงข้อมูลจาก API `/pokemon-form/{id}/`
	formAPIURL := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon-form/%s/", id.ID)
	var pokeFormData PokemonAPIResponse
	if err := fetchAPI(formAPIURL, &pokeFormData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch pokemon form data"})
		return
	}

	var filterStats []PokemonStat
	for _, stat := range pokeData.Stats {
		if stat.Stat.Name == "hp" || stat.Stat.Name == "attack" {
			filterStats = append(filterStats, stat)
		}
	}

	formatstrSprites := Pokemonsprites{
		BackDefault:      emptywithnull(pokeData.Sprites.BackDefault),
		BackFemale:       emptywithnull(pokeData.Sprites.BackFemale),
		BackShiny:        emptywithnull(pokeData.Sprites.BackShiny),
		BackShinyFemale:  emptywithnull(pokeData.Sprites.BackShinyFemale),
		FrontDefault:     emptywithnull(pokeData.Sprites.FrontDefault),
		FrontFemale:      emptywithnull(pokeData.Sprites.FrontFemale),
		FrontShiny:       emptywithnull(pokeData.Sprites.FrontShiny),
		FrontShinyFemale: emptywithnull(pokeData.Sprites.FrontShinyFemale),
	}

	responseData := PokemonResponse{
		Stats:   filterStats,
		Name:    pokeFormData.Name,
		Sprites: formatstrSprites,
	}

	c.JSON(http.StatusOK, responseData)
}

func main() {

	r := gin.Default()
	r.POST("/pokemon", getPokemon)
	log.Fatal(r.Run(":8080"))
}
