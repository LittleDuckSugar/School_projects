from flask import Blueprint, render_template, request, redirect, url_for, flash
from app.models import db, Tag, Image, Category
import os
from werkzeug.utils import secure_filename
from PIL import Image as Picture
import pymysql
import math



image = Blueprint('image', __name__, url_prefix='/')


@image.route('/', methods=['GET', 'POST'])
def index():
    return redirect(url_for('image.importImg'))

@image.route('/import', methods=['GET', 'POST'])
def importImg():
    message_error = ""
    
    # recupére tout les tags
    tags = Tag.query.all()
    images = Image.query.all()

    if request.method == 'POST':
        typeImg = request.form.get('typeImg')
        isProduct = request.form.get('isProduct')
        isHuman = request.form.get('isHuman')
        isInstitutional = request.form.get('isInstitutional')
        credit = request.form.get('credit')
        rightOfUse = request.form.get('rightOfUse')
        copyrightImg = request.form.get('copyrightImg')
        endOfUse = request.form.get('endOfUse')
        category = request.form.get('category')
        filenames = request.files.getlist('filename')
        tagsChecked = request.form.getlist('tag')

        if isProduct == "oui" :
            isProduct = True
        else : 
            isProduct = False

        if isHuman == "oui" :
            isHuman = True
        else : 
            isHuman = False

        if isInstitutional == "oui" :
            isInstitutional = True
        else : 
            isInstitutional = False

        if rightOfUse == "oui" :
            rightOfUse = True
        else : 
            rightOfUse = False

        count_filename = 0
        for filename in filenames : 
            message_error = ""
            resp = uploadImage(filename)
            if not resp : 
                message_error = "Le type du fichier n'est pas sous le bon format"
            else :
                filenamefinal = str(filename.filename).replace(" ", "_").replace("é", "e").replace("è", "e").replace("à", "a").replace("â", "a").replace("ä", "a").replace("ê", "e").replace("ï", "i").replace("î", "i").replace("ñ", "n").replace("ô", "o").replace("ù", "u").replace(",", "_")

                for img in images :
                    if img.name == filenamefinal:
                        message_error = "cette image existe déjà dans la base de données"
                        count_filename += 1

                if message_error != "cette image existe déjà dans la base de données" :
                    formatImg = get_format_image(filenamefinal)

                    categoryImg = Category.query.filter_by(name=category).first()

                    image = Image(filenamefinal, typeImg, isProduct, isHuman, isInstitutional, formatImg , credit, rightOfUse, copyrightImg, endOfUse, categoryImg)
                    # db.session.add(image)

                    for tag in tagsChecked :
                        tagImg = Tag.query.filter_by(id=tag).first()

                        image.tags.append(tagImg)
                    db.session.add(image)
                    
                    db.session.commit()

        if len(filenames) == 1 :
            message_error = str(len(filenames) - count_filename) + " image a été ajouté et " + str(count_filename) + " existait déjà"
        else :   
            if count_filename == 1 :
                message_error = str(len(filenames) - count_filename) + " images ont été ajouté et " + str(count_filename) + " existait déjà sur les " + str(len(filenames)) + " images"
            else :  
                message_error = str(len(filenames) - count_filename) + " images ont été ajouté et " + str(count_filename) + " existaient déjà sur les " + str(len(filenames)) + " images"
        
        

    return render_template('pages/import.html', tags=tags,  message_error= message_error)




@image.route('/find', methods=['GET', 'POST'])
def findImg():
    tags = Tag.query.all()

    pagination = 1

    if request.method == 'POST':
        filename = request.form.get('filename')
        typeImg = request.form.get('typeImg')
        isProduct = request.form.get('isProduct')
        isHuman = request.form.get('isHuman')
        isInstitutional = request.form.get('isInstitutional')
        credit = request.form.get('credit')
        category = request.form.get('category')
        filenames = request.files.get('filename')
        formatImg = request.form.get('format')
        tagsChecked = request.form.getlist('tag')

        result = execute_requete_sql(filename,typeImg,isProduct,isHuman, isInstitutional, credit, formatImg, category, tagsChecked)

        images = []
        for r in result :
            images.append({'id': r[0], 'name': r[1]})

    
    else :
        images = Image.query.limit(28).all()

    return render_template('pages/find.html', tags=tags, images=images, pagination=pagination, lenImage=int((math.floor(len(images) /4))))



@image.route('/update/<int:id>', methods=['GET', 'POST'])
def updateImg(id):
    print(id)
    tags = Tag.query.all()
    image = Image.query.filter_by(id=id).first()

    tags_image = ""
    for tag in image.tags :
        tags_image += str(tag.name) + ","

    if request.method == 'POST':
        filename = request.form.get('filename')
        typeImg = request.form.get('typeImg')
        isProduct = request.form.get('isProduct')
        isHuman = request.form.get('isHuman')
        isInstitutional = request.form.get('isInstitutional')
        credit = request.form.get('credit')
        category = request.form.get('category')
        filenames = request.files.get('filename')
        formatImg = request.form.get('format')
        tagsChecked = request.form.getlist('tag')

        if isProduct == "oui" :
            isProduct = True
        else : 
            isProduct = False

        if isHuman == "oui" :
            isHuman = True
        else : 
            isHuman = False

        if isInstitutional == "oui" :
            isInstitutional = True
        else : 
            isInstitutional = False


        categoryImg = Category.query.filter_by(name=category).first()

        old_name = image.name

        image.name = filename
        image.typeImg = typeImg
        image.isProduct = isProduct
        image.isHuman = isHuman
        image.isInstitutional = isInstitutional
        image.credit = credit
        image.category = categoryImg

        
        if tagsChecked != [] :
            for tag in tagsChecked :
                tagImg = Tag.query.filter_by(id=tag).first()
                image.tags.append(tagImg)

        db.session.add(image)
        db.session.commit()

        os.rename('app/static/img/' + str(old_name), 'app/static/img/' + str(filename))
        return redirect(url_for('image.updateImg',id=id))



    return render_template('pages/update.html', tags=tags, image=image, tags_image=tags_image)


@image.route('/delete/<int:id>', methods=['GET', 'POST'])
def delete(id):
    image = Image.query.filter_by(id=id).first()
    image.tags = []
    db.session.add(image)
    db.session.commit()
    db.session.delete(image)

    os.remove('app/static/img/' + str(image.name))
    db.session.commit()

    return redirect(url_for('image.findImg'))


@image.route('/addtag', methods=['GET', 'POST'])
def addTag():
    message = ""
    if request.method == 'POST':
        tag = request.form.get('tag')
        tag = Tag(tag)
        db.session.add(tag)
        db.session.commit()
        message = "vous avez ajouté un tag"


    return render_template('pages/addtag.html', message=message)



def execute_requete_sql(filename, typeImg,isProduct,isHuman, isInstitutional, credit, formatImg, category, tagsChecked) :
    """ Fonction qui execute la requete sql pour filtrer les images

    Args :
        all [string] : les differentes colonnes en bdd
    
    Return [list]
    """
    requetes = get_requete_sql_image(typeImg,isProduct,isHuman, isInstitutional, credit, formatImg, category, tagsChecked)
    db = pymysql.connect("localhost", "root", "", "pomona")
    cursor = db.cursor()

    select = "select * from image "

    where = "where image.name LIKE '%{}%' ".format(filename)
    for requete in requetes :
        where += requete + " "

    join = ""
    if category != "" :
        join += "join category on image.category_id = category.id "
    
    if tagsChecked != [] :
        join += "join image_tag on image.id = image_tag.image_id "

    order = "order by image.name limit 28"

    sql = select + join + where + order
    print(sql)

    cursor.execute(sql)
    result = cursor.fetchall()

    return result




def get_requete_sql_image(typeImg,isProduct,isHuman, isInstitutional, credit, formatImg, category, tagsChecked):
    """ Fonction qui recupere la requete sql si les inputs sont rempli ou non

    Args :
        all [string] : les differentes colonnes en bdd
    
    Return [list]
    """
    # on stocke dans requetes tout les "and name =" si l'input n'est pas vide, si il est vide on ne veut pas filtrer par cet input
    requetes = []
    if typeImg != "" :
        requetes.append("and typeImg = '" + str(typeImg) + "'")
    
    if isProduct != "" :
        if isProduct == "oui" :
            requetes.append("and isProduct = " + str(True))
        else : 
            requetes.append("and isProduct = " + str(False))

    if isHuman != "" :
        if isHuman == "oui" :
            requetes.append("and isHuman = " + str(True))
        else : 
            requetes.append("and isHuman = " + str(False))

    if isInstitutional != "" :
        if isInstitutional == "oui" :
            requetes.append("and isInstitutional = " + str(True))
        else : 
            requetes.append("and isInstitutional = " + str(False))
    
    if credit != "" :
        requetes.append("and credit = '" + str(credit) + "'")

    if formatImg != "" :
        requetes.append("and formatImg = '" + str(formatImg) + "'")

    if category != "" :
        cat = Category.query.filter_by(name=category).first()
        requetes.append("and category_id = '" + str(cat.id) + "'")

    if tagsChecked != [] :
        for tag in tagsChecked :
            requetes.append("and tag_id = '" + str(tag) + "'")

    return requetes

# select * from image 
# join category on image.category_id = category.id
# where category_id = 1 order by image.name




def get_format_image(filename):
    """ Fonction qui recupere le format de l'image (horizontal ou vertical)

    Args:
        filename [string] : nom de l'image
    
    Return [string] : horizontal ou vertical
    """
    imagename = "app/static/img/" + str(filename)
    image = Picture.open(imagename)

    width,height=image.size  

    if width >= height :
        return "horizontal"
    return "vertical"

def allowed_image(filename):
    """ Fonction qui permet de verifier si l'extension de l'image est bonne

    Args:
        filename [string] : nom de l'image
    
    Return [Boolean]
    """

    # Nous voulons uniquement des fichiers avec un. dans le nom de fichier
    if not "." in filename:
        return False

    # Séparez l'extension du nom de fichier
    ext = filename.rsplit(".", 1)[1]

    # Vérifiez si l'extension est dans ALLOWED_IMAGE_EXTENSIONS
    if ext.upper() in ["JPEG", "JPG", "PNG"]:
        return True
    else:
        return False




def uploadImage(image):
    """ Fonction qui upload une image dans le bon dossier 

    Args:
        image [request.files] : image
    
    Return [Boolean]
    """
    if allowed_image(image.filename):
        if image.mimetype == 'image/png' or image.mimetype == 'image/jpg' or image.mimetype == 'image/jpeg':

            filename = secure_filename(image.filename)
            uploads_dir = 'app/static/img/'
            os.makedirs(uploads_dir, exist_ok=True)
            image.save(os.path.join(uploads_dir, filename))

            return True

    return False



