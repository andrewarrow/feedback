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

function vote(guid, e) {
  e.preventDefault();
  const xhr = new XMLHttpRequest();
  xhr.open('POST', '/vote/'+guid+'/');
  xhr.setRequestHeader('Content-Type', 'application/json');
  xhr.addEventListener('load', function(event) {
     if (xhr.status != 200) {
        document.getElementById('flash').innerHTML = event.target.response;
     } else {
        document.getElementById('v'+guid).innerHTML = '&nbsp;&nbsp;&nbsp;';
        document.getElementById('flash').innerHTML = '';
     }
  });
  xhr.send();
}

function sendFormAsJson(name, e) {
  e.preventDefault();
  const form = document.getElementById('form1');
  const xhr = new XMLHttpRequest();
  xhr.open('POST', '/'+name+'/');
  xhr.setRequestHeader('Content-Type', 'application/json');
  xhr.addEventListener('load', function(event) {
     if (xhr.status != 200) {
        document.getElementById('flash').innerHTML = event.target.response;
     } else {
        document.getElementById(name).innerHTML = event.target.response;
        document.getElementById('name').value = '';
        document.getElementById('flash').innerHTML = '';
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

