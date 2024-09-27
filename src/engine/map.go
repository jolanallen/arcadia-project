package engine

import (
	"encoding/json"    // Permet de travailler avec des fichiers JSON.
	"fmt"              // Utilisé pour l'affichage de messages d'erreur.
	"os"               // Utilisé pour la gestion des fichiers (lecture, fermeture, etc.).
	"path"             // Permet de manipuler les chemins de fichiers.

	rl "github.com/gen2brain/raylib-go/raylib" // Librairie raylib pour la gestion graphique.
)

type Layer struct {
	Data    []int   `json:"data"`    // Les données des tuiles (ID des tuiles dans chaque calque).
	Height  int     `json:"height"`  // Hauteur du calque en nombre de tuiles.
	ID      int     `json:"id"`      // Identifiant du calque.
	Name    string  `json:"name"`    // Nom du calque (utile pour identifier des couches spécifiques comme "collisions").
	Opacity float32 `json:"opacity"` // Transparence du calque.
	Type    string  `json:"type"`    // Type de calque (ex : "tilelayer").
	Visible bool    `json:"visible"` // Visibilité du calque.
	Width   int     `json:"width"`   // Largeur du calque en nombre de tuiles.
	X       int     `json:"x"`       // Coordonnée X du calque.
	Y       int     `json:"y"`       // Coordonnée Y du calque.
}

type TileSet struct {
	Columns     int    `json:"columns"`     // Nombre de colonnes de tuiles dans l'image du tileset.
	FirstGid    int    `json:"firstgid"`    // Premier identifiant de tuile dans ce tileset.
	Image       string `json:"image"`       // Chemin vers l'image du tileset.
	ImageHeight int    `json:"imageheight"` // Hauteur de l'image du tileset.
	ImageWidth  int    `json:"imagewidth"`  // Largeur de l'image du tileset.
	Margin      int    `json:"margin"`      // Espace entre les tuiles dans l'image.
	Name        string `json:"name"`        // Nom du tileset.
	Spacing     int    `json:"spacing"`     // Espacement entre les tuiles dans l'image.
	TileCount   int    `json:"tilecount"`   // Nombre total de tuiles dans le tileset.
	TileHeight  int    `json:"tileheight"`  // Hauteur d'une tuile.
	TileWidth   int    `json:"tilewidth"`   // Largeur d'une tuile.
}

type MapJSON struct {
	CompressionLevel int       `json:"compressionLevel"` // Niveau de compression du fichier de carte.
	Height           int       `json:"height"`           // Hauteur de la carte en nombre de tuiles.
	Infinite         bool      `json:"infinite"`         // Indique si la carte est infinie.
	Layers           []Layer   `json:"layers"`           // Liste des calques de la carte.
	NextLayerID      int       `json:"nextlayerid"`      // Prochain ID de calque.
	NextObjectID     int       `json:"nextobjectid"`     // Prochain ID d'objet.
	Orientation      string    `json:"orientation"`      // Orientation de la carte (orthogonale, isométrique, etc.).
	RenderOrder      string    `json:"renderorder"`      // Ordre de rendu des calques.
	TiledVersion     string    `json:"tiledversion"`     // Version du logiciel Tiled utilisée.
	TileHeight       int       `json:"tileheight"`       // Hauteur d'une tuile.
	TileSets         []TileSet `json:"tilesets"`         // Liste des tilesets utilisés dans la carte.
	TileWidth        int       `json:"tilewidth"`        // Largeur d'une tuile.
	Type             string    `json:"type"`             // Type de fichier ("map").
	Version          string    `json:"version"`          // Version de la carte.
	Width            int       `json:"width"`            // Largeur de la carte en nombre de tuiles.
}


func (e *Engine) InitMap(mapFile string) {
	file, err := os.ReadFile(mapFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	json.Unmarshal(file, &e.MapJSON)

	//Load all required textures from TileSets
	for _, TileSet := range e.MapJSON.TileSets {
		path := path.Dir(mapFile) + "/"
		e.Sprites[TileSet.Name] = rl.LoadTexture(path + TileSet.Image)
	}
}

func (e *Engine) RenderMap() {
	// Prépare les rectangles source (dans l'image du tileset) et destination (dans la fenêtre).
	// Seuls X et Y changeront à chaque tuile, la largeur et la hauteur restent fixes.
	srcRectangle := rl.Rectangle{X: 0, Y: 0, Width: float32(e.MapJSON.TileHeight), Height: float32(e.MapJSON.TileHeight)}
	destRectangle := rl.Rectangle{X: 0, Y: 0, Width: float32(e.MapJSON.TileWidth), Height: float32(e.MapJSON.TileWidth)}
	column_counter := 0 // Compte le nombre de colonnes pour gérer les retours à la ligne.

	for _, Layer := range e.MapJSON.Layers { // Parcourt tous les calques (layers) de la carte.
		for _, tile := range Layer.Data { // Parcourt toutes les tuiles dans chaque calque.
			if tile != 0 { // Si le tile est différent de 0 (0 signifie aucune tuile à cet endroit).
				wantedTileSet := e.MapJSON.TileSets[0]
				for _, TileSet := range e.MapJSON.TileSets { 
					if TileSet.FirstGid <= tile {
						wantedTileSet = TileSet
					}
				}

				index := tile - wantedTileSet.FirstGid

				srcRectangle.X = float32(index)
				srcRectangle.Y = 0

				if index >= wantedTileSet.Columns { 
					srcRectangle.X = float32(index % wantedTileSet.Columns)
					srcRectangle.Y = float32(index / wantedTileSet.Columns)
				}
				if Layer.Name == "objet" {
					e.ColisionListe = append(e.ColisionListe, rl.NewRectangle(destRectangle.X - 32, destRectangle.Y - 32, destRectangle.Width, destRectangle.Height))
				}
				srcRectangle.X *= float32(e.MapJSON.TileWidth)
				srcRectangle.Y *= float32(e.MapJSON.TileHeight)
				rl.DrawTexturePro(
					e.Sprites[wantedTileSet.Name],
					srcRectangle,
					destRectangle,
					rl.Vector2{X: 0, Y: 0},
					0,
					rl.White,
				)
			}

			
			destRectangle.X += 16
			column_counter += 1
			if destRectangle.Y == 0 {
				if column_counter >= e.MapJSON.Width {
					destRectangle.X = 0
					destRectangle.Y += 16
					column_counter = 0
				}
				
			}
			if column_counter >= e.MapJSON.Width {
				destRectangle.X = 0
				destRectangle.Y += 16
				column_counter = 0
			}

		}
		destRectangle.X, destRectangle.Y, column_counter = 0, 0, 0
	}
}
