<?php

require_once('../config/config.php');

//Get uploaded file data using $_FILES array
$tmp_name = $_FILES['file']['tmp_name']; // get the temporary file name of the file on the server
$filename = $_FILES['file']['name']; // get the name of the file
$filesize = $_FILES['file']['size']; // get size of the file for size validation
$filetype = $_FILES['file']['type']; // get type of the file
$fileerror = $_FILES['file']['error']; // get the error (if any)

//validate form field for attaching the file
if ($fileerror > 0) {
    echo ("Ошибка загрузки или файлы не загружены.");
    exit;
}

// Upload attachment file
if (!empty($_FILES["file"]["name"])) {

    // File path config
    $targetDir = UPLOAD_DIR;
    $fileName = basename($filename);
    $targetFilePath = $targetDir . $fileName;
    $fileType = pathinfo($targetFilePath, PATHINFO_EXTENSION);

    // Allow certain file formats
    $allowTypes = array('pdf', 'doc', 'docx', 'jpg', 'png', 'jpeg');
    if (in_array($fileType, $allowTypes)) {
        // Check file size
        if ($filesize > 2000000) {
            error_log ("ERROR 5");
            echo ("Размер файла должен быть меньше 2 МБ.");
            exit;
        } else {
            // Upload file to the server
            if (move_uploaded_file($tmp_name, $targetFilePath)) {
                $uploadedFile = $targetFilePath;
                echo "Done";
            } else {
                echo ("При загрузке файла произошла серверная ошибка. Необходимо связаться с технчиеской поддердкой.");
                echo ("\n");
                echo ($filesize);
                echo ("\n");
                echo ($targetFilePath);
                echo ("\n");
                echo ($tmp_name);
                echo ("\n");
                echo ($errorMSG);
                echo ("\n");
                echo ($fileType);
                exit;
            }
        }
    } else {
        echo ("К сожалению, разрешено загружать только файлы PDF, DOC, JPG, JPEG и PNG.");
        exit;
    }
}
?>