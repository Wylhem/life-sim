# Jeu de la vie
Bienvenue dans une simulation du jeu de la vie inventé par John Horton Conway en 1970. 

## 📏 Règles
1. **Le plateau** : 
   - Le jeu se déroule sur une grille de cellules qui peuvent être vivantes ou mortes. La grille peut être infinie, mais est souvent limitée pour les besoins de simulation.

2. **Les cellules** :
   - Chaque cellule a huit voisins, qui sont les cellules adjacentes horizontalement, verticalement et diagonalement.

3. **Les états** :
   - Les cellules peuvent être dans deux états : vivante ou morte.

4. **Évolution** :
   - Le jeu se déroule en tours (ou générations), et l'état de chaque cellule change à chaque tour selon les règles suivantes :

      - **Naissance** : Une cellule morte avec exactement trois cellules voisines vivantes devient vivante (naît) au tour suivant.
      - **Survie** : Une cellule vivante avec deux ou trois cellules voisines vivantes reste vivante au tour suivant.
      - **Mort par isolement** : Une cellule vivante avec moins de deux cellules voisines vivantes meurt (d'isolement) au tour suivant.
      - **Mort par surpopulation** : Une cellule vivante avec plus de trois cellules voisines vivantes meurt (de surpopulation) au tour suivant.

Ces règles simples peuvent produire des comportements très complexes et intéressants sur la grille, donnant lieu à des motifs qui se déplacent, oscillent ou restent stables.

Si vous avez besoin d'une démonstration visuelle ou d'un programme pour simuler le Jeu de la vie, je peux vous aider à en créer un.

## 👇 Installation

Pour installer ce projet, suivez les étapes suivantes :
  
```bash

git clone https://github.com/votrecompte/votreprojet.git

cd votreprojet

go mod download
```
```bash
tree

