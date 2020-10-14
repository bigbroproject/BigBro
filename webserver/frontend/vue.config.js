let cssConfig = {};

if (process.env.NODE_ENV == "production") {
    cssConfig = {
        extract: {
            filename: "[name].css",
            chunkFilename: "[name].css"
        }
    };
}



module.exports = {
    lintOnSave: false,
    devServer: {
        proxy: 'http://localhost:8181',
    },
    outputDir:"../www",
    publicPath:"/dashboard/",
    assetsDir:"assets",
    // chainWebpack: config => {
    //     let limit = 9999999999999999;
    //
    //     config.module
    //         .rule("images")
    //         .test(/\.(png|gif|jpg)(\?.*)?$/i)
    //         .use("url-loader")
    //         .loader("url-loader")
    //         .tap(options => Object.assign(options, { limit: limit }));
    //     config.module
    //         .rule("fonts")
    //         .test(/\.(woff2?|eot|ttf|otf)(\?.*)?$/i)
    //         .use("url-loader")
    //         .loader("url-loader")
    //         .options({
    //             limit: limit
    //         });
    //
    // }
}