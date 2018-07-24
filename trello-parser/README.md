# Trello parser

Une truelle pour générer rapidement la section Sources depuis le board trello.

Il y a un peu de set-up à faire pour obtenir le token de Trello, mais après, ça roule !


# Obtenir le token

https://trello.com/app-key

Ouais, parce que j'ai la flemme d'intégrer la séquence d'autorisation.


# Usage 

	./parser -appid "******" -token "******" -card "3 juillet"

Si plusieurs cards sont retournées par la recherche, l'app propose de choisir la bonne :

	3  results found. Please select:
	5b367504d46dae19a6b63400         3 juillet
	5b3c32b646e1a70b4496cff8         10 Juillet 2018
	5ab904026969c4519bd7cff3         27 Mars 2018
	Enter the selected ID:

À la fin, ça sort un output.md dans le répertoire courant

# Cross-compile

	$ GOOS=windows GOARCH=386 go build -o parser.exe main.go
