function cacheSlider()
{
    var pop = document.getElementById("pop");
    pop.style.display = "none";//display slider
}


//init slider with image in arg
function slider(image)
{
    var index = list_files.indexOf(image);

    var pop = document.getElementById("pop");
    pop.style.display = "flex";//display slider

    pop.getElementsByClassName("after")[0].src = "/static/images/loader1.svg";
    pop.getElementsByClassName("after")[0].onclick = null;
    pop.getElementsByClassName("before")[0].src = "/static/images/loader1.svg";

    var images = pop.getElementsByClassName("images")[0];
    images.innerHTML = "";
    
    var img_dom = document.createElement("img");
    
    img_dom.className = "visible";
    images.appendChild(img_dom);
    img_dom.onload = function()
    {
        loadNext(index);
        loadPrevious(index,img_dom);
    }
    img_dom.src = image;

}

function loadNext(i)
{
    console.log("next !");
    if(i < list_files.length - 1)
    {
        var pop = document.getElementById("pop");
        var images = pop.getElementsByClassName("images")[0];
        var img_dom = document.createElement("img");
        img_dom.onload = function()
        {
            pop.getElementsByClassName("after")[0].src = "/static/images/arrow-right.svg";
            pop.getElementsByClassName("after")[0].onclick = function(){ next(i) };
            console.log("ok next!");
        }
        img_dom.src = list_files[i+1];
        images.appendChild(img_dom);
    }
}


function loadPrevious(i,image)
{
    if(i > 0)
    {
        var pop = document.getElementById("pop");
        var images = pop.getElementsByClassName("images")[0];
        var img_dom = document.createElement("img");
        img_dom.onload = function()
        {
            pop.getElementsByClassName("before")[0].src = "/static/images/arrow-left.svg";
            pop.getElementsByClassName("before")[0].onclick = function(){ previous(i) };
            console.log("ok before!");
        }
        img_dom.src = list_files[i-1];
        images.insertBefore(img_dom,image);
    }
}

function next(i)
{
    var pop = document.getElementById("pop");
    var enCours = pop.getElementsByClassName("visible")[0];
    enCours.classList.remove("visible");
    var suivant = enCours.nextElementSibling;
    if(suivant != null)
    {
        suivant.classList.add("visible");
        if(suivant.nextElementSibling == null)
        {
            pop.getElementsByClassName("after")[0].src = "/static/images/loader1.svg";
            pop.getElementsByClassName("after")[0].onclick = null;
            loadNext(i+1);
        }
        else
        {
            pop.getElementsByClassName("after")[0].onclick = function(){ next(i+1) };
        }
        
        pop.getElementsByClassName("before")[0].src = "/static/images/arrow-left.svg";
        pop.getElementsByClassName("before")[0].onclick = function(){ previous(i+1) };
        
    }
}

function previous(i)
{
    var pop = document.getElementById("pop");
    var enCours = pop.getElementsByClassName("visible")[0];
    enCours.classList.remove("visible");
    var precedent = enCours.previousElementSibling;
    if(precedent != null)
    {
        precedent.classList.add("visible");
        if(precedent.previousElementSibling == null)
        {
            pop.getElementsByClassName("before")[0].src = "/static/images/loader1.svg";
            pop.getElementsByClassName("before")[0].onclick = null;
            loadPrevious(i-1,precedent)
        }
        else
        {
            pop.getElementsByClassName("before")[0].onclick = function(){ previous(i-1) };
        }
        
        pop.getElementsByClassName("after")[0].src = "/static/images/arrow-right.svg";
        pop.getElementsByClassName("after")[0].onclick = function(){ next(i-1) };
        
    }
}