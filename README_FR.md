# TP

Un outil simple pour faciliter la création de solutions C# et la configuration du repo git.

Conçu pour être utilisé par les étudiants d'EPITA, mais pourrait fonctionner dans d'autres cas.

## Utilisation

Les étudiants d'EPITA doivent s'assurer de cloner le TP avant d'exécuter cet outil.

Une fois dans le répertoire créé par la commande clone, vous pouvez exécuter :

```bash
tp <nom du projet> [FLAGS]
```

```<nom du projet>``` sera le nom donné à la solution.

FLAGS :

    - ng : cela ne créera pas de fichier .gitignore
    - nr : cela ne créera pas de fichier README
    - ngr : ne créera ni .gitignore ni README
    - nrg : même chose que -ngr

Si vous rencontrez des erreurs ou avez des questions, n'hésitez pas à me contacter sur discord : thomas.lemago