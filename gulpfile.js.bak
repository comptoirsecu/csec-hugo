var gulp = require('gulp');
var plugins = require('gulp-load-plugins')();
var img = {
  cover: {   src: "src/images/covers/",
            dst: "static/images/covers/" },
  misc: {   src: "src/images/misc/",
            dst: "static/images/misc/" },
  thumbnail: {   src: "src/images/thumbnails/",
            dst: "static/images/thumbnails/" },
  };
var exec = require('child_process').exec;
var runSequence = require('run-sequence');
gulp.task('img', function(callback) {
  runSequence(['img:covers', 'img:thumbnails', 'img:misc'], ['img:covers:clean', 'img:thumbnails:clean', 'img:misc:clean'], callback);
})

gulp.task('img:covers:clean', function() {
  return gulp.src("src/images/covers/*.{jpg,jpeg,gif,png,webp}")
        .pipe(plugins.clean())
        .pipe(gulp.dest("src/images/covers/archive/"))
})

gulp.task('img:misc:clean', function() {
  return gulp.src("src/images/misc/*.{jpg,jpeg,gif,png,webp}")
        .pipe(plugins.clean())
        .pipe(gulp.dest("src/images/misc/archive/"))
})

gulp.task('img:thumbnails:clean', function() {
  return gulp.src("src/images/thumbnails/*.{jpg,jpeg,gif,png,webp}")
        .pipe(plugins.clean())
        .pipe(gulp.dest("src/images/thumbnails/archive/"))
})

gulp.task('img:misc:gif', function () {
  return gulp.src("src/images/misc/*.gif").pipe(gulp.dest(img.misc.dst));
})

gulp.task('img:misc', ['img:misc:gif'], function() {
  return  gulp.src(img.misc.src + "*.{png,jpg,jpeg}")
  .pipe(plugins.responsive({'*.{jpg,jpeg,png}':
      {  width: 1000, rename:
          { extname: '.jpg',
          prefix: new Date().toISOString().split('T')[0] + '-'
          }
      }},
  {
    withMetadata: false,
    background: "#ECF0F1",
    flatten: true,
    errorOnEnlargement: false,
    errorOnUnusedConfig: false,
  }))
  .pipe(gulp.dest(img.misc.dst));
});

gulp.task('img:thumbnails', function() {
  return gulp.src(img.thumbnail.src + "*.{png,jpg,jpeg,gif}")
  .pipe(plugins.responsive({'*': [
      { quality: 40, width: 200, height: 200, rename:
          { extname: '.jpg' } },
      { quality: 40, width: 400, height: 400, rename:
          { extname: '-@2x.jpg' } }]},
  {
    withMetadata: false,
    embed: true,
    flatten: true,
    background: "#ECF0F1",
    errorOnEnlargement: false,
    withoutEnlargement: false,
    errorOnUnusedConfig: false,
  }))
  .pipe(gulp.dest(img.thumbnail.dst));
});


gulp.task('img:covers', function() {
  return gulp.src(img.cover.src + "*.{png,jpg,jpeg,gif}")
  .pipe(plugins.responsive({
    '*': [{
        quality: 60,
        width: 1920,
        height: 1080,
        rename: {
          extname: '.jpg',
          prefix: new Date().toISOString().split('T')[0] + '-'
        }
    },{
      width: 384,
      height: 216,
      rename: {
        //prefix: new Date().toISOString().split('T')[0] + '-',
        suffix: '-small',
        extname: '.jpg',
        prefix: new Date().toISOString().split('T')[0] + '-'
      },
      // format option can be omitted because
      // format of output image is detected from new filename
      // format: 'jpeg'
      // Do not enlarge the output image if the input image are already less than the required dimensions.
      progressive: true,
      quality: 60,
    }, {
      width: 640,
      height: 360,
      rename: {
        //prefix: new Date().toISOString().split('T')[0] + '-',
        suffix: '-medium',
        extname: '.jpg',
        prefix: new Date().toISOString().split('T')[0] + '-'
      },
      progressive: true,
      quality: 60,
    }, {
      width: 1024,
      height: 576,
      rename: {
        //prefix: new Date().toISOString().split('T')[0] + '-',
        suffix: '-large',
        extname: '.jpg',
        prefix: new Date().toISOString().split('T')[0] + '-'
      },
      progressive: true,
      quality: 60,
    }, {
      width: 640,
      height: 360,
      rename: {
        //prefix: new Date().toISOString().split('T')[0] + '-',
        suffix: '-medium',
        extname: '.webp',
        prefix: new Date().toISOString().split('T')[0] + '-'
      },
      quality: 50,
    }, {
      width: 1024,
      //height: 379, //576,
      rename: {
        //prefix: new Date().toISOString().split('T')[0] + '-',
        suffix: '-large',
        extname: '.webp',
        prefix: new Date().toISOString().split('T')[0] + '-'
      },
      quality: 50,
    }],
    }, {
    // Global configuration for all images
    // Strip all metadata
    withMetadata: false,
    crop: 'attention',
    flatten: true,
    background: "#0a0a0a",
    // Do not emit the error when image is enlarged.
    errorOnEnlargement: false,
    errorOnUnusedConfig: false,
  }))
  .pipe(gulp.dest(img.cover.dst));
});


gulp.task('lint:html', function () {
  return gulp.src(['content/**/*.html', 'layout/**/*.html', 'public/**/*.html', 'archetypes/**/*.html'])
  .pipe(plugins.htmlhint())
  .pipe(plugins.htmlhint.failReporter());
});

gulp.task('clean', function() {
  return gulp.src('public')
  .pipe(plugins.clean());
});


gulp.task('hugo', function (cb) {
  exec('hugo --gc --minify --cleanDestinationDir', function (err, stdout, stderr) {
    console.log(stdout);
    console.log(stderr);
    cb(err);
  });
});


gulp.task('default', ['build'])
gulp.task('build', ['hugo']);
