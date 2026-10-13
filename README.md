# Go-reloaded

## Présentation
Le projet Go-reloaded est un outil en ligne de commande, écrit en Go, qui corrige et met en forme automatiquement un texte. Il lit un fichier d'entrée, applique une série de règles d'édition et d'autocorrection, puis écrit le résultat dans un fichier de sortie.

## Fonctionnalités
- `(hex)` et `(bin)` : convertir le mot précédent (hexadécimal ou binaire) en décimal
- `(up)`, `(low)`, `(cap)` : mettre le mot précédent en majuscules, en minuscules ou avec une majuscule initiale
- `(up, n)`, `(low, n)`, `(cap, n)` : appliquer la transformation aux `n` mots précédents
- Ponctuation `. , ! ? : ;` : collée au mot précédent, groupes comme `...` ou `!?` conservés
- Apostrophes `' '` : collées aux mots qu'elles entourent
- `a` devient `an` devant une voyelle ou un `h`
- Les retours à la ligne sont conservés

## Technologies
- Go (bibliothèque standard uniquement : `os`, `strings`, `strconv`, `fmt`)

## Prérequis
- Go 1.16 ou plus récent

## Installation
Cloner le dépôt puis se placer dans le dossier du projet :
```bash
git clone <url_du_depot>
cd go-reloaded
```

## Configuration
Aucune configuration ni variable d'environnement nécessaire.

## Lancement
```bash
go run . <fichier_entree> <fichier_sortie>
```
Exemple :
```bash
go run . sample.txt result.txt
```

## Tests
```bash
go test -v
```