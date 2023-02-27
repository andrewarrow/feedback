
function sendFormAsJson(e) {
  e.preventDefault();
  const form = document.getElementById('form1');
  const xhr = new XMLHttpRequest();
  xhr.open('POST', '/models');
  xhr.setRequestHeader('Content-Type', 'application/json');
  xhr.addEventListener('load', function(event) {
    console.log('Response received: ', event.target.response);
    document.getElementById('models').innerHTML = event.target.response;
  });

  const formData = new FormData(form);
  const formDataJson = {};
  formData.forEach(function(value, key) {
    formDataJson[key] = value;
  });

  const jsonData = JSON.stringify(formDataJson);
  xhr.send(jsonData);

}

