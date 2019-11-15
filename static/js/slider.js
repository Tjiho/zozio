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

    //pop.getElementsByClassName("after")[0].src = "/static/images/loader1.svg";
    pop.getElementsByClassName("after")[0].onclick = null;
    //pop.getElementsByClassName("before")[0].src = "/static/images/loader1.svg";

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
            // pop.getElementsByClassName("after")[0].src = "/static/images/arrow-right.svg";
        	pop.getElementsByClassName("after")[0].classList.remove("progress");
            pop.getElementsByClassName("after")[0].onclick = function(){ next(i) };
    		img_dom.classList.add("next-img");
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
            // pop.getElementsByClassName("before")[0].src = "/static/images/arrow-left.svg";
        	pop.getElementsByClassName("before")[0].classList.remove("progress");
            pop.getElementsByClassName("before")[0].onclick = function(){ previous(i) };
    		img_dom.classList.add("previous-img");
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
    enCours.classList.add("previous-img");
    var suivant = enCours.nextElementSibling;
    var precedent = enCours.previousElementSibling;
    if(suivant != null)
    {
        suivant.classList.add("visible");
    	suivant.classList.remove("next-img");
    	if (precedent != null)
			precedent.classList.remove("previous-img");
        if(suivant.nextElementSibling == null)
        {
            pop.getElementsByClassName("after")[0].classList.add("progress")
            pop.getElementsByClassName("after")[0].onclick = null;
            loadNext(i+1);
        }
        else
        {
            pop.getElementsByClassName("after")[0].onclick = function(){ next(i+1) };
        	suivant.nextElementSibling.classList.add("next-img")
		}
        
        pop.getElementsByClassName("before")[0].classList.remove("progress");
        pop.getElementsByClassName("before")[0].onclick = function(){ previous(i+1) };
        
    }
}

function previous(i)
{
    var pop = document.getElementById("pop");
    var enCours = pop.getElementsByClassName("visible")[0];
    enCours.classList.remove("visible");
    enCours.classList.add("next-img");
    
    var precedent = enCours.previousElementSibling;
    var suivant = enCours.nextElementSibling;
    if(precedent != null)
    {

    	precedent.classList.remove("previous-img");
        precedent.classList.add("visible");
    	if (suivant != null)
			suivant.classList.remove("next-img");
        if(precedent.previousElementSibling == null)
        {
            // pop.getElementsByClassName("before")[0].src = "/static/images/loader1.svg";
            pop.getElementsByClassName("before")[0].onclick = null;
            loadPrevious(i-1,precedent)
        }
        else
        {
            pop.getElementsByClassName("before")[0].onclick = function(){ previous(i-1) };
        	precedent.previousElementSibling.classList.add("previous-img");

		}
        
        // pop.getElementsByClassName("after")[0].src = "/static/images/arrow-right.svg";
        pop.getElementsByClassName("after")[0].onclick = function(){ next(i-1) };
        
    }
}
