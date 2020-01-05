
function nightMode()
{
    document.querySelector("body").classList.toggle('night')
    if (document.querySelector("body").classList.contains('night'))
        sendNightRequest(1)
    else
        sendNightRequest(0)
}

function sendNightRequest(status)
{
  var xhttp = new XMLHttpRequest();
  xhttp.onreadystatechange = function() {
    if (this.readyState == 4 && this.status == 200) {
        console.log('nightMode done !')
    }
  };
  xhttp.open("POST", "/nightMode", true);
  xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
  xhttp.send("nightMode="+status);

}
