var url = 'folder/index.html?param=#23dd&noob=yes'; //or specify one

var encodedUrl = encodeURIComponent(url);
console.log(encodedUrl);


var decodedUrl = decodeURIComponent(url);
console.log(decodedUrl);



//Ref: https://www.sitepoint.com/jquery-decode-url-string/
