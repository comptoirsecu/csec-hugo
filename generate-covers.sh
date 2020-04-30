#!/bin/bash
# Usage: 
#    --date=2020-04-23 to force a date and not use today's
#    --prefix=/path/to/csec-hugo folder if not run from that folder

PREFIX="." # no trailing slash
PICPATH="src/images/covers" # no trailing slash
DT=`date --rfc-3339=date`

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
	convert "${f}" -quality 60 -size 1920x1080 "${PREFIX}/static/images/covers/${DT}-${DESTFILE}.jpg"
	ls "${PREFIX}/static/images/covers/${DT}-${DESTFILE}.jpg"
	convert "${f}" -quality 60 -size 384x216 "${PREFIX}/static/images/covers/${DT}-${DESTFILE}-small.jpg"
	ls "${PREFIX}/static/images/covers/${DT}-${DESTFILE}-small.jpg"
	convert "${f}" -quality 60 -size 640x360 "${PREFIX}/static/images/covers/${DT}-${DESTFILE}-medium.jpg"
	ls "${PREFIX}/static/images/covers/${DT}-${DESTFILE}-medium.jpg"
	convert "${f}" -quality 60 -size 1024x576 "${PREFIX}/static/images/covers/${DT}-${DESTFILE}-large.jpg"
	ls "${PREFIX}/static/images/covers/${DT}-${DESTFILE}-large.jpg"
	convert "${f}" -quality 60 -size 640x360 "${PREFIX}/static/images/covers/${DT}-${DESTFILE}-medium.webp"
	ls "${PREFIX}/static/images/covers/${DT}-${DESTFILE}-medium.webp"
	convert "${f}" -quality 60 -size 1024x576 "${PREFIX}/static/images/covers/${DT}-${DESTFILE}-large.webp"
	ls "${PREFIX}/static/images/covers/${DT}-${DESTFILE}-large.webp"
done
