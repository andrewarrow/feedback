
function doLogout(e) {
  e.preventDefault();
  const form = document.getElementById('logout');
  form.submit();
}

function doDeletePost(guid, e) {
  e.preventDefault();
  const form = document.getElementById('f'+guid);
  form.submit();
}

function sendFormAsJson(e) {
  e.preventDefault();
  const form = document.getElementById('form1');
  const xhr = new XMLHttpRequest();
  xhr.open('POST', '/models/');
  xhr.setRequestHeader('Content-Type', 'application/json');
  xhr.addEventListener('load', function(event) {
     if (xhr.status != 200) {
        document.getElementById('flash').innerHTML = event.target.response;
     } else {
        document.getElementById('models').innerHTML = event.target.response;
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

