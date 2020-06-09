# Vote-Go-Vue
 Exigences : installation du Node Js, Yarn, Docker et golang
Pour ce faire il faut suivre les étapes suivantes : 
1)	Il faut se rendre sur le répertoire client et en cliquant sur Shift + clique droite choisir lancer PowerShell ou CMD comme indiqué ci-dessous,Ensuite lancer la commander yarn install
2)	Après avoir lancer Yarn faut revenir à la racine du projet et faire la même manipulation pour ouvrir une 2ième fenêtre de commande et taper la commande docker-compose up -d si tout ce passe bien et pour s’assurer faut lancer la commande docker-compose ps et il faut obtenir un résultat de Client => Up et Db =>Up
3)	Toujours dans la racine lancer une 3ième fenêtre de commande et taper la commande suivante go run main.go faut être en écoute à un port pour que ca fonctionne si jamais ce n’est pas le cas veuillez vérifier l’antivirus 
4)	Et pour finir lancer le navigateur et lancer Localhost :8080 ou 127.0.0.1 :8080 pour démarrer l'application
