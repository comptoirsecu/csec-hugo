<?php
# ffmpeg -i EP53_ffmpeg_vbr3.m4a -i chap.meta -map_chapters 1 -codec copy avec_chapitres.m4a
$content = file_get_contents("r5.chapters.txt");

$lines = array_reverse(explode(PHP_EOL, $content)); // De la fin au début pour construire les endtimes

$ch = array();

foreach($lines as $line) {
	# 00:10:20.078 raisons projets (suite)
	$blocks = explode(" ", $line, 2);
	
	$tcodes = explode(":", preg_replace("/\.\d{3}/","",$blocks[0])); // remove .225 ms
	$prevtime = isset($timecode) ? $timecode : 0;
	$timecode = 0;
	$timecode += $tcodes[0]*3600;
	$timecode += $tcodes[1]*60;
	$timecode += $tcodes[2];

	$chapter = $blocks[1];

// TIMEBASE met le compteur start/end en secondes
	$ch[] = <<<EOF

[CHAPTER]
TIMEBASE=1/1000
START=${timecode}000
END=${prevtime}000
title=${chapter}

EOF;
}

$ch[] = ";FFMETADATA1".PHP_EOL;

echo implode("", array_reverse($ch));
