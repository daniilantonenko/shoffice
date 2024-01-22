function ValidateIPaddress(ipaddress) {
  if (/^(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$/.test(ipaddress)) {
    return (true)
  }
  return (false)
}

function SendPostAndReturnImage(ipaddress){
  var xhr = new XMLHttpRequest();
    xhr.open("POST", "/qrcode", true);
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    xhr.onreadystatechange = function () {
      if (xhr.readyState === 4 && xhr.status === 200) {
        document.getElementById("result").src = xhr.responseText;
      }
    };
    xhr.send(ipaddress);
}

var selectedOption = document.getElementById("selectedOption");
var diffrentInput = document.getElementById("inputDiffrent");
var diffrentGroup = document.getElementById("groupDiffrent");
var valueAddress = "";

// Show or hide "user address" input field
if (selectedOption != null) {
  selectedOption.addEventListener('change', function () {
    if(selectedOption.value === "diffrent")
    {
      diffrentGroup.classList.remove("hide");
    }else{
      diffrentGroup.classList.add("hide");
    }
  })
}

function sendForm() {

  // Selecting a path depending on the filling of the "user address" input field
  if (selectedOption != null) {
    if (selectedOption.value != "diffrent") {
      // Get value from options
      valueAddress = document.getElementById("selectedOption").value;
    } else {
      // Get value from diffrent input
      valueAddress = document.getElementById("inputDiffrent").value;
    }
  } else {
    // Get value from diffrent input
    valueAddress = document.getElementById("inputDiffrent").value;
  }

  if (ValidateIPaddress(valueAddress)) {
    // IP is valid
    if (selectedOption != null) {
      diffrentInput.classList.remove("is-invalid")
    }

    SendPostAndReturnImage(valueAddress);

    return false;
  } else {
    // IP is invalid
    diffrentInput.classList.add("is-invalid")

    return false;
  }
}