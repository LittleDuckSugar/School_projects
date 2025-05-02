// parsing du fichier changelog.md pour l'afficher sur la page blog

// parsing du fichier changelog.md dans la variable text
fetch('../md/CHANGELOG.md')
.then((response) => {
    return response.text();
})
.then((text) => {
    // conversion de la mise en page markdown en html
    parsedText = parseMarkdown(text);
    //regroupement des commits par date dans la page blog
    htmlText = getDate(parsedText)
    //affichage du changelog dans la balise output sur la page blog
    document.getElementById('output').innerHTML=htmlText;
});

function parseMarkdown(markdownText) {
    // remplacement de la syntaxe markdown par les balises html
    const htmlText = markdownText
    .replace(/^### (.*$)/gim, '<h3>$1</h3>')
    .replace(/^## (.*$)/gim, '<h3>$1</h3>')
    .replace(/^# (.*$)/gim, '<h1>$1</h1>')
    .replace(/^\> (.*$)/gim, '<blockquote>$1</blockquote>')
    .replace(/\*\*(.*)\*\*/gim, '<b>$1</b>')
    .replace(/\*(.*)\*/gim, '<i>$1</i>')
    .replace(/!\[(.*?)\]\((.*?)\)/gim, "<img alt='$1' src='$2' />")
    .replace(/\[(.*?)\]\((.*?)\)/gim, "<a href='$2'>$1</a>")
    .replace(/\n$/gim, '<br />') // do something to this markdown text
    .replace(/^date = ((.*$))/gim, '<h5>$1</h5>')

    return htmlText.trim();
}

function getDate(htmlText){
    let n1 = 1;
    let n2 = 1;
    var date = "";
    var nb =0
     while (n1 != -1 || n2 != -1){        
        //isolation de la variable de date a chaque commit
        n1 = htmlText.search("<h5>");
        n2 = htmlText.search("</h5>");

        var newdate = htmlText.slice(n1+4, n2);

        lenstr = htmlText.size;
        
        if (n1 != -1 || n2 != -1){
        // on vérifie que la date est différente de celle du commit précédent
            if (date == newdate){
                // si la date est identique alors on supprime juste la date du commit
                htmlText = htmlText.slice(0, n1) + htmlText.slice(n2+5, -1);
            } else {
                // si la date est différente de celle du commit précédent alors le commit vient d'un jour différent
                //alors on met en page la date pour qu'elle soit affichée en titre
                if (nb != 0){
                    htmlText = htmlText.slice(0, n1) + "</article>\n<article class = 'cardContent'>" + "<h2> Mise a jour du : " + newdate + "</h2>" + htmlText.slice(n2+5, -1);
                }else{
                    htmlText = htmlText.slice(0, n1) + "<article class = 'cardContent'>" + "<h2> Mise a jour du : " + newdate + "</h2>" + htmlText.slice(n2+5, -1);
                }
                nb++
            }
        }else{
            htmlText += "</article>";
        }
        // on sauvegarde la date pour vérifier le commit suivant
        date = newdate;
    }
    return htmlText;
}