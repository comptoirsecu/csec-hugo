#!/bin/bash
# Usage: 
#    --date=2020-04-23 to force a date and not use today's
#    --prefix=/path/to/csec-hugo folder if not run from that folder

PREFIX="." # no trailing slash
PICPATH="src/images/thumbnails" # no trailing slash

for i in "$@"
do
case $i in
    -p=*|--prefix=*)
    PREFIX="${i#*=}"

    ;;
    -s=*|--date=*)
    DT="${i#*=}"

    ;;
esac
done

# Prepend prefix
PICPATH=${PREFIX}/${PICPATH}

for f in `ls ${PICPATH}/*.jpg`
do
	DESTFILE=`basename -- $f` 
	DESTFILE=${DESTFILE%.*}
	convert "${f}" -quality 40 -resize 200x200 "${PREFIX}/static/images/thumbnails/${DESTFILE}.jpg"
	ls "${PREFIX}/static/images/thumbnails/${DESTFILE}.jpg"
	convert "${f}" -quality 40 -resize 400x400 "${PREFIX}/static/images/thumbnails/${DESTFILE}-@2x.jpg"
	ls "${PREFIX}/static/images/thumbnails/${DESTFILE}-@2x.jpg"
done
