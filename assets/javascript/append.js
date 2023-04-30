
function sendFormAsJsonAndAppend(e) {
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
      const jsonResponse = xhr.response;

        var newTrElement = document.createElement('tr');
        var newTableElement = document.createElement('tbody');
        newTrElement.innerHTML = jsonResponse.table_large;
        newTableElement.innerHTML = jsonResponse.table_small;

        document.getElementById(jsonResponse.table_large_div).prepend(newTrElement);
        document.getElementById(jsonResponse.table_small_div).prepend(newTableElement);

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
