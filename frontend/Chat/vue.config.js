module.exports = {
    publicPath: './',
    devServer: {
        host: '0.0.0.0',
        disableHostCheck: true,
        proxy: {
            '/': {
                target: 'http://localhost:8081/',
                ws: true,
                changeOrigin: true
            }
        }
    },
};