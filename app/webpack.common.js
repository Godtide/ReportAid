const path = require('path')
const fs  = require('fs');
const htmlWebpackPlugin = require('html-webpack-plugin');
const cleanWebpackPlugin = require('clean-webpack-plugin');
const { CheckerPlugin } = require('awesome-typescript-loader')
const { TsConfigPathsPlugin } = require('awesome-typescript-loader');
const lessToJs = require('less-vars-to-js');
const themeVariables = lessToJs(fs.readFileSync(path.join(__dirname, './src/stylesheets/themeVariables.less'), 'utf8'));
themeVariables["@icon-url"] = 'https://fonts.googleapis.com/css?family=Source+Sans+Pro:regular,bold,italic&subset=latin,latin-ext';

var config = {
  node: {
    console: true,
    fs: 'empty',
    net: 'empty',
    tls: 'empty'
  },
  entry: {
    app: [
      'babel-polyfill',
      './src/index.tsx'
    ]
  },
  output: {
    path: path.resolve(__dirname, 'build'),
    filename: 'bundle.js'
  },
  resolve: {
    extensions: [".ts", ".tsx", ".js", ".json"]
  },
  plugins: [
    new cleanWebpackPlugin(['build']),
    new htmlWebpackPlugin({
      template: './src/index.html',
      inject: 'body',
    }),
    new CheckerPlugin(),
    new TsConfigPathsPlugin(/* { configFileName, compiler } */)
  ],
  externals: {
    "react": "React",
    "react-dom": "ReactDOM"
  },
  module: {
    rules: [
      {
        test: /\.less$/,
        use: [
          {
            loader: "style-loader" // creates style nodes from JS strings
          },
          {
            loader: "css-loader" // translates CSS into CommonJS
          },
          {
            loader: "less-loader", // compiles Less to CSS
            options: {
              javascriptEnabled: true,
              modifyVars: themeVariables
            }
          }
        ]
      },
      {
        test: /\.(jp(e*)g|png|gif|svg|pdf|ico)$/i,
        use: [
          {
            loader: 'url-loader',
            options: {
              limit: 8192,
              name: '[sha512:hash:base64:7].[ext]'
            }
          }
        ]
      },
      {
        test: /\.js$/,
        use: ["source-map-loader"],
        enforce: "pre"
      },
      {
        test: /\.jsx?/,
        include: path.resolve(__dirname, 'app'),
        exclude: /node_modules/,
        loader: 'babel-loader'
      },
      // All files with a '.ts' or '.tsx' extension will be handled by 'awesome-typescript-loader'.
      {
        test: /\.tsx?$/,
        loader: "awesome-typescript-loader"
      }
    ]
  }
}

module.exports = config
