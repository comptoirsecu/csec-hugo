var gulp = require('gulp');
var plugins = require('gulp-load-plugins')();
var css_dest = "static/css";
var sass_src = "src/scss/app.scss";
var runSequence = require('run-sequence');
var javascript_src = ['src/js/enabled/foundation.core.js', 'src/js/enabled/foundation.util.*.js', 'src/js/enabled/*.js'];
var javascript_dest = "static/js";
var img = {
  cover: {   src: "src/images/covers/*.{png,jpg,jpeg,gif}",
            dst: "static/images/covers/" },
  misc: {   src: "src/images/misc/*.{png,jpg,jpeg,gif}",
            dst: "static/images/misc/" },
  thumbnail: {   src: "src/images/thumbnails/*.{png,jpg,jpeg,gif}",
            dst: "static/images/thumbnails/" },
  };

var sassPaths = [
  'src/scss/foundation',
  'src/scss/normalize-scss/sass',
  'src/scss/sassy-lists/stylesheets'
];
var exec = require('child_process').exec;



gulp.task('sass:dev', function() {
  return gulp.src(sass_src)
    .pipe(plugins.sass({
      includePaths: sassPaths,
      outputStyle: 'nested' // if css 'compressed' **file size**
    })
      .on('error', plugins.sass.logError))
    .pipe(plugins.autoprefixer({
      browsers: ['last 2 versions', 'ie >= 9']
    }))
    .pipe(gulp.dest(css_dest));
});

gulp.task('sass:prod', function() {
  return gulp.src(sass_src)
    .pipe(plugins.sass({
      includePaths: sassPaths,
      outputStyle: 'compressed' // if css 'compressed' **file size**
    })
      .on('error', plugins.sass.logError))
    .pipe(plugins.autoprefixer({
      browsers: ['last 2 versions', 'ie >= 9']
    }))
    .pipe(gulp.dest(css_dest));
});

gulp.task('img:misc', function() {
  return gulp.src(img.misc.src)
  .pipe(plugins.responsive({'*': [
      { quality: 50, width: 400, rename:
          { extname: '.jpg', suffix: '-small'} },
      { quality: 50, width: 200, rename:
          { extname: '.jpg'} }]},
  {
    withMetadata: false,
    embed: true,
    background: "#fff",
    errorOnEnlargement: false,
    errorOnUnusedConfig: false,
  }))
  .pipe(gulp.dest(img.misc.dst));
});

gulp.task('img:thumbnails', function() {
  return gulp.src(img.thumbnail.src)
  .pipe(plugins.responsive({'*': [
      { quality: 50, width: 400, height: 400, rename:
          { extname: '.jpg', suffix: '-small'} },
      { quality: 50, width: 200, height: 200, rename:
          { extname: '.jpg'} }]},
  {
    withMetadata: false,
    crop: true,
    background: "#fff",
    errorOnEnlargement: false,
    errorOnUnusedConfig: false,
  }))
  .pipe(gulp.dest(img.thumbnail.dst));
});



gulp.task('img:covers', function() {
  return gulp.src(img.cover.src)
  .pipe(plugins.responsive({
    '*': [{
        quality: 60,
        width: 1920,
        height: 1080,
        rename: {
          extname: '.jpg',
          //prefix: new Date().toISOString().split('T')[0] + '-'
        }
    },{
      width: 384,
      height: 216,
      rename: {
        //prefix: new Date().toISOString().split('T')[0] + '-',
        suffix: '-small',
        extname: '.jpg',
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
      },
      quality: 50,
    }, {
      width: 1024,
      //height: 379, //576,
      rename: {
        //prefix: new Date().toISOString().split('T')[0] + '-',
        suffix: '-large',
        extname: '.webp',
      },
      quality: 50,
    }],
    }, {
    // Global configuration for all images
    // Strip all metadata
    withMetadata: false,
    crop: true,
    background: "#000",
    // Do not emit the error when image is enlarged.
    errorOnEnlargement: false,
    errorOnUnusedConfig: false,
  }))
  .pipe(gulp.dest(img.cover.dst));
});

gulp.task('sass:dev', function() {
  return gulp.src(sass_src)
    .pipe(plugins.sass({
      includePaths: sassPaths,
      outputStyle: 'nested' // if css 'compressed' **file size**
    })
      .on('error', plugins.sass.logError))
    .pipe(plugins.autoprefixer({
      browsers: ['last 2 versions', 'ie >= 9']
    }))
    .pipe(gulp.dest(css_dest));
});

gulp.task('sass:prod', function() {
  return gulp.src(sass_src)
    .pipe(plugins.sass({
      includePaths: sassPaths,
      outputStyle: 'compressed' // if css 'compressed' **file size**
    })
      .on('error', plugins.sass.logError))
    .pipe(plugins.autoprefixer({
      browsers: ['last 2 versions', 'ie >= 9']
    }))
    .pipe(gulp.dest(css_dest));
});


gulp.task('css:prod', function() {
  return gulp.src('public/css/app.css')
  .pipe(plugins.purifycss(['./public/**/*.js', './public/**/*.html']))
  .pipe(plugins.shorthand())
  .pipe(plugins.csso())
  .pipe(gulp.dest('public/css/'))
});


gulp.task('lint:html', function () {
  return gulp.src(['content/**/*.html', 'layout/**/*.html', 'public/**/*.html', 'archetypes/**/*.html'])
  .pipe(plugins.htmlhint())
  .pipe(plugins.htmlhint.failReporter());
});

gulp.task('javascript:dev', function () {
  return gulp.src(javascript_src)
    .pipe(plugins.babel({ presets: ['es2015']}))
    .pipe(plugins.concat('app.js'))
    .pipe(gulp.dest(javascript_dest));
});

gulp.task('javascript:prod', function () {
  return gulp.src(javascript_src)
    .pipe(plugins.babel({ presets: ['es2015']}))
    .pipe(plugins.concat('app.js'))
    .pipe(plugins.uglify())
    .pipe(gulp.dest(javascript_dest));
});

gulp.task('html:prod', function() {
  return gulp.src('public/**/*.html')
    .pipe(plugins.htmlmin({collapseWhitespace: true}))
    .pipe(gulp.dest('public'));
});

gulp.task('clean', function() {
  return gulp.src('public')
  .pipe(plugins.clean());
});



gulp.task('hugo', function (cb) {
  exec('hugo', function (err, stdout, stderr) {
    console.log(stdout);
    console.log(stderr);
    cb(err);
  });
});

gulp.task('default', ['sass:dev', 'javascript:dev'], function() {
  gulp.watch(['src/scss/**/*.scss'], ['sass:dev']);
  gulp.watch(['src/js/enabled/**/*.js'], ['javascript:dev']);
});

gulp.task('build', function(callback) {
  runSequence('clean', ['sass:prod', 'javascript:prod'], 'hugo', 'css:prod', callback);
});
