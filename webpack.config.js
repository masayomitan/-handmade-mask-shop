var debug   = process.env.NODE_ENV !== "production";
var webpack = require('webpack');
var path    = require('path');

module.exports = {
    context: path.join(__dirname, "src"),
    entry: { 
      js: "./js/app.js",
      css: "./css/style.css"
    },
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
                      presets: ['@babel/preset-react', '@babel/preset-env']
                  }
              }]
          },
          {
              // 対象となるファイルの拡張子
              test: /\.css/,
              // ローダー名
              use: [
                  // linkタグに出力する機能
                  "style-loader",
                  // CSSをバンドルするための機能
                  {
                      loader: "css-loader",
                          options: {
                            // オプションでCSS内のurl()メソッドの取り込みを禁止する
                            url: false,
                          }
                  }
                ]
          }
          ]
    },
    
    plugins: debug ? [] : [
      new webpack.optimize.OccurrenceOrderPlugin(),
      new webpack.optimize.UglifyJsPlugin({ mangle: false, sourcemap: false }),
    ]
  
};