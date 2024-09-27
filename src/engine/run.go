package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib" // Importation du package raylib pour gérer le rendu graphique et l'interaction.
)

func (engine *Engine) Run() {
	rl.SetTargetFPS(60) // Définit le nombre d'images par seconde à 60, garantissant une fluidité du jeu.

	for engine.IsRunning { // Boucle principale du jeu qui s'exécute tant que le jeu est en cours d'exécution.

		rl.BeginDrawing() // Commence une nouvelle frame de dessin.

		// Gestion de l'état principal du menu (ex. : HOME, SETTINGS, PLAY).
		switch engine.StateMenu {
		case HOME:
			engine.HomeRendering() // Rendu de l'écran d'accueil.
			engine.HomeLogic()     // Gère la logique de l'écran d'accueil (interactions, etc.).

		case SETTINGS:
			engine.SettingsLogic() // Gère la logique des paramètres du jeu (options, contrôles, etc.).

		case PLAY:
			// Gestion de l'état interne du jeu lorsque le joueur est en mode "PLAY".
			switch engine.StateEngine {
			case INGAME:
				engine.InGameRendering() // Rendu du jeu (carte, joueur, etc.).
				engine.InGameLogic()     // Gère la logique du jeu (mouvements, collisions, etc.).

			case PAUSE:
				engine.PauseRendering() // Rendu de l'écran de pause.
				engine.PauseLogic()     // Gère la logique en mode pause (reprendre ou quitter).

			case INVENTORY:
				engine.InventoryRendering() // rendu de  l'écran d'inventaire.
				engine.InventoryLogic()     // Gère la logique de l'inventaire (interaction avec les objets).

			case INFIGHT:
				engine.FightRendering() // Rendu de l'écran pendant un combat.
				engine.FightLogic()     // Gère la logique de combat (attaques, défenses, etc.).

			case GAMEOVER:
				engine.GAMEOver() // Affiche l'écran de Game Over quand le joueur perd.

			case WIN:
				engine.YouWin() // Affiche l'écran de victoire quand le joueur gagne.

			case LORE:
				engine.LoreRendering() // Rendu de l'écran pour afficher le lore (histoire).
				engine.LoreLogic()     // Gère la logique de l'écran du lore (interaction avec l'histoire).
			}
		}

		rl.EndDrawing() // Termine le rendu de la frame actuelle.
	}
}
