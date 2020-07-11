var stringToColour = function(str) {
    var hash = 0;
    for (var i = 0; i < str.length; i++) {
        hash = str.charCodeAt(i) + ((hash << 5) - hash);
    }
    var colour = '#';
    for (var i = 0; i < 3; i++) {
        var value = (hash >> (i * 8)) & 0xFF;
        colour += ('00' + value.toString(16)).substr(-2);
    }
    return colour;
}

function stringToInt(str)
{
    var total = 0;
    for(var lettre of str)
    {
        total = total + lettre.charCodeAt(0)/2;
    }
    return total;
}

function coloration()
{
    //color = ["#27bfe7","#5227e7","#c927e7","#f16bcd","#e71e50","#e78027","#e7bd27","#27e763","#4baf6b"];
    color = ["#fef9d1","#79a50a","#fd966d","#1a8ffb","#6b59cd","#38d9da","#ff3299","#e7bd27","fb02fe"];
    for(var ele of Array.from(document.getElementsByClassName("dossier")))
    {
        var str = ele.getElementsByTagName("h3");
        if (str.length > 0)
        {
          str = str[0].innerText;
          value = Math.round(stringToInt(str)) % 8;
          ele.style.borderColor = stringToColour(str);
          ele.style.backgroundColor = stringToColour(str)+'88';
        }
    }
}

coloration();
