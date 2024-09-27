package engine

import (
	"main/src/entity" // Importation des entités du jeu, telles que Player, Monster, et Button.

	rl "github.com/gen2brain/raylib-go/raylib" // Importation du package raylib pour la gestion graphique.
)

// Définition d'un type personnalisé "menu" pour représenter les états du menu principal.
type menu int

// Constantes pour les différents états du menu.
const (
	HOME     menu = iota // État pour l'écran d'accueil.
	SETTINGS menu = iota // État pour les paramètres du jeu.
	PLAY     menu = iota // État pour commencer ou reprendre le jeu.
)

// Définition d'un type personnalisé "engine" pour représenter les différents états du jeu.
type engine int

// Constantes pour les différents états de l'engine du jeu.
const (
	INGAME    engine = iota // Le joueur est dans le jeu actif (mode exploration).
	PAUSE     engine = iota // Le jeu est en pause.
	GAMEOVER  engine = iota // L'écran Game Over.
	LORE      engine = iota // Écran affichant l'histoire (lore) du jeu.
	WIN       engine = iota // Écran de victoire.
	INFIGHT   engine = iota // Le joueur est en combat.
	INVENTORY engine = iota // L'écran d'inventaire où le joueur peut gérer ses objets.
)

// Structure principale du moteur de jeu "Engine" qui contient les états et les ressources du jeu.
type Engine struct {
	Title                rl.Texture2D    // Texture pour le titre du jeu.
	Background           rl.Texture2D    // Texture pour le fond d'écran.
	BgSourceX            int             // Coordonnée X de la source de l'arrière-plan.
	BgSourceY            int             // Coordonnée Y de la source de l'arrière-plan.
	BackgroundFrameCount int             // Compteur de frames pour l'animation de fond.
	QuitButton           entity.Button   // Bouton pour quitter le jeu.
	StartButton          entity.Button   // Bouton pour démarrer ou reprendre le jeu.
	ScreenWidth          int32           // Largeur de la fenêtre de jeu.
	ScreenHeight         int32           // Hauteur de la fenêtre de jeu.
	Timer                float64         // Timer global pour diverses utilisations (mouvements, etc.).
	InventoryUI          rl.Texture2D    // Texture de l'interface utilisateur pour l'inventaire.
	ColisionListe        []rl.Rectangle  // Liste des rectangles de collision dans le jeu.
	GameOver             rl.Texture2D    // Texture pour l'écran de Game Over.
	Win                  rl.Texture2D    // Texture pour l'écran de victoire.
	BigJump              bool            // Indicateur si le joueur peut faire un "grand saut".
	SupSpeed             bool            // Indicateur si le joueur a une vitesse supplémentaire (power-up).
	Robotsentence        []string        // Liste des phrases prononcées par les robots (ou d'autres dialogues).
	
	Player   entity.Player   // Structure représentant le joueur du jeu.
	Monsters []entity.Monster // Liste des monstres présents dans le jeu.

	Music       rl.Music      // Musique de fond du jeu.
	MusicVolume float32       // Volume de la musique.

	Sprites          map[string]rl.Texture2D // Carte pour les sprites utilisés dans le jeu (ex. : différents objets, personnages).
	SpriteLife       rl.Texture2D            // Texture représentant la jauge de vie du joueur.
	SpriteMoney      rl.Texture2D            // Texture représentant la monnaie ou les points du joueur.
	SpriteInventaire rl.Texture2D            // Texture représentant l'inventaire du joueur.

	Camera rl.Camera2D // Caméra pour la gestion de la vue dans le jeu (suivi du joueur, par exemple).

	MapJSON MapJSON // Représentation de la carte du jeu, chargée à partir d'un fichier JSON.

	IsRunning   bool   // Indique si le jeu est en cours d'exécution.
	StateMenu   menu   // État actuel du menu (HOME, SETTINGS, PLAY, etc.).
	StateEngine engine // État actuel du jeu (INGAME, PAUSE, GAMEOVER, etc.).

	StartedFightCountFrames int         // Compteur de frames depuis le début du combat.
	StartedFight            rl.Texture2D // Texture affichée au début d'un combat.

	FondFight rl.Texture2D // Texture pour le fond lors des combats.

	FontMedieval rl.Font // Police d'écriture médiévale utilisée pour les textes dans le jeu.
	FontFreshman rl.Font // Police d'écriture de type "Freshman" utilisée pour d'autres textes.
}
