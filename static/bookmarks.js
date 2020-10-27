// ON START
const IP = 'http://127.0.0.1:8000';
let bookmarks = document.getElementById('bookmarks');


///// CREATE NEW ELEMENT
function createElement(element) {
  let e = document.createElement(element);
  return e
}


///// CLOSURE FUNCTION /////
function closure(id, route) {
  return function () {
    delRequest(id, route);
  }
}


///// CREATE NEW BOOKMARK /////
function newBookmark(id, url, description) {
  let link = createElement('a');
  link.href = url;
  link.innerHTML = description;
  link.target = '_blank';
  bookmarks.appendChild(link);

  let del = createElement('button');
  del.id = 'del';
  del.onclick = closure(id, '/delBookmark');
  del.innerHTML = 'delete';
  bookmarks.appendChild(del);

  bookmarks.appendChild(createElement('br'));
}


///// 'ADD NEW' BUTTON /////
function addNewButton() {
  let addNew = createElement('button');
  addNew.id = 'addNew';
  addNew.innerHTML = 'add new';
  addNew.onclick = addBookmark;
  bookmarks.appendChild(addNew);
}


///// LOAD BOOKMARKS /////
async function loadBookmarks() {
  bookmarks.innerHTML = '';

  // menu asterix animation
  document.getElementById('menu_bm').style = 'font-size:45px';
  document.getElementById('mark_bm').innerHTML = '*';
  document.getElementById('menu_nt').style = 'font-size:40px';
  document.getElementById('mark_nt').innerHTML = '';

  // load array of response object containing bookmarks
  let response = await getRequest('/loadBookmarks');
  if (response) {
    for (bm of response) {
      newBookmark(bm.Id, bm.Url, bm.Description);
    }
  }

  bookmarks.appendChild(createElement('br'));
  addNewButton();
}


/// ADD BOOKMARK /////
async function addBookmark() {
  // delete 'add new' button
  document.getElementById('addNew').remove();

  let br1 = createElement('br');
  br1.id = 'br1';
  bookmarks.appendChild(br1);

  // input url
  let inputUrl = createElement('input');
  inputUrl.id = 'inputUrl';
  inputUrl.placeholder = 'enter url';
  inputUrl.type = 'text';
  bookmarks.appendChild(inputUrl);
  inputUrl.focus();

  let br2 = createElement('br');
  br2.id = 'br2';
  bookmarks.appendChild(br2);

  // input description
  let inputDesc = createElement('input');
  inputDesc.id = 'inputDesc';
  inputDesc.placeholder = 'enter description';
  inputDesc.type = 'text';
  bookmarks.appendChild(inputDesc);

  let br3 = createElement('br');
  br3.id = 'br3';
  bookmarks.appendChild(br3);

  // save button
  let save = createElement('button');
  save.id = 'save';
  save.innerHTML = 'save';
  save.onclick = saveURL;
  bookmarks.appendChild(save);

  // cancel button
  let cancel = createElement('button');
  cancel.id = 'cancel';
  cancel.innerHTML = 'cancel';
  cancel.onclick = cancelButton;
  bookmarks.appendChild(cancel);
}


///// SAVE URL /////
function saveURL() {
  let inputUrl = document.getElementById('inputUrl').value;
  let inputDesc = document.getElementById('inputDesc').value;

  if (inputUrl && inputDesc) {
    let data = {
      id: 0,
      url: inputUrl,
      description: inputDesc
    }
    console.log(data);
    postRequest(data, "/addBookmark");
  }
}


///// CANCEL BUTTON /////
function cancelButton() {
  document.getElementById('br1').remove();
  document.getElementById('inputUrl').remove();
  document.getElementById('br2').remove();
  document.getElementById('inputDesc').remove();
  document.getElementById('br3').remove();
  document.getElementById('save').remove();
  document.getElementById('cancel').remove();
  addNewButton();
}


///// GET REQUEST /////
async function getRequest(route) {
  let url = `${IP}${route}`;

  let response = await fetch(url);
  if (response.ok) {
    let data = await response.json()
    return data
  }
}


///// POST REQUEST /////
async function postRequest(data, route) {
  let url = `${IP}${route}`;
  let package = {
    method: 'POST',
    header: {'content-type': 'application/json'},
    body: JSON.stringify(data)
  }

  let response = await fetch(url, package);
  if (response.ok) {
    loadBookmarks();
  }
}


///// DELETE REQUEST /////
async function delRequest(data, route) {
  let url = `${IP}${route}`;
  let package = {
    method: 'DELETE',
    header: {'content-type': 'application/json'},
    body: JSON.stringify(data)
  }

  let response = await fetch(url, package);
  if (response.ok) {
    loadBookmarks();
  }
}

/////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////

///// LOAD NOTES /////
function loadNotes() {
  bookmarks.innerHTML = '';
  document.getElementById('menu_bm').style = 'font-size:40px';
  document.getElementById('mark_bm').innerHTML = '';
  document.getElementById('menu_nt').style = 'font-size:45px';
  document.getElementById('mark_nt').innerHTML = '*';
}
