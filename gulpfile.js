var gulp = require('gulp');
var plugins = require('gulp-load-plugins')();
var dest_css = "static/css";
var sass_src = "src/scss/app.scss";
var runSequence = require('run-sequence');
var javascript_src = ['src/js/enabled/foundation.core.js', 'src/js/enabled/foundation.util.*.js', 'src/js/enabled/*.js'];
var dest_javascript = "static/js";
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
    .pipe(gulp.dest(dest_css));
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
    .pipe(plugins.shorthand())
    .pipe(plugins.csso())
    .pipe(gulp.dest(dest_css));
});


gulp.task('uncss', function() {
  return gulp.src('public/css/app.css')
  .pipe(plugins.uncss({
          html: ['public/**/*.html']
      }))
  .pipe(gulp.dest('public/css/app.css'))
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
    .pipe(gulp.dest(dest_javascript));
});

gulp.task('javascript:prod', function () {
  return gulp.src(javascript_src)
    .pipe(plugins.babel({ presets: ['es2015']}))
    .pipe(plugins.concat('app.js'))
    .pipe(plugins.uglify())
    .pipe(gulp.dest(dest_javascript));
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
  runSequence('clean', ['sass:prod', 'javascript:prod'], 'hugo', 'html:prod', callback);
});
