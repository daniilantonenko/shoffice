//Пример начального JavaScript для отключения отправки формы, если есть недопустимые поля
(() => {
    'use strict'

    // Получить все формы, к которым мы хотим применить пользовательские стили проверки Bootstrap.
    const forms = document.querySelectorAll('.needs-validation')

    // Перебрать их и предотвратить отправку
    Array.from(forms).forEach(form => {
        form.addEventListener('submit', event => {
            if (!form.checkValidity()) {
                event.preventDefault()
                event.stopPropagation()
            }

            form.classList.add('was-validated')
        }, false)
    })
})()


$('#152-fz').click(function () {
    if ($('#submit').is(':disabled')) {
        $('#submit').removeAttr('disabled');
    } else {
        $('#submit').attr('disabled', 'disabled');
    }
});


$("#sendForm").submit(function (event) {
    event.preventDefault();

    var fileName = $('#inputFile').val(); // имя файла

    if (fileName != "") {
        // событие при отправке формы
        $(this).hide();

        // проверка согласия на обработку персональных данных
        if ($('input.form-check-input').is(':checked')) {
            //submitForm();
            $.ajax({
                //url: $(this).attr("action"),
                //
                type: "POST",
                url: "/action/index.php",

                //data: form_data,
                data: new FormData(this),
                processData: false,
                contentType: false,
                success: function (output) {
                    if (output == "Done") {
                        console.log("Файл успешно загружен")
                    } else {
                        alert("Непредвиденная ошибка, необходимо обратиться к системному администратору.");
                    }

                },
                error: function (output) {
                    alert(output);
                }
            });
        } else {
            console.log("Необходимо подтвердить согласание на обработку данных.")
        }
        $("#endSendForm").show();

        console.log("submit")
        return false;
    }
});

function submitForm() {
    console.log("submitForm()");

}