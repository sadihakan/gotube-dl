var search = require('youtube-search');
const youtubedl = require('youtube-dl')

const argm = process.argv.slice(2)

var opts = {
    maxResults: parseInt(argm[0]),
    key: argm[1]
};

search(argm[2], opts, function (err, results) {
    if (err) return console.log(err);

    youtubedl.getInfo([results[0].link], function (err, info) {
        if (err) return console.log(err);
        console.log(info.url)
    })
});


// var opts = {
//     maxResults: 1,
//     key: 'AIzaSyC0-9QZFnXu6SpFmpkaWkktWvyIDhqXR1M'
// };