function b64EncodeUnicode(str) {
    // first we use encodeURIComponent to get percent-encoded Unicode,
    // then we convert the percent encodings into raw bytes which
    // can be fed into btoa.
    return btoa(encodeURIComponent(str).replace(/%([0-9A-F]{2})/g,
        function toSolidBytes(match, p1) {
            return String.fromCharCode('0x' + p1);
    }));
}

function rus_to_latin ( str ) {
    
    var ru = {
        'а': 'a', 'б': 'b', 'в': 'v', 'г': 'g', 'д': 'd', 
        'е': 'e', 'ё': 'e', 'ж': 'j', 'з': 'z', 'и': 'i', 
        'к': 'k', 'л': 'l', 'м': 'm', 'н': 'n', 'о': 'o', 
        'п': 'p', 'р': 'r', 'с': 's', 'т': 't', 'у': 'u', 
        'ф': 'f', 'х': 'h', 'ц': 'c', 'ч': 'ch', 'ш': 'sh', 
        'щ': 'shch', 'ы': 'y', 'э': 'e', 'ю': 'u', 'я': 'ya'
    }, n_str = [];
    
    str = str.replace(/[ъь]+/g, '').replace(/й/g, 'i');
    
    for ( var i = 0; i < str.length; ++i ) {
       n_str.push(
              ru[ str[i] ]
           || ru[ str[i].toLowerCase() ] == undefined && str[i]
           || ru[ str[i].toLowerCase() ].toUpperCase()
       );
    }
    
    return n_str.join('');
}

function setCorrectAnswer(a) {
    var changeLin = document.querySelector('.grade');
    changeLin.innerText = a;
}

var changeLink = document.querySelector('.breadcrumb').querySelector(".breadcrumb-item").querySelector("a");

changeLink.setAttribute("href", "#");

var subject = changeLink.textContent
var answer = document.querySelector('.qtext').textContent;

changeLink.onclick = function() {
    var r = fetch(`http://localhost:8080/${rus_to_latin(subject)}?answer=${b64EncodeUnicode(rus_to_latin(answer))}`, {
   headers: {
      'Accept': 'application/json'
   }
}).then(response => response.text())
   .then(text => setCorrectAnswer(text))
}