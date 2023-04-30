function fillInFieldName() {
  const flavor = document.getElementById('flavor');
  const name = document.getElementById('name');
  name.value = flavor.value;
}

function doLogout(e) {
  e.preventDefault();
  const form = document.getElementById('logout');
  form.submit();
}

function doDeletePost(guid, e) {
  e.preventDefault();
  const form = document.getElementById('f'+guid);
        console.log(form)
  form.submit();
}

function sendFormAsJson(e) {
  e.preventDefault();
  const route = e.target.action;
  const form = e.target;
  const xhr = new XMLHttpRequest();
  xhr.open('POST', route);
  xhr.setRequestHeader('Content-Type', 'application/json');
  xhr.addEventListener('load', function(event) {
     if (xhr.status != 200) {
        document.getElementById('flash').innerHTML = event.target.response;
     } else {
        document.getElementById('flash').innerHTML = event.target.response;
     }
  });

  const formData = new FormData(form);
  const formDataJson = {};
  formData.forEach(function(value, key) {
    formDataJson[key] = value;
  });

  const jsonData = JSON.stringify(formDataJson);
  xhr.send(jsonData);

}

function sendFormAsJsonAndReplace(e) {
  e.preventDefault();
  const route = e.target.action;
  const form = e.target;
  const formData = new FormData(form);

  var queryString = "?";
  formData.forEach(function(value, key) {
    queryString += key + "=" + value + "&";
  });

  queryString = queryString.slice(0, -1);

  const xhr = new XMLHttpRequest();
  xhr.open('POST', route);
  xhr.responseType = 'json';
  xhr.setRequestHeader('Content-Type', 'application/json');
  xhr.setRequestHeader('Feedback-Ajax', 'true');
  xhr.addEventListener('load', function(event) {
     if (xhr.status != 200) {
        document.getElementById('flash').innerHTML = event.target.response;
     } else {
        history.pushState(null, null, route+queryString);
      const jsonResponse = xhr.response;
      //console.log(jsonResponse);
      const div = jsonResponse.div;

        document.getElementById(div).innerHTML = jsonResponse.html;
        document.getElementById('flash').innerHTML = '';
        setAllLinks();
     }
  });

  const formDataJson = {};
  formData.forEach(function(value, key) {
    formDataJson[key] = value;
  });

  const jsonData = JSON.stringify(formDataJson);
  xhr.send(jsonData);

}

function getViaAjax(e) {
  e.preventDefault();
  document.getElementById("feedback-ajax").innerHTML = "<div class='p-10 font-bold text-6xl'>Please wait...</div>";
  document.getElementById('flash').innerHTML = '';

  const route = e.target.href;
  history.pushState(null, null, route);
  const xhr = new XMLHttpRequest();
  xhr.responseType = 'json';
  xhr.open('GET', route);
  xhr.setRequestHeader('Feedback-Ajax', 'true');
  xhr.addEventListener('load', function() {
    handleLoadEvent(xhr, e.target.href);
  });
  xhr.send();
}

function handleLoadEvent(xhr, href) {

   if (xhr.status != 200) {
      document.getElementById('flash').innerHTML = xhr.response;
      document.getElementById('feedback-ajax').innerHTML = '';
   } else {
      const jsonResponse = xhr.response;
      //console.log(jsonResponse);
      const div = jsonResponse.div;
      const next = jsonResponse.next;

      document.getElementById(div).innerHTML = jsonResponse.html;
      document.getElementById('flash').innerHTML = '';
      setAllLinks();

      if (next != null) {
        const route = href + "?offset="+next;
        xhr.open('GET', route);
        xhr.responseType = 'json';
        xhr.setRequestHeader('Feedback-Ajax', 'true');
        xhr.send();
      } 
   }
}

function setAllLinks() {
  const links = document.getElementsByTagName("a"); 
  for (let i = 0; i < links.length; i++) {
    var classes = links[i].className;
    if (classes.indexOf("no-ajax") == -1) {
      links[i].onclick = getViaAjax; 
    }
  }
}

document.addEventListener("DOMContentLoaded", function() {
  setAllLinks();
  window.onpopstate = function(event) {
    location.reload();
  };
});

