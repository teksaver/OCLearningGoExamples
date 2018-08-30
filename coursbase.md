# Apprenez à programmer en Go par la pratique
[//]: # (Un modèle fondé sur Go, à adapter pour chaque langage)

[//]: # (Certaines parties du cours a vocation à être identique pour tous les langages : n’hésitez pas à commenter pour proposer des modifications, dans une volonté d’amélioration continue)

[//]: # (Les parties à adapter en fonction du langage comprennent les explications spécifiques, les exemples de code et les exercices codevolve )

## Chapitre de présentation (objectifs, méthodologie, prérequis)

Dans ce cours, vous allez tout d’abord faire une plongée vertigineuse et implémenter, pas à pas, un premier programme.. Puis nous allons, ensemble, revoir les éléments essentiels à tout développement informatique : les structures de données, les fonctions et le contrôle du flux.

Prêt à plonger ? c’est parti !

> **:information_source:** 
> Ce cours s’appuie sur le langage Go. Vous pouvez également le suivre en C#, Kotlin, Java, JavaScript, PHP, Python…  

# Partie 1 Implémentez une TODO LIST

## Chapitre 1 : Modélisez le besoin

Faire ses courses,  apprendre à coder, passer voir un ami.. penser à tout pendant une journée bien chargée nécessite une organisation formidable ! Pour ne rien oublier, rien de tel qu’une TODO-LIST. 

![Texte alternatif](./todolist.png "Légende de l'image")


Conceptuellement, une TODO-LIST est un simple ensemble ordonné de tâches à ne pas oublier. En tant qu’utilisateur, nous avons besoin d’ajouter des tâches, de les supprimer, de les modifier… et de les visualiser. Pour cela, un stylo, un  post-it, une feuille volante, tous les moyens sont bons. Et pourquoi pas un petit programme ? 

Si vous avez suivi notre cours sur la logique informatique, vous savez qu’un programme informatique peut s’appuyer sur un concept que nous connaissons tous : un objet !

(boite bleue) Nous ferons régulièrement référence à des concepts du cours de logique informatique. Vous pouvez faire une pause sur celui-ci et aller le suivre, ou lire les chapitres recommandés au fur et à mesure.

A partir du besoin exprimé précédemment, nous pouvons définir les opérations  suivantes :
Ajouter une tâche à la liste,
Supprimer une tâche,
Modifier une tâche,
Visualiser l’ensemble des tâches

Afin de conserver l’information, il nous faut la stocker dans un type de données adapté. Chaque tâche pouvant être décrite par un simple texte, nous allons donc utiliser, pour chaque tâche, ce que nous appelons une chaîne de caractères.

Enfin, comme nous avons plusieurs tâches, nous allons les regrouper dans une structure de données adaptée : idéalement, une liste. 

Avant de passer au code, résumons rapidement notre besoin à l’aide d’un diagramme UML. Ce type de diagramme permet de modéliser rapidement un besoin de manière visuelle et dans un format connu de tous les développeurs.. idéal pour échanger sur une idée !




Un diagramme UML dit de classes modélise un objet en 3 parties : son nom, ses propriétés et ses opérations. Voici une modélisation visuelle de notre objet TodoList : 

TodoList
-taches: liste de chaines de caractères
+ajouter(tache)
+supprimer(tache)
+modifier(tache)
+visualiser()


Passons maintenant à la pratique !

Chapitre 2 : Exécutez votre premier programme
Comme nous l’avons vu au chapitre précédent, modéliser un objet permet de définir les opérations qu’il est capable de réaliser. Un programme, lui, vise à définir la séquence dans laquelle ces opérations vont avoir lieu. 
Pour notre premier programme, nous allons mettre en oeuvre la séquence suivante :
Créer une liste vide
Ajouter 4 tâches à la liste : se lever, prendre une douche, petit-déjeuner et partir au travail
Afficher le contenu de la liste
Supprimer l’étape du petit déjeuner 
Afficher le contenu de la liste

Voyons comment réaliser ce programme en Go :

Indiquez votre intention avec des commentaires
La première chose que nous pouvons écrire dans un programme est un élément qui ne sert pas à l’ordinateur, mais plutôt à un collègue développeur cherchant à comprendre le programme - très souvent, vous-même dans quelques semaines ! Il s’agit des commentaires. Afin de vous préparer pour l’entreprise et/ou à devenir contributeur OpenSource, nous allons écrire ces commentaires, ainsi que le reste du code, en anglais

[info bulle] l’anglais est le langage universel du développeurs. Les mots clés de la plupart des langages sont issus de l’anglais. Ecrire en anglais vous permet également de partager vos codes sur les forums d’entraide comme StackOverflow. Ne négligez pas cette possibilité.


En Go, un commentaire peut s’écrire en préfixant la phrase par //. 

Afin de n’oublier aucune étape, nous pouvons ainsi commencer par écrire ceci : 

    // Create an empty list of tasks named myTodoList
    // Add all 4 tasks to the list
    // Display tasks in the list
    // Remove breakfast
    // Display tasks after removal

Ecrivons maintenant le code correspondant à chaque étape : 

Créez une liste vide
Comme nous l’avons vu lors du chapitre précédent, il est nécessaire de créer une structure de données pour enregistrer les différentes tâches. Idéalement, cette structure doit être :
ordonnée : les éléments doivent s’ajouter à la fin de la liste
dynamique : on souhaite pouvoir ajouter autant d’éléments que nécessaire 
typée : on souhaite restreindre le type d’éléments que l’on peut ajouter à la liste. Ici, on ne stocke que des tâches, représentées par des chaînes de caractères.

Go fournit une structure respectant ces 3 contraintes : le slice. Voici le code permettant de créer un slice vide :

    // Create an empty list of tasks named myTodoList
    var myTodoList []string

Ce code comporte 3 éléments :
var est un mot-clé du langage indiquant que nous allons définir une variable. Une variable se définit comme un symbole qui associe un nom à une valeur, stockée quelque part dans la mémoire de l’ordinateur.
myTodoList est le nom donné à la variable. Ce nom permet d’accéder à la mémoire.
[]string est le type de la variable. Comme son nom l’indique, il définit le type de données que l’on peut stocker. Ici,
 string indique que ce que l’on manipule est une chaîne de caractères.
 [] signifie que nous avons bien un slice capable de stocker plusieurs chaînes

Nous avons désormais à notre disposition un slice prêt à accueillir nos tâches. 

Ajoutez 4 tâches à la liste

 // Add 4 tasks to the list
    myTodoList = append(myTodoList, "Wake up", "Shower", "Have breakfast", "Go to work")

Cette seconde ligne de code est constituée des éléments suivants :
append est une fonction du langage permettant d’ajouter un ou plusieurs éléments à une liste
Le premier argument de la fonction est le nom du slice auquel il faut ajouter un ou plusieurs éléments
[info bulle] Les arguments sont les éléments situés entre () après le nom de la fonction et séparés par des virgules. Ils fournissent les éléments nécessaires au bon fonctionnement de la fonction.
Les 4 arguments suivants sont des chaines de caractères. Une chaine de caractère est ainsi représentée par les lettres entourées de “ “. 
= est l’opérateur d’affectation. Il permet de stocker dans la variable à gauche de = le résultat de l’appel de la fonction append. Ici nous remplaçons la valeur actuelle de myTodoList par la liste initiale (vide) à laquelle ont été ajoutés les 4 nouvelles tâches

Affichez le contenu de la liste


    // Display tasks in the list
    displayListOfString(myTodoList)

Cette 3ème instruction appelle une fonction d'affichage en lui indiquant que la chaîne de caractères à afficher est myTodoList. 

L’affichage d’une liste nécessite en effet un ensemble d’opérations, qui ont été codées par ailleurs dans la fonction myTodoList : c’est le principe de la boite noire du cours de logique. A partir du moment où une fonction vous est fournie ou que vous l’avez codée, elle fait partie de votre boîte à outils. Vous n’avez qu’à lui fournir les entrées que la fonction attend pour obtenir la sortie qu’elle calcule.

[bulle info] vous pouvez consulter l’implémentation de cette fonction dans l’exercice qui vous est proposé en fin de chapitre

Supprimez l’étape du petit-déjeuner
Cette étape n’a absolument pas l’objectif d’être implémenté dans la vie réelle. L’ordinateur va pourtant l’exécuter sans sourciller, ce qui montre que la puissance ne fait pas tout.

L’opération de suppression dans une liste peut être très lourde en temps de calcul. Il faut en effet décaler tous les éléments qui viennent après l’élément à supprimer afin de conserver l’ordre de la liste, comme l’illustre le schéma ci-après (à refaire)









La plupart des langages qui proposent une structure de liste vous fournissent des fonctions prédéfinies 


Ce n’est malheureusement pas le cas en Go : pour rappel, nous utilisons une structure appelée slice, qui ne fournit pas toutes les opérations courantes pour une liste.

[infobulle] La raison pour cela est justement la complexité en temps de calcul de cette opération. Le langage vous encourage à utiliser une structure plus adaptée si vous avez 1) beaucoup de données et 2) beaucoup d’opérations de suppression à prévoir

Le code de suppression a donc été écrit pour vous dans une fonction removeFromSliceByIndex qui supprime l’élément voulu en fonction de sa position. Il nous reste à appeler cette fonction avec les bons arguments

    // Remove breakfast
    myTodoList = removeFromSliceByIndex(myTodoList, 2)

Le premier argument est le slice dans lequel nous souhaitons supprimer un élément
Le second est l’index de l’élément à supprimer. Comme dans la majorité des langages, cet index commence à 0 : l’index du petit-déjeuner qui est le 3ème élément est donc 2.
[infobulle] L’indexation à 0 provient du langage C : un tableau est représenté par un ensemble de cases mémoires contiguës. L’index indique de combien de cases il faut se décaler en partant du premier élément. Le premier élément est donc bien indexé à 0. 



L’étape suivant vérifie le bon fonctionnement du programme en affichant la liste après suppression

Affichez la liste après suppression
Le code est identique au précédent, seul le résultat change

    // Display tasks after removal
    displayListOfString(myTodoList)

Une fois lancé, ce programme va afficher :
Here is the list
Task N° 0  is  Wake up
Task N° 1  is  Shower
Task N° 2  is  Have breakfast
Task N° 3  is  Go to work

Here is the list
Task N° 0  is  Wake up
Task N° 1  is  Shower
Task N° 2  is  Go to work

Modifiez une tache
Pour terminer, nous allons utiliser le temps gagné par la suppression d’une tâche pour en modifier une : nous allons remplacer la douche par un bain

    // change shower to bath
    myTodoList[1] = "Take Bath"


Exécutez et modifiez le code
L’exercice ci-dessous va vous permettre de pratiquer directement dans votre navigateur. Vous allez pouvoir
Exécutez le code fourni et voir le résultat de l’exécution
Vous exercer en apportant les modification suivantes :
Afficher la liste finale après la dernière modification
Modifier le programme pour ajouter 3 nouvelles tâches

[codevolve]

Maintenant que vous avez exécuté votre premier programme, il ne nous reste plus qu’à résumer ce que nous avons appris

Résumé
Un programme met en oeuvre une séquence d’instructions exécutée par la machine, mais doit également être compréhensible par d’autres développeurs. Pour cela, nous avons dans ce chapitre ajouté des commentaires expliquant l’objectif de chaque partie du code.

Les commentaires permettent d’expliquer une intention à l’instant T, mais l’expérience montre qu’ils ont tendance à diverger du code avec le temps et la complexité. Dans une base de code complexe il n’est pas rare que les commentaires, au lieu d’apporter l’aide attendue, induisent en erreur par défaut de mise-à-jour.

Vous aurez également noté que ce programme est finalement assez éloigné de la modélisation que nous avons conçue lors du premier chapitre.

Le chapitre suivant montre comment nous pouvons organiser notre code afin qu’il constitue lui-même la meilleure des documentations.

Chapitre 2 : Améliorez la lisibilité de votre code en définissant un objet
Le code du chapitre précédent possède l’avantage d’être fonctionnel. Cependant, il dépend largement de la qualité de la documentation pour être compréhensible.

Etudions maintenant le code Go suivant :


    var myTodoList TodoList
    myTodoList.Add("Wake up")
    myTodoList.Add("Shower")
    myTodoList.Add("Have breakfast")
    myTodoList.Add("Go to work")
    myTodoList.Display()
    myTodoList.Remove(2)
    myTodoList.Display()
    myTodoList.Rename(1, "Take bath")
    myTodoList.Display()


Ce code n’est pas commenté, mais le nom des opérations indique clairement ce qu’il se passe. 
Nous déclarons une variable appelée myTodoList de type TodoList. 
Nous envoyons des messages à cette variable pour réaliser une opération.

Par exemple,     myTodoList.Add("Go to work") envoie le message Add(“Go to work”) à myTodoList., ce qui a l’effet attendu d’ajouter la tâche Go To Work à la liste.

[questionbulle] Mais qu’est-ce qui permet à myTodoList d’accepter ces messages ?

Il s’avère que myTodoList est un objet créé à partir du modèle TodoList. En programmation objet, on appelle généralement ce modèle une classe. En Go, ce modèle est défini par un type de données appelé struct. Nous utiliserons indifféremment le terme classe ou structure dans la suite de ce chapitre.

Afin de permettre le fonctionnement de notre programme, nous allons implémenter l’équivalent du schéma UML défini au chapitre 1, en prenant soin de renommer les différents éléments pour correspondre aux noms utilisés dans notre programme.

TodoList
-tasks: list of string
+Add(task:string)
+Display()
+Remove(task:string)
+Rename(task:string)



Comme nous pouvons le voir, nous devons définir
Le nom de la structure : ici, TodoList
Une structure de données pour enregistrer l’état de la TodoList : idéalement, une liste de chaînes de caractères
Les opérations d’ajout, suppression, modification et visualisation

Nous allons dans un premier temps créer la structure et son contenu, puis implémenter chaque opération.

Définissez le nom et les champs
En Go, nous pouvons écrire le code suivant :
    // TodoList  models  common operations on a Todo List   
type TodoList struct {
    tasks []string
}

En Go, le commentaire placé avant une structure est recommandé : il permet de générer une documentation pour les développeurs souhaitant réutiliser ce code.

[infobulle] Un point fort de la programmation objet est qu’elle permet de définir des
 composants réutilisables. Un ensemble de classes mis à disposition d’autres développeurs est appelé une librairie.

La structure contient le seul élément utile pour le besoin modélisé : la liste des tâches implémentée sous la forme d’un slice de chaînes de caractères. 

En Go, une variable définie dans une structure est appelé un champ.

La suite du code permet d’implémenter les différentes opérations sous la forme de méthodes
Implémentez la méthode d’ajout
Lorsque nous avons écrit notre premier programme, nous avons utilisé des fonctions : ces fonctions ingurgitent en entrée les arguments nécessaires, effectuent une séquence d’opération puis affichent ou retournent un résultat.

Dans une classe, nous définissons des méthodes : il s’agit de fonctions qui sont rattachées à l’objet : elles ont donc accès directement aux champs de l’objet. 

Nous pouvons donc écrire :

// Add a string at the end of the slice
func (t *TodoList) Add(value string) {
    t.tasks = append(t.tasks, value)
}

Vous ne comprenez pas nécessairement tout ce code, et c’est normal ! l’important ici est de se rendre compte que la méthode Add n’a besoin que d’un seul argument : la valeur à ajouter. Le champ tasks est accessible directement depuis l’intérieur de la méthode. 

Dans la notation t.tasks,  t est une référence vers l’objet. Elle représente l’objet réel qui sera construit à partir du modèle défini par la structure.

Le fonctionnement est identique au chapitre précédent : nous utilisons la fonction append pour créer une nouvelle liste à partir de la liste existante et de la chaîne à ajouter et nous affectons le résultat à la t.tasks. 

A chaque appel de la méthode Add, la chaîne de caractères fournie en argument de la méthode est donc ajoutée à l’objet issu de TodoList.

Voyons maintenant la méthode permettant de visualisation du contenu.

Implémentez la méthode de visualisation 
Pour visualiser le contenu d’une liste, il est nécessaire de parcourir l’ensemble de la liste et de l’afficher élément par élément.

En Go, nous commençons par un commentaire de documentation puis nous définissons la méthode

  // Display prints all the tasks of the list
func (t TodoList) Display() {

La méthode ayant accès au champ contenant la liste, il n’est pas nécessaire de définir un paramètre. Il n’y a qu’à écrire le code permettant le parcours de celle-ci. Dans un premier temps, nous affichons un message. 

    fmt.Println("Here is the list")

Puis nous écrivons le code du parcours

    for index, task := range t.tasks {
   	 fmt.Println("Task N°", index, " is ", task)
    }

La construction for … range permet de répéter l’opération tant qu’il y a des éléments dans la variable après le range, issu t.tasks.

L’index indique la position de l’élément tandis que task contient la chaîne décrivant la tâche. 

Nous reviendrons sur les structures de répétitions dans la 4ème partie de ce cours.

Implémentez la suppression


Implémentez la modification


Pratiquez !
Dans l’exercice CodeVolve ci-dessous, vous allez
Faire tourner l’exemple présenté dans ce chapitre
Ecrire le programme permettant d’utiliser un objet créé à partir du modèle suivant :

Team
-players: list of string
+Add(task:string)
+Display()
+SetNumber(number:int)



[lien codevolve]

Il est désormais temps de clore ce chapitre avec un résumé de nos acquis.


Résumé
Un programme orienté objet permet d’apporter de la lisibilité. Il s’écrit en 2 étapes
Ecriture d’un modèle pour l’objet définissant ses propriétés et ses opérations. 
La séquence d’opérations est réalisée par la création d’un objet puis l’envoi de messages à celui-ci

Dans le chapitre suivant, nous allons apporter une fonctionnalité supplémentaire : marquer les tâches qui sont terminées.
Chapitre 3 : Indiquez les tâches qui sont terminées grâce à une structure de données adaptée
⇒ map/dictionnaire au lieu de la liste
Chapitre 4 : Gérez plus de fonctionnalités avec les objets
⇒ diagramme UML avec l’association TodoList --------*> Tasks où Task est désormais un objet
Chapitre 5 : Pensez interface et comportement

Quelle que soit la structure de données choisie, le comportement de base reste le même : nous voulons ajouter, modifier, supprimer et afficher.
De la même manière que nous avons modélisé la tâche comme un objet doté d’opérations, nous pouvons modéliser notre liste comme une implémentation d’une interface. Cela permettra à n’importe quelle application utilisant une liste, quelle que soit son type, d’utiliser les opérations communes.

En Go, la définition d’opérations commune se fait par la déclaration d’une interface

Si nous voulons vraiment

Nous allons commencer par définir nos structures de données. Il nous faut un moyen d’enregistrer
l’objet global, TodoList
la liste de tâches, taches


En Go, il va nous falloir définir une structure qui définit un modèle abstrait à partir duquel nous pourrons créer notre objet. 

[bulle info] dans beaucoup de langages la structure définissant un modèle s’appelle une classe. 


Partie 2 : Structurez vos données de manière pertinente
Chapitre 2.1 : Démystifiez la notion de variable
En mathématiques et en logique, on définit une variable comme un ”symbole utilisé pour marquer un rôle dans un prédicat, une formule ou un algorithme”.


Comme nous l’avons vu dans le cours de logique informatique, elle est constituée de 3 parties :

Un nom
Un type
Une valeur

Si l’on reprend la définition, un point fondamental est qu’une variable marque un rôle. Un enjeu clé est donc de lui donner un nom représentatif.




Le nom est 

Combien font  a + 7 lorsque a vaut 4 ?

Une variable possède


selon les langages, il est nécessaire de la déclarer 


 Chapitre 2.2 : Découvrez le typage des variables
Nous avons vu au chapitre précédent qu’une composante fondamentale d’une variable était son type. Certains langages de programmation vous demandent de spécifier ce type, tandis que d’autres vont le deviner automatiquement en fonction de la valeur que vous allez affecter à la variable. Vous entendrez ainsi parler de typage statique ou dynamique, ainsi que de typage fort ou faible, et implicite/explicite.

Ce choix dépend généralement de la manière dont le langage est

Le cas le plus contraignant est le typage statique : dans ce cas, le type de la variable est fixé une fois pour toute. Toute modification du type nécessite une conversion explicite et l’affectation du résultat dans une nouvelle variable.

De l’autre côté du spectre, le typage dynamique faible laisse une liberté totale au programmeur : la valeur peut évoluer sans contrainte tout au long du cycle de vie. 




Un typage statique fournit une contrainte à la variable tandis qu’un typage dynamique contraint 




Static/Dynamic Typing is about when type information is acquired (Either at compile time or at runtime)
Strong/Weak Typing is about how strictly types are distinguished (e.g. whether the language tries to do an implicit conversion from strings to numbers).
ex : https://stackoverflow.com/questions/11328920/is-python-strongly-typed


Typahe
Statique
Dynamique
Faible


JavaScript, PHP
Fort
Go, Java, Scala
Python









Partie3 : Gérez le flux de votre programme

Chapitre 3.1 : Sélectionnez la condition appropriée
Chapitres 3.2 : Utilisez la bonne structure de répétition
Chapitre 3.3: Gérez les erreurs / exceptions
Chapitre 3.4 : Découvrez la complexité

Partie 4 : Modularisez votre code avec les fonctions
Chapitre 3.1 : Découvrez les fonctions systèmes
Chapitre 3.2 : le point d’entrée du programme et ses arguments
Chapitre 3.3 : Testez qu’une fonction répond à votre besoin
Chapitre 4.4 : Écrivez votre propre fonction
Chapitre 4.5 : Deboguez votre fonction
