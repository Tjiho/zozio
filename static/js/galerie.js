function stringToInt(str)
{
    var total = 0;
    for(var lettre of str)
    {
        total = total + lettre.charCodeAt(0);
    }    
    return total;
}

function coloration()
{   
    color = ["#27bfe7","#5227e7","#c927e7","#f16bcd","#e71e50","#e78027","#e7bd27","#27e763","#4baf6b"];
    for(var ele of document.getElementsByClassName("dossier"))
    {
        var str = ele.getElementsByTagName("h3")[0].innerText;
        
        value = stringToInt(str) % 9;
        
        ele.style.backgroundColor = color[value];
    }
}

coloration();