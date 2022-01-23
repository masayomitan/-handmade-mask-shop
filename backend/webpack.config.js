var debug   = process.env.NODE_ENV !== "production";
var webpack = require('webpack');
var path    = require('path');
const tailwindCss = require('tailwindcss')
const autoprefixer = require('autoprefixer')

module.exports = {
    context: path.join(__dirname, "src"),
    entry: { 
      js: "./js/app.js",

    },
    stats: 'errors-only',
    output: {
      path: __dirname + "/src/",
      filename: "[name]_bundle.js",
    },
    module: {
        rules: [{
          test: /\.js$|jsx/,
              exclude: /(node_modules|bower_components)/,
              use: [{
                  loader: 'babel-loader',
                  options: {
                      presets: ['@babel/preset-react', '@babel/preset-env'],
                      plugins: ['@babel/plugin-transform-runtime'],
                  }
              }]
          },
          {
          test: /\.css/,
            use: [{
              loader: "css-loader", 
            }]
          },
          {
            test: /\.css/,
            use: [{
              loader: 'postcss-loader',
              options: {
                postcssOptions: {
                  ident: 'postcss',
                  plugins: [
                    tailwindCss,
                    autoprefixer
                  ],
                },
              }
            }]
          },
        ],
    },
    
    plugins: debug ? [] : [
      new webpack.optimize.OccurrenceOrderPlugin(),
      new webpack.optimize.UglifyJsPlugin({ mangle: false, sourcemap: false }),
    ]
  
};