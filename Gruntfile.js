module.exports = function(grunt) {

  grunt.initConfig({

    // COPY VENDOR =============================================================

    bowercopy: {
      options: {
        srcPrefix: "bower_components"
      },
      scripts: {
        options: {
          destPrefix: "dist/vendor"
        },
        files: {
          'moment/moment.min.js': 'moment/min/moment.min.js',
          'jquery/jquery.min.js': 'jquery/dist/jquery.min.js',
          'jquery/jquery.min.map': 'jquery/dist/jquery.min.map',
          'angular/angular.min.js': 'angular/angular.min.js',
          'angular/angular.min.js.map': 'angular/angular.min.js.map',
          'angular-ui-router/angular-ui-router.js': 'angular-ui-router/release/angular-ui-router.js',
          'bootstrap/bootstrap-theme.min.css': 'bootstrap/dist/css/bootstrap-theme.min.css',
          'bootstrap/bootstrap-theme.min.css.map': 'bootstrap/dist/css/bootstrap-theme.min.css.map',
          'bootstrap/bootstrap.min.js': 'bootstrap/dist/js/bootstrap.min.js',
        }
      }
    },

    // IMAGES =================================================================

    imagemin: {
      dynamic: {
        files: [{
          expand: true,
          cwd: 'public/images/',
          src: ['**/*.{png,jpg,gif}'],
          dest: 'dist/images/'
        }]
      }
    },

    // ANGULAR TEMPLATES =======================================================

    ngtemplates: {
      myapp: {
        src:      'public/src/**/*.html',
        dest:     'dist/template.js'
      }
    },

    // JS TASKS ================================================================
    jshint: {
      all: ['public/src/**/*.js'] 
    },

    uglify: {
      build: {
        options: {
          beautify : true,
          mangle   : true
        },
        files: {
          'dist/app.min.js': ['public/src/**/*.js', 'public/src/*.js', '!app/app.js', 'app/app.js']
        }
      }
    },

    // CSS TASKS ===============================================================
    less: {
      build: {
        files: {
          'dist/style.css': 'public/src/style/style.less'
        }
      }
    },

    cssmin: {
      build: {
        files: {
          'dist/style.min.css': 'dist/style.css'
        }
      }
    },

    // COOL TASKS ==============================================================
    watch: {
      css: {
        files: ['public/src/style/**/*.less'],
        tasks: ['less', 'cssmin']
      },
      js: {
        files: ['public/src/**/*.js'],
        tasks: ['jshint', 'uglify']
      },
      html: {
        files: ['public/src/**/*.html'],
        tasks: ['ngtemplates']
      }
    },

    concurrent: {
      options: {
        logConcurrentOutput: true
      },
      tasks: ['watch']
    }   

  });

  grunt.loadNpmTasks('grunt-contrib-imagemin');
  grunt.loadNpmTasks('grunt-contrib-jshint');
  grunt.loadNpmTasks('grunt-contrib-uglify');
  grunt.loadNpmTasks('grunt-contrib-less');
  grunt.loadNpmTasks('grunt-contrib-cssmin');
  grunt.loadNpmTasks('grunt-contrib-watch');
  grunt.loadNpmTasks('grunt-nodemon');
  grunt.loadNpmTasks('grunt-concurrent');
  grunt.loadNpmTasks('grunt-bowercopy');
  grunt.loadNpmTasks('grunt-angular-templates');

  grunt.registerTask('default', ['bowercopy', 'imagemin', 'ngtemplates', 'less', 'cssmin', 'jshint', 'uglify']);

};