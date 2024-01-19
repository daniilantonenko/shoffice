function ValidateIPaddress(ipaddress) {  
    if (/^(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$/.test(ipaddress)) {  
      return (true)  
    }  
    return (false)  
    } 

    var selectedOption = document.getElementById("selectedOption");
    var diffrentInput = document.getElementById("inputDiffrent");

    if (selectedOption != null){
      selectedOption.addEventListener('change',function(){
        diffrentInput.classList.remove("d-none")
      })
    }

    function sendForm() {
      // selectedOption == null || selectedOption.value == "diffrent"

      selectedOption = document.getElementById("selectedOption");
      diffrentInput = document.getElementById("inputDiffrent");

      if (selectedOption != null){
        if (selectedOption.value != "diffrent"){
          // Get value from options
          selectedOption = document.getElementById("selectedOption").value;
        }else{
          // Get value from diffrent input
          selectedOption = document.getElementById("inputDiffrent").value;
        }
      }else{
        // Get value from diffrent input
        selectedOption = document.getElementById("inputDiffrent").value;
      }

      if (ValidateIPaddress(selectedOption)){
        // IP is valid
        var xhr = new XMLHttpRequest();
        xhr.open("POST", "/qrcode", true);
        xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
        xhr.onreadystatechange = function() {
          if (xhr.readyState === 4 && xhr.status === 200) {
            document.getElementById("result").src = xhr.responseText;
          }
        };
        xhr.send(selectedOption);
        return false;
      }else{
        // IP is invalid
        // TODO: add message "You have entered an invalid IP address!")
        diffrentInput.classList.add("is-invalid")
        
        return false;
      }
    }