
window.onload = function() {
    var pop = document.getElementById("pop");
    pop.addEventListener("click", (e) => {
        console.log("not implemented")
    });
}



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

    img_dom.onclick = (e) => { window.open(list_original[index]) }
}

// enable click for next img
function enableClickNext(index_image)
{
    var pop = document.getElementById("pop");
    pop.getElementsByClassName("next-img")[0].classList.remove("progress");
    pop.getElementsByClassName("next-img")[0].onclick = function(){ next(index_image) };
    document.getElementById("pop--next-button").classList.remove("progress");
    document.getElementById("pop--next-button").onclick = function(){ next(index_image) };

}

// disable click for next img and set cursor as loading
function disableClickNext()
{
    var pop = document.getElementById("pop");
    pop.getElementsByClassName("next-img")[0].classList.add("progress");
    pop.getElementsByClassName("next-img")[0].onclick = null
    document.getElementById("pop--next-button").classList.add("progress");
    document.getElementById("pop--next-button").onclick = null
}

//display next miniature
function loadNext(index_image)
{
    if(index_image < list_files.length - 1)
    {
        var pop = document.getElementById("pop");
        var images = pop.getElementsByClassName("images")[0];
        var img_dom = document.createElement("img");
        img_dom.classList.add("next-img");
        img_dom.classList.add("progress");
        img_dom.onload = () => enableClickNext(index_image)
        img_dom.src = list_files[index_image+1];
        images.appendChild(img_dom);
    }
}

// enable click for previous img
function enableClickPrevious(index_image)
{
    var pop = document.getElementById("pop");
    pop.getElementsByClassName("previous-img")[0].classList.remove("progress");
    pop.getElementsByClassName("previous-img")[0].onclick = function(){ previous(index_image) };
    document.getElementById("pop--previous-button").onclick = function(){ previous(index_image) };
}

// disable click for previous img and set cursor as loading
function disableClickPrevious()
{
    var pop = document.getElementById("pop");
    pop.getElementsByClassName("previous-img")[0].classList.add("progress");
    pop.getElementsByClassName("previous-img")[0].onclick = null
    document.getElementById("pop--previous-button").classList.add("progress");
    document.getElementById("pop--previous-button").onclick = null
}

function loadPrevious(index_image,image)
{
    if(index_image > 0)
    {
        var pop = document.getElementById("pop");
        var images = pop.getElementsByClassName("images")[0];
        var img_dom = document.createElement("img");
        img_dom.classList.add("previous-img");
        img_dom.classList.add("progress");
        img_dom.onload = () => enableClickPrevious(index_image)
        img_dom.src = list_files[index_image-1];
        images.insertBefore(img_dom,image);
    }
}

function resetState(previousImg,currentImg,nextImg,index_image)
{
    pop.querySelectorAll(".previous-img").forEach(e => e.classList.remove('previous-img'));
    pop.querySelectorAll(".next-img").forEach(e => e.classList.remove('next-img'));
    pop.querySelectorAll(".visible").forEach(e => e.classList.remove('visible'));

    if(previousImg != null)
    {
        previousImg.classList.add("previous-img")
        enableClickPrevious(index_image)
    }
    else
    {
        loadPrevious(index_image,currentImg)
    }

    if(nextImg != null)
    {
        nextImg.classList.add("next-img")
        enableClickNext(index_image)
    }
    else
    {
        loadNext(index_image)
    }

    currentImg.classList.add('visible')
    console.log(currentImg)
    currentImg.onclick = (e) => { window.open(list_original[index_image]) }
}


function next(index_image)
{
    var pop = document.getElementById("pop");
    var enCours = pop.getElementsByClassName("visible")[0];
    var suivant = enCours.nextElementSibling;
    var precedent = enCours.previousElementSibling;
    resetState(enCours,suivant,suivant.nextElementSibling,index_image+1)
}

function previous(index_image)
{
    var pop = document.getElementById("pop");
    var enCours = pop.getElementsByClassName("visible")[0];
    var precedent = enCours.previousElementSibling;
    var suivant = enCours.nextElementSibling;
    resetState(precedent.previousElementSibling,precedent,enCours,index_image-1)
}
